package service

import (
	"context"

	"github.com/canonical/rebac-admin-ui-handlers/v1/interfaces"
	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"

	"example/cmd/database"
)

// ResourcesService implements the `ResourcesService` interface.
type ResourcesService struct {
	Database *database.Database
}

// For doc/test sake, to hint that the struct needs to implement a specific interface.
var _ interfaces.ResourcesService = &ResourcesService{}

// ResourcesService defines an abstract backend to handle Resources related operations.
func (s *ResourcesService) ListResources(ctx context.Context, params *resources.GetResourcesParams) (*resources.PaginatedResponse[resources.Resource], error) {
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

	return Paginate(s.Database.ListUserResources(), params.Size, params.Page, params.NextToken, params.NextPageToken)

}
