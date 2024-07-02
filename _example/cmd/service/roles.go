// Copyright 2024 Canonical Ltd.

package service

import (
	"context"

	v1 "github.com/canonical/rebac-admin-ui-handlers/v1"
	"github.com/canonical/rebac-admin-ui-handlers/v1/interfaces"
	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"

	"example/cmd/database"
)

// RolesService implements the `RolesService` interface.
type RolesService struct {
	Database *database.Database
}

// For doc/test sake, to hint that the struct needs to implement a specific interface.
var _ interfaces.RolesService = &RolesService{}

// ListRoles returns a page of Role objects of at least `size` elements if available.
func (s *RolesService) ListRoles(ctx context.Context, params *resources.GetRolesParams) (*resources.PaginatedResponse[resources.Role], error) {
	// For the sake of this example we allow everyone to call this method. If it's not
	// the case, you can do the following to get the user:
	//
	//    raw, _ := v1.GetIdentityFromContext(ctx)
	//    user, _ := raw.(*User)
	//
	// And return this error if the user is not authorized:
	//
	//    return nil, v1.NewAuthorizationError("user cannot add group")
	//

	return Paginate(s.Database.ListRoles(), params.Size, params.Page, params.NextToken, params.NextPageToken, false)
}

// CreateRole creates a single Role.
func (s *RolesService) CreateRole(ctx context.Context, role *resources.Role) (*resources.Role, error) {
	// For the sake of this example we allow everyone to call this method.

	added, err := s.Database.AddRole(role)
	if err != nil {
		return nil, v1.NewInvalidRequestError("already exists")
	}
	return added, nil
}

// GetRole returns a single Role.
func (s *RolesService) GetRole(ctx context.Context, roleId string) (*resources.Role, error) {
	// For the sake of this example we allow everyone to call this method.

	role := s.Database.GetRole(roleId)
	if role == nil {
		return nil, v1.NewNotFoundError("")
	}
	return role, nil
}

// UpdateRole updates a Role.
func (s *RolesService) UpdateRole(ctx context.Context, role *resources.Role) (*resources.Role, error) {
	// For the sake of this example we allow everyone to call this method.

	updated := s.Database.UpdateRole(ctx, role)
	if updated == nil {
		return nil, v1.NewNotFoundError("")
	}
	return updated, nil
}

// DeleteRole deletes a Role
// returns (true, nil) in case a Role was successfully deleted
// returns (false, error) in case something went wrong
// implementors may want to return (false, nil) for idempotency cases.
func (s *RolesService) DeleteRole(ctx context.Context, roleId string) (bool, error) {
	deleted := s.Database.DeleteRole(roleId)

	if !deleted {
		// For idempotency, we return a nil error; the `false` value indicates
		// that the entry was already deleted/missing.
		return false, nil
	}
	return true, nil
}

// GetRoleEntitlements returns a page of Entitlements for Role `roleId`.
func (s *RolesService) GetRoleEntitlements(ctx context.Context, roleId string, params *resources.GetRolesItemEntitlementsParams) (*resources.PaginatedResponse[resources.EntityEntitlement], error) {
	// For the sake of this example we allow everyone to call this method.

	relatives := s.Database.GetRoleEntitlements(roleId)
	if relatives == nil {
		return nil, v1.NewNotFoundError("")
	}
	return Paginate(relatives, params.Size, params.Page, params.NextToken, params.NextPageToken, false)

}

// PatchRoleEntitlements performs addition or removal of an Entitlement to/from a Role.
func (s *RolesService) PatchRoleEntitlements(ctx context.Context, roleId string, entitlementPatches []resources.RoleEntitlementsPatchItem) (bool, error) {
	// For the sake of this example we allow everyone to call this method.

	additions := []string{}
	removals := []string{}
	for _, p := range entitlementPatches {
		if p.Op == "add" {
			additions = append(additions, database.EntitlementToString(p.Entitlement))
		} else if p.Op == "remove" {
			removals = append(removals, database.EntitlementToString(p.Entitlement))
		}
	}

	changed := s.Database.PatchRoleEntitlements(roleId, additions, removals)
	if changed == nil {
		return false, v1.NewNotFoundError("")
	}
	return *changed, nil
}
