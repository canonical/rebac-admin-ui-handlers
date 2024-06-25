package database

import (
	"context"
	"errors"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

func (db *Database) ListIdentities() []resources.Identity {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return GetMapValues(db.Identities)
}

func (db *Database) AddIdentity(identity *resources.Identity) (*resources.Identity, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	// Names should be unique.
	_, ok := db.Identities[identity.Email]
	if ok {
		return nil, errors.New("already exists")
	}

	id := identity.Email
	entry := *identity
	entry.Id = &id
	db.Identities[id] = entry
	db.Persist()
	return &entry, nil
}

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

func (db *Database) UpdateIdentity(ctx context.Context, identity *resources.Identity) *resources.Identity {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	if _, ok := db.Identities[*identity.Id]; !ok {
		return nil
	}
	db.Identities[*identity.Id] = *identity
	db.Persist()
	return identity
}

func (db *Database) DeleteIdentity(identityId string) bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	if _, ok := db.Identities[identityId]; !ok {
		return false
	}
	delete(db.Identities, identityId)
	db.Group2Identity.RemoveRight(identityId)
	db.Identity2Entitlement.RemoveLeft(identityId)
	db.Identity2Role.RemoveLeft(identityId)
	db.Persist()
	return true
}

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

func (db *Database) PatchIdentityGroups(identityId string, additions, removals []string) *bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	if _, ok := db.Identities[identityId]; !ok {
		return nil
	}
	result := db.Group2Identity.PatchRight(identityId, additions, removals)
	db.Persist()
	return &result
}

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

func (db *Database) PatchIdentityRoles(identityId string, additions, removals []string) *bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	if _, ok := db.Identities[identityId]; !ok {
		return nil
	}
	result := db.Identity2Role.PatchLeft(identityId, additions, removals)
	db.Persist()
	return &result
}

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

func (db *Database) PatchIdentityEntitlements(identityId string, additions, removals []string) *bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	if _, ok := db.Identities[identityId]; !ok {
		return nil
	}
	result := db.Identity2Entitlement.PatchLeft(identityId, additions, removals)
	db.Persist()
	return &result
}
