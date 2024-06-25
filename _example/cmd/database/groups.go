package database

import (
	"context"
	"errors"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

func (db *Database) ListGroups() []resources.Group {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return GetMapValues(db.Groups)
}

func (db *Database) AddGroup(group *resources.Group) (*resources.Group, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	// Names should be unique.
	_, ok := db.Groups[group.Name]
	if ok {
		return nil, errors.New("already exists")
	}

	id := group.Name
	entry := *group
	entry.Id = &id
	db.Groups[id] = entry
	db.Persist()
	return &entry, nil
}

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

func (db *Database) UpdateGroup(ctx context.Context, group *resources.Group) *resources.Group {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	if _, ok := db.Groups[*group.Id]; !ok {
		return nil
	}

	db.Groups[*group.Id] = *group
	db.Persist()
	return group
}

func (db *Database) DeleteGroup(groupId string) bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	if _, ok := db.Groups[groupId]; !ok {
		return false
	}
	defer db.Persist()
	delete(db.Groups, groupId)
	db.Group2Identity.RemoveLeft(groupId)
	db.Group2Role.RemoveLeft(groupId)
	db.Group2Entitlement.RemoveLeft(groupId)
	db.Persist()
	return true
}

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

func (db *Database) PatchGroupIdentities(groupId string, additions, removals []string) *bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	if _, ok := db.Groups[groupId]; !ok {
		return nil
	}
	result := db.Group2Identity.PatchLeft(groupId, additions, removals)
	db.Persist()
	return &result
}

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

func (db *Database) PatchGroupRoles(groupId string, additions, removals []string) *bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	if _, ok := db.Groups[groupId]; !ok {
		return nil
	}
	result := db.Group2Role.PatchLeft(groupId, additions, removals)
	db.Persist()
	return &result
}

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

func (db *Database) PatchGroupEntitlements(groupId string, additions, removals []string) *bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	if _, ok := db.Groups[groupId]; !ok {
		return nil
	}
	result := db.Group2Entitlement.PatchLeft(groupId, additions, removals)
	db.Persist()
	return &result
}
