package database

import (
	"context"
	"errors"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// ListRoles returns the list of roles.
func (db *Database) ListRoles() []resources.Role {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return getMapValues(db.Roles)
}

// AddRole adds a new role.
func (db *Database) AddRole(role *resources.Role) (*resources.Role, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	// Names should be unique.
	_, ok := db.Roles[role.Name]
	if ok {
		return nil, errors.New("already exists")
	}

	id := role.Name
	entry := *role
	entry.Id = &id
	db.Roles[id] = entry
	db.isDirty = true
	return &entry, nil
}

// GetRole returns a role identified by given ID. If nothing found, the method
// returns nil.
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

// UpdateRole updates a role to the given value. If nothing found, the method
// returns nil.
func (db *Database) UpdateRole(ctx context.Context, role *resources.Role) *resources.Role {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	if _, ok := db.Roles[*role.Id]; !ok {
		return nil
	}

	db.Roles[*role.Id] = *role
	db.isDirty = true
	return role
}

// DeleteRole deletes a role identified by given ID. If nothing found, the
// method returns false. On a successful deletion, the method returns true.
func (db *Database) DeleteRole(roleId string) bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	if _, ok := db.Roles[roleId]; !ok {
		return false
	}

	delete(db.Roles, roleId)
	db.Group2Role.RemoveRight(roleId)
	db.Identity2Role.RemoveRight(roleId)
	db.Role2Entitlement.RemoveLeft(roleId)
	db.isDirty = true
	return true
}

// GetRoleEntitlements returns entitlements associated with a role identified by given ID.
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

// PatchRoleEntitlements patches entitlements associated with a role identified
// by given ID. If nothing found, the method returns nil. If nothing changes
// after applying the patch, the method returns (a pointer to) false; otherwise,
// it returns (a pointer to) true.
func (db *Database) PatchRoleEntitlements(roleId string, additions, removals []string) *bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	if _, ok := db.Roles[roleId]; !ok {
		return nil
	}
	result := db.Role2Entitlement.PatchLeft(roleId, additions, removals)
	db.isDirty = result
	return &result
}
