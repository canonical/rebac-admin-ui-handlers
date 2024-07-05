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

// ListGroup returns the list of groups.
func (db *Database) ListGroups() []resources.Group {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return getMapValues(db.Groups)
}

// AddGroup adds a new group.
func (db *Database) AddGroup(group *resources.Group) (*resources.Group, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	// Names should be unique.
	_, ok := db.Groups[group.Name]
	if ok {
		return nil, errors.New("already exists")
	}

	id := group.Name
	entry := *group
	entry.Id = &id
	db.Groups[id] = entry
	db.isDirty = true
	return &entry, nil
}

// GetGroup returns a group identified by given ID. If nothing found, the method
// returns nil.
func (db *Database) GetGroup(groupId string) *resources.Group {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	result, ok := db.Groups[groupId]
	if !ok {
		return nil
	}
	return &result
}

// UpdateGroup updates a group to the given value. If nothing found, the method
// returns nil.
func (db *Database) UpdateGroup(ctx context.Context, group *resources.Group) *resources.Group {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	if _, ok := db.Groups[*group.Id]; !ok {
		return nil
	}

	db.Groups[*group.Id] = *group
	db.isDirty = true
	return group
}

// DeleteGroup deletes a group identified by given ID. If nothing found, the
// method returns false. On a successful deletion, the method returns true.
func (db *Database) DeleteGroup(groupId string) bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	if _, ok := db.Groups[groupId]; !ok {
		return false
	}

	delete(db.Groups, groupId)
	db.Group2Identity.RemoveLeft(groupId)
	db.Group2Role.RemoveLeft(groupId)
	db.Group2Entitlement.RemoveLeft(groupId)
	db.isDirty = true
	return true
}

// GetGroupIdentities returns identities associated with a group identified by given ID.
func (db *Database) GetGroupIdentities(groupId string) []resources.Identity {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	if _, ok := db.Groups[groupId]; !ok {
		return nil
	}
	return mapStringSlice(db.Group2Identity.GetRights(groupId), func(s string) resources.Identity {
		return db.Identities[s]
	})
}

// PatchGroupIdentities patches identities associated with a group identified
// by given ID. If nothing found, the method returns nil. If nothing changes
// after applying the patch, the method returns (a pointer to) false; otherwise,
// it returns (a pointer to) true.
func (db *Database) PatchGroupIdentities(groupId string, additions, removals []string) *bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	if _, ok := db.Groups[groupId]; !ok {
		return nil
	}
	result := db.Group2Identity.PatchLeft(groupId, additions, removals)
	db.isDirty = result
	return &result
}

// GetGroupRoles returns roles associated with a group identified by given ID.
func (db *Database) GetGroupRoles(groupId string) []resources.Role {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	if _, ok := db.Groups[groupId]; !ok {
		return nil
	}
	return mapStringSlice(db.Group2Role.GetRights(groupId), func(s string) resources.Role {
		return db.Roles[s]
	})
}

// PatchGroupRoles patches roles associated with a group identified by given ID.
// If nothing found, the method returns nil. If nothing changes after applying
// the patch, the method returns (a pointer to) false; otherwise, it returns (a
// pointer to) true.
func (db *Database) PatchGroupRoles(groupId string, additions, removals []string) *bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	if _, ok := db.Groups[groupId]; !ok {
		return nil
	}
	result := db.Group2Role.PatchLeft(groupId, additions, removals)
	db.isDirty = result
	return &result
}

// GetGroupEntitlements returns entitlements associated with a group identified by given ID.
func (db *Database) GetGroupEntitlements(groupId string) []resources.EntityEntitlement {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	if _, ok := db.Groups[groupId]; !ok {
		return nil
	}
	return mapStringSlice(db.Group2Entitlement.GetRights(groupId), func(s string) resources.EntityEntitlement {
		return EntitlementFromString(s)
	})
}

// PatchGroupEntitlements patches entitlements associated with a group
// identified by given ID. If nothing found, the method returns nil. If nothing
// changes after applying the patch, the method returns (a pointer to) false;
// otherwise, it returns (a pointer to) true.
func (db *Database) PatchGroupEntitlements(groupId string, additions, removals []string) *bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	if _, ok := db.Groups[groupId]; !ok {
		return nil
	}
	result := db.Group2Entitlement.PatchLeft(groupId, additions, removals)
	db.isDirty = result
	return &result
}
