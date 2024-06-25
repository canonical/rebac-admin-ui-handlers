package service

import (
	"context"

	v1 "github.com/canonical/rebac-admin-ui-handlers/v1"
	"github.com/canonical/rebac-admin-ui-handlers/v1/interfaces"
	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"

	"example/cmd/database"
)

type IdentitiesService struct {
	Database *database.Database
}

// For doc/test sake, to hint that the struct needs to implement a specific interface.
var _ interfaces.IdentitiesService = &IdentitiesService{}

// ListIdentities returns a page of Identity objects of at least `size` elements if available
func (s *IdentitiesService) ListIdentities(ctx context.Context, params *resources.GetIdentitiesParams) (*resources.PaginatedResponse[resources.Identity], error) {
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

	return Paginate(s.Database.ListIdentities(), params.Size, params.Page, params.NextToken, params.NextPageToken)
}

// CreateIdentity creates a single Identity.
func (s *IdentitiesService) CreateIdentity(ctx context.Context, identity *resources.Identity) (*resources.Identity, error) {
	// For the sake of this example we allow everyone to call this method.

	added, err := s.Database.AddIdentity(identity)
	if err != nil {
		return nil, v1.NewInvalidRequestError("already exists")
	}
	return added, nil
}

// GetIdentity returns a single Identity.
func (s *IdentitiesService) GetIdentity(ctx context.Context, identityId string) (*resources.Identity, error) {
	// For the sake of this example we allow everyone to call this method.

	identity := s.Database.GetIdentity(identityId)
	if identity == nil {
		return nil, v1.NewNotFoundError("")
	}
	return identity, nil
}

// UpdateIdentity updates an Identity.
func (s *IdentitiesService) UpdateIdentity(ctx context.Context, identity *resources.Identity) (*resources.Identity, error) {
	// For the sake of this example we allow everyone to call this method.

	updated := s.Database.UpdateIdentity(ctx, identity)
	if updated == nil {
		return nil, v1.NewNotFoundError("")
	}
	return updated, nil
}

// DeleteIdentity deletes an Identity
// returns (true, nil) in case an identity was successfully delete
// return (false, error) in case something went wrong
// implementors may want to return (false, nil) for idempotency cases
func (s *IdentitiesService) DeleteIdentity(ctx context.Context, identityId string) (bool, error) {
	deleted := s.Database.DeleteIdentity(identityId)

	if !deleted {
		// For idempotency, we return a nil error; the `false` value indicates
		// that the entry was already deleted/missing.
		return false, nil
	}
	return true, nil
}

// GetIdentityGroups returns a page of Groups for identity `identityId`.
func (s *IdentitiesService) GetIdentityGroups(ctx context.Context, identityId string, params *resources.GetIdentitiesItemGroupsParams) (*resources.PaginatedResponse[resources.Group], error) {
	// For the sake of this example we allow everyone to call this method.

	relatives := s.Database.GetIdentityGroups(identityId)
	if relatives == nil {
		return nil, v1.NewNotFoundError("")
	}
	return Paginate(relatives, params.Size, params.Page, params.NextToken, params.NextPageToken)
}

// PatchIdentityGroups performs addition or removal of a Group to/from an Identity.
func (s *IdentitiesService) PatchIdentityGroups(ctx context.Context, identityId string, groupPatches []resources.IdentityGroupsPatchItem) (bool, error) {
	// For the sake of this example we allow everyone to call this method.

	additions := []string{}
	removals := []string{}
	for _, p := range groupPatches {
		if p.Op == "add" {
			additions = append(additions, p.Group)
		} else if p.Op == "remove" {
			removals = append(removals, p.Group)
		}
	}

	changed := s.Database.PatchIdentityGroups(identityId, additions, removals)
	if changed == nil {
		return false, v1.NewNotFoundError("")
	}
	return *changed, nil
}

// GetIdentityRoles returns a page of Roles for identity `identityId`.
func (s *IdentitiesService) GetIdentityRoles(ctx context.Context, identityId string, params *resources.GetIdentitiesItemRolesParams) (*resources.PaginatedResponse[resources.Role], error) {
	// For the sake of this example we allow everyone to call this method.

	relatives := s.Database.GetIdentityRoles(identityId)
	if relatives == nil {
		return nil, v1.NewNotFoundError("")
	}
	return Paginate(relatives, params.Size, params.Page, params.NextToken, params.NextPageToken)

}

// PatchIdentityRoles performs addition or removal of a Role to/from an Identity.
func (s *IdentitiesService) PatchIdentityRoles(ctx context.Context, identityId string, rolePatches []resources.IdentityRolesPatchItem) (bool, error) {
	// For the sake of this example we allow everyone to call this method.

	additions := []string{}
	removals := []string{}
	for _, p := range rolePatches {
		if p.Op == "add" {
			additions = append(additions, p.Role)
		} else if p.Op == "remove" {
			removals = append(removals, p.Role)
		}
	}

	changed := s.Database.PatchIdentityRoles(identityId, additions, removals)
	if changed == nil {
		return false, v1.NewNotFoundError("")
	}
	return *changed, nil
}

// GetIdentityEntitlements returns a page of Entitlements for identity `identityId`.
func (s *IdentitiesService) GetIdentityEntitlements(ctx context.Context, identityId string, params *resources.GetIdentitiesItemEntitlementsParams) (*resources.PaginatedResponse[resources.EntityEntitlement], error) {
	// For the sake of this example we allow everyone to call this method.

	relatives := s.Database.GetIdentityEntitlements(identityId)
	if relatives == nil {
		return nil, v1.NewNotFoundError("")
	}
	return Paginate(relatives, params.Size, params.Page, params.NextToken, params.NextPageToken)
}

// PatchIdentityEntitlements performs addition or removal of an Entitlement to/from an Identity.
func (s *IdentitiesService) PatchIdentityEntitlements(ctx context.Context, identityId string, entitlementPatches []resources.IdentityEntitlementsPatchItem) (bool, error) {
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

	changed := s.Database.PatchIdentityEntitlements(identityId, additions, removals)
	if changed == nil {
		return false, v1.NewNotFoundError("")
	}
	return *changed, nil
}
