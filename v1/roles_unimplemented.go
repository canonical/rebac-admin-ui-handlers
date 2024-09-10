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

package v1

import (
	"context"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// unimplementedRolesService represents a *Null Object* implementation of the the `RolesService` interface.
type unimplementedRolesService struct{}

func (s unimplementedRolesService) ListRoles(ctx context.Context, params *resources.GetRolesParams) (*resources.PaginatedResponse[resources.Role], error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedRolesService) CreateRole(ctx context.Context, role *resources.Role) (*resources.Role, error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedRolesService) GetRole(ctx context.Context, roleId string) (*resources.Role, error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedRolesService) UpdateRole(ctx context.Context, role *resources.Role) (*resources.Role, error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedRolesService) DeleteRole(ctx context.Context, roleId string) (bool, error) {
	return false, NewNotImplementedError("")
}

func (s unimplementedRolesService) GetRoleEntitlements(ctx context.Context, roleId string, params *resources.GetRolesItemEntitlementsParams) (*resources.PaginatedResponse[resources.EntityEntitlement], error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedRolesService) PatchRoleEntitlements(ctx context.Context, roleId string, entitlementPatches []resources.RoleEntitlementsPatchItem) (bool, error) {
	return false, NewNotImplementedError("")
}
