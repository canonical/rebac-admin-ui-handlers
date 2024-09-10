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

// unimplementedGroupsService represents a *Null Object* implementation of the the `GroupsService` interface.
type unimplementedGroupsService struct{}

func (s unimplementedGroupsService) ListGroups(ctx context.Context, params *resources.GetGroupsParams) (*resources.PaginatedResponse[resources.Group], error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedGroupsService) CreateGroup(ctx context.Context, group *resources.Group) (*resources.Group, error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedGroupsService) GetGroup(ctx context.Context, groupId string) (*resources.Group, error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedGroupsService) UpdateGroup(ctx context.Context, group *resources.Group) (*resources.Group, error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedGroupsService) DeleteGroup(ctx context.Context, groupId string) (bool, error) {
	return false, NewNotImplementedError("")
}

func (s unimplementedGroupsService) GetGroupIdentities(ctx context.Context, groupId string, params *resources.GetGroupsItemIdentitiesParams) (*resources.PaginatedResponse[resources.Identity], error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedGroupsService) PatchGroupIdentities(ctx context.Context, groupId string, identityPatches []resources.GroupIdentitiesPatchItem) (bool, error) {
	return false, NewNotImplementedError("")
}

func (s unimplementedGroupsService) GetGroupRoles(ctx context.Context, groupId string, params *resources.GetGroupsItemRolesParams) (*resources.PaginatedResponse[resources.Role], error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedGroupsService) PatchGroupRoles(ctx context.Context, groupId string, rolePatches []resources.GroupRolesPatchItem) (bool, error) {
	return false, NewNotImplementedError("")
}

func (s unimplementedGroupsService) GetGroupEntitlements(ctx context.Context, groupId string, params *resources.GetGroupsItemEntitlementsParams) (*resources.PaginatedResponse[resources.EntityEntitlement], error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedGroupsService) PatchGroupEntitlements(ctx context.Context, groupId string, entitlementPatches []resources.GroupEntitlementsPatchItem) (bool, error) {
	return false, NewNotImplementedError("")
}
