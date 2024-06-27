package service

import (
	"context"

	v1 "github.com/canonical/rebac-admin-ui-handlers/v1"
	"github.com/canonical/rebac-admin-ui-handlers/v1/interfaces"
	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"

	"example/cmd/database"
)

// IdentityProvidersService implements the `IdentityProvidersService` interface.
type IdentityProvidersService struct {
	Database *database.Database
}

// For doc/test sake, to hint that the struct needs to implement a specific interface.
var _ interfaces.IdentityProvidersService = &IdentityProvidersService{}

// ListAvailableIdentityProviders returns the static list of supported identity providers.
func (s *IdentityProvidersService) ListAvailableIdentityProviders(ctx context.Context, params *resources.GetAvailableIdentityProvidersParams) (*resources.PaginatedResponse[resources.AvailableIdentityProvider], error) {
	return Paginate(s.Database.ListAvailableIdentityProviders(), params.Size, params.Page, params.NextToken, params.NextPageToken)
}

// ListIdentityProviders returns a list of registered identity providers configurations.
func (s *IdentityProvidersService) ListIdentityProviders(ctx context.Context, params *resources.GetIdentityProvidersParams) (*resources.PaginatedResponse[resources.IdentityProvider], error) {
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

	return Paginate(s.Database.ListIdentityProviders(), params.Size, params.Page, params.NextToken, params.NextPageToken)
}

// RegisterConfiguration register a new authentication provider configuration.
func (s *IdentityProvidersService) RegisterConfiguration(ctx context.Context, provider *resources.IdentityProvider) (*resources.IdentityProvider, error) {
	// For the sake of this example we allow everyone to call this method.

	added, err := s.Database.AddIdentityProvider(provider)
	if err != nil {
		return nil, v1.NewInvalidRequestError("either already exists or `name` is not provided")
	}
	return added, nil
}

// DeleteConfiguration removes an authentication provider configuration identified by `id`.
func (s *IdentityProvidersService) DeleteConfiguration(ctx context.Context, id string) (bool, error) {
	deleted := s.Database.DeleteIdentityProvider(id)

	if !deleted {
		// For idempotency, we return a nil error; the `false` value indicates
		// that the entry was already deleted/missing.
		return false, nil
	}
	return true, nil
}

// GetConfiguration returns the authentication provider configuration identified by `id`.
func (s *IdentityProvidersService) GetConfiguration(ctx context.Context, id string) (*resources.IdentityProvider, error) {
	// For the sake of this example we allow everyone to call this method.

	idp := s.Database.GetIdentityProvider(id)
	if idp == nil {
		return nil, v1.NewNotFoundError("")
	}
	return idp, nil
}

// UpdateConfiguration update the authentication provider configuration identified by `id`.
func (s *IdentityProvidersService) UpdateConfiguration(ctx context.Context, provider *resources.IdentityProvider) (*resources.IdentityProvider, error) {
	// For the sake of this example we allow everyone to call this method.

	updated := s.Database.UpdateIdentityProvider(ctx, provider)
	if updated == nil {
		return nil, v1.NewInvalidRequestError("either not found or `id` is missing")
	}
	return updated, nil
}
