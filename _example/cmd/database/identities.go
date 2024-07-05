// Copyright (C) 2024 Canonical Ltd.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package database

import (
	"context"
	"errors"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// ListIdentities returns the list of identities.
func (db *Database) ListIdentities() []resources.Identity {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return getMapValues(db.Identities)
}

// AddIdentity adds a new identity.
func (db *Database) AddIdentity(identity *resources.Identity) (*resources.Identity, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	// Names should be unique.
	_, ok := db.Identities[identity.Email]
	if ok {
		return nil, errors.New("already exists")
	}

	id := identity.Email
	entry := *identity
	entry.Id = &id
	db.Identities[id] = entry
	db.isDirty = true
	return &entry, nil
}

// GetIdentity returns an identity identified by given ID. If nothing found, the
// method returns nil.
func (db *Database) GetIdentity(identityId string) *resources.Identity {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	result, ok := db.Identities[identityId]
	if !ok {
		return nil
	}
	return &result
}

// UpdateIdentity updates an identity to the given value. If nothing found, the
// method returns nil.
func (db *Database) UpdateIdentity(ctx context.Context, identity *resources.Identity) *resources.Identity {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	if _, ok := db.Identities[*identity.Id]; !ok {
		return nil
	}
	db.Identities[*identity.Id] = *identity
	db.isDirty = true
	return identity
}

// DeleteIdentity deletes an identity identified by given ID. If nothing found,
// the method returns false. On a successful deletion, the method returns true.
func (db *Database) DeleteIdentity(identityId string) bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	if _, ok := db.Identities[identityId]; !ok {
		return false
	}
	delete(db.Identities, identityId)
	db.Group2Identity.RemoveRight(identityId)
	db.Identity2Entitlement.RemoveLeft(identityId)
	db.Identity2Role.RemoveLeft(identityId)
	db.isDirty = true
	return true
}

// GetIdentityGroups returns groups associated with an identity identified by given ID.
func (db *Database) GetIdentityGroups(identityId string) []resources.Group {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	if _, ok := db.Identities[identityId]; !ok {
		return nil
	}
	return mapStringSlice(db.Group2Identity.GetLefts(identityId), func(s string) resources.Group {
		return db.Groups[s]
	})
}

// PatchIdentityGroups patches groups associated with an identity identified by
// given ID. If nothing found, the method returns nil. If nothing changes after
// applying the patch, the method returns (a pointer to) false; otherwise, it
// returns (a pointer to) true.
func (db *Database) PatchIdentityGroups(identityId string, additions, removals []string) *bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	if _, ok := db.Identities[identityId]; !ok {
		return nil
	}
	result := db.Group2Identity.PatchRight(identityId, additions, removals)
	db.isDirty = result
	return &result
}

// GetIdentityRoles returns roles associated with an identity identified by given ID.
func (db *Database) GetIdentityRoles(identityId string) []resources.Role {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	if _, ok := db.Identities[identityId]; !ok {
		return nil
	}
	return mapStringSlice(db.Identity2Role.GetRights(identityId), func(s string) resources.Role {
		return db.Roles[s]
	})
}

// PatchIdentityRoles patches roles associated with an identity identified by
// given ID. If nothing found, the method returns nil. If nothing changes after
// applying the patch, the method returns (a pointer to) false; otherwise, it
// returns (a pointer to) true.
func (db *Database) PatchIdentityRoles(identityId string, additions, removals []string) *bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	if _, ok := db.Identities[identityId]; !ok {
		return nil
	}
	result := db.Identity2Role.PatchLeft(identityId, additions, removals)
	db.isDirty = result
	return &result
}

// GetIdentityEntitlements returns entitlements associated with an identity identified by given ID.
func (db *Database) GetIdentityEntitlements(identityId string) []resources.EntityEntitlement {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	if _, ok := db.Identities[identityId]; !ok {
		return nil
	}
	return mapStringSlice(db.Identity2Entitlement.GetRights(identityId), func(s string) resources.EntityEntitlement {
		return EntitlementFromString(s)
	})
}

// PatchIdentityEntitlements patches entitlements associated with an identity
// identified by given ID. If nothing found, the method returns nil. If nothing
// changes after applying the patch, the method returns (a pointer to) false;
// otherwise, it returns (a pointer to) true.
func (db *Database) PatchIdentityEntitlements(identityId string, additions, removals []string) *bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	if _, ok := db.Identities[identityId]; !ok {
		return nil
	}
	result := db.Identity2Entitlement.PatchLeft(identityId, additions, removals)
	db.isDirty = result
	return &result
}
