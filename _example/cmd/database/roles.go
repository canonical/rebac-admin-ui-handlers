package database

import (
	"context"
	"errors"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

func (db *Database) ListRoles() []resources.Role {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return GetMapValues(db.Roles)
}

func (db *Database) AddRole(role *resources.Role) (*resources.Role, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	// Names should be unique.
	_, ok := db.Roles[role.Name]
	if ok {
		return nil, errors.New("already exists")
	}

	id := role.Name
	entry := *role
	entry.Id = &id
	db.Roles[id] = entry
	db.Persist()
	return &entry, nil
}

func (db *Database) GetRole(roleId string) *resources.Role {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	result, ok := db.Roles[roleId]
	if !ok {
		return nil
	}
	return &result
}

func (db *Database) UpdateRole(ctx context.Context, role *resources.Role) *resources.Role {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	if _, ok := db.Roles[*role.Id]; !ok {
		return nil
	}

	db.Roles[*role.Id] = *role
	db.Persist()
	return role
}

func (db *Database) DeleteRole(roleId string) bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	if _, ok := db.Roles[roleId]; !ok {
		return false
	}
	defer db.Persist()
	delete(db.Roles, roleId)
	db.Group2Role.RemoveRight(roleId)
	db.Identity2Role.RemoveRight(roleId)
	db.Role2Entitlement.RemoveLeft(roleId)
	db.Persist()
	return true
}

func (db *Database) GetRoleEntitlements(roleId string) []resources.EntityEntitlement {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	if _, ok := db.Roles[roleId]; !ok {
		return nil
	}
	return mapStringSlice(db.Role2Entitlement.GetRights(roleId), func(s string) resources.EntityEntitlement {
		return EntitlementFromString(s)
	})
}

func (db *Database) PatchRoleEntitlements(roleId string, additions, removals []string) *bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	if _, ok := db.Roles[roleId]; !ok {
		return nil
	}
	result := db.Role2Entitlement.PatchLeft(roleId, additions, removals)
	db.Persist()
	return &result
}
