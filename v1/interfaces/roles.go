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

package interfaces

import (
	"context"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// RolesService defines an abstract backend to handle Roles related operations.
type RolesService interface {
	// ListRoles returns a page of Role objects of at least `size` elements if available.
	ListRoles(ctx context.Context, params *resources.GetRolesParams) (*resources.PaginatedResponse[resources.Role], error)
	// CreateRole creates a single Role.
	CreateRole(ctx context.Context, role *resources.Role) (*resources.Role, error)

	// GetRole returns a single Role.
	GetRole(ctx context.Context, roleId string) (*resources.Role, error)
	// UpdateRole updates a Role.
	UpdateRole(ctx context.Context, role *resources.Role) (*resources.Role, error)
	// DeleteRole deletes a Role
	// returns (true, nil) in case a Role was successfully deleted
	// returns (false, error) in case something went wrong
	// implementors may want to return (false, nil) for idempotency cases.
	DeleteRole(ctx context.Context, roleId string) (bool, error)

	// GetRoleEntitlements returns a page of Entitlements for Role `roleId`.
	GetRoleEntitlements(ctx context.Context, roleId string, params *resources.GetRolesItemEntitlementsParams) (*resources.PaginatedResponse[resources.EntityEntitlement], error)
	// PatchRoleEntitlements performs addition or removal of an Entitlement to/from a Role.
	PatchRoleEntitlements(ctx context.Context, roleId string, entitlementPatches []resources.RoleEntitlementsPatchItem) (bool, error)
}
