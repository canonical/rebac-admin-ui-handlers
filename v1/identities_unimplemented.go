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

// unimplementedIdentitiesService represents a *Null Object* implementation of the the `IdentitiesService` interface.
type unimplementedIdentitiesService struct{}

func (s unimplementedIdentitiesService) ListIdentities(ctx context.Context, params *resources.GetIdentitiesParams) (*resources.PaginatedResponse[resources.Identity], error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedIdentitiesService) CreateIdentity(ctx context.Context, identity *resources.Identity) (*resources.Identity, error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedIdentitiesService) GetIdentity(ctx context.Context, identityId string) (*resources.Identity, error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedIdentitiesService) UpdateIdentity(ctx context.Context, identity *resources.Identity) (*resources.Identity, error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedIdentitiesService) DeleteIdentity(ctx context.Context, identityId string) (bool, error) {
	return false, NewNotImplementedError("")
}

func (s unimplementedIdentitiesService) GetIdentityGroups(ctx context.Context, identityId string, params *resources.GetIdentitiesItemGroupsParams) (*resources.PaginatedResponse[resources.Group], error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedIdentitiesService) PatchIdentityGroups(ctx context.Context, identityId string, groupPatches []resources.IdentityGroupsPatchItem) (bool, error) {
	return false, NewNotImplementedError("")
}

func (s unimplementedIdentitiesService) GetIdentityRoles(ctx context.Context, identityId string, params *resources.GetIdentitiesItemRolesParams) (*resources.PaginatedResponse[resources.Role], error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedIdentitiesService) PatchIdentityRoles(ctx context.Context, identityId string, rolePatches []resources.IdentityRolesPatchItem) (bool, error) {
	return false, NewNotImplementedError("")
}

func (s unimplementedIdentitiesService) GetIdentityEntitlements(ctx context.Context, identityId string, params *resources.GetIdentitiesItemEntitlementsParams) (*resources.PaginatedResponse[resources.EntityEntitlement], error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedIdentitiesService) PatchIdentityEntitlements(ctx context.Context, identityId string, entitlementPatches []resources.IdentityEntitlementsPatchItem) (bool, error) {
	return false, NewNotImplementedError("")
}
