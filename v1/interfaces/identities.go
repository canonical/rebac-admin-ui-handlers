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

// IdentitiesService defines an abstract backend to handle Identities related operations.
type IdentitiesService interface {

	// ListIdentities returns a page of Identity objects of at least `size` elements if available
	ListIdentities(ctx context.Context, params *resources.GetIdentitiesParams) (*resources.PaginatedResponse[resources.Identity], error)
	// CreateIdentity creates a single Identity.
	CreateIdentity(ctx context.Context, identity *resources.Identity) (*resources.Identity, error)

	// GetIdentity returns a single Identity.
	GetIdentity(ctx context.Context, identityId string) (*resources.Identity, error)

	// UpdateIdentity updates an Identity.
	UpdateIdentity(ctx context.Context, identity *resources.Identity) (*resources.Identity, error)
	// DeleteIdentity deletes an Identity
	// returns (true, nil) in case an identity was successfully delete
	// return (false, error) in case something went wrong
	// implementors may want to return (false, nil) for idempotency cases
	DeleteIdentity(ctx context.Context, identityId string) (bool, error)

	// GetIdentityGroups returns a page of Groups for identity `identityId`.
	GetIdentityGroups(ctx context.Context, identityId string, params *resources.GetIdentitiesItemGroupsParams) (*resources.PaginatedResponse[resources.Group], error)
	// PatchIdentityGroups performs addition or removal of a Group to/from an Identity.
	PatchIdentityGroups(ctx context.Context, identityId string, groupPatches []resources.IdentityGroupsPatchItem) (bool, error)

	// GetIdentityRoles returns a page of Roles for identity `identityId`.
	GetIdentityRoles(ctx context.Context, identityId string, params *resources.GetIdentitiesItemRolesParams) (*resources.PaginatedResponse[resources.Role], error)
	// PatchIdentityRoles performs addition or removal of a Role to/from an Identity.
	PatchIdentityRoles(ctx context.Context, identityId string, rolePatches []resources.IdentityRolesPatchItem) (bool, error)

	// GetIdentityEntitlements returns a page of Entitlements for identity `identityId`.
	GetIdentityEntitlements(ctx context.Context, identityId string, params *resources.GetIdentitiesItemEntitlementsParams) (*resources.PaginatedResponse[resources.EntityEntitlement], error)
	// PatchIdentityEntitlements performs addition or removal of an Entitlement to/from an Identity.
	PatchIdentityEntitlements(ctx context.Context, identityId string, entitlementPatches []resources.IdentityEntitlementsPatchItem) (bool, error)
}
