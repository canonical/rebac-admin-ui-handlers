package service

import (
	"context"
	"example/cmd/database"

	"github.com/canonical/rebac-admin-ui-handlers/v1/interfaces"
	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// EntitlementsService implements the `EntitlementsService` interface.
type EntitlementsService struct {
	Database *database.Database
}

// For doc/test sake, to hint that the struct needs to implement a specific interface.
var _ interfaces.EntitlementsService = &EntitlementsService{}

// ListEntitlements returns the list of entitlements in JSON format.
func (s *EntitlementsService) ListEntitlements(ctx context.Context, params *resources.GetEntitlementsParams) (*resources.PaginatedResponse[resources.EntityEntitlement], error) {
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

	return Paginate(s.Database.ListUserEntitlements(), params.Size, params.Page, params.NextToken, params.NextPageToken)
}

// RawEntitlements returns the list of entitlements as raw text.
func (s *EntitlementsService) RawEntitlements(ctx context.Context) (string, error) {
	// This should return the raw OpenFGA authorization model, so we just return
	// a fake one:

	return s.Database.GetAuthModel(), nil
}
