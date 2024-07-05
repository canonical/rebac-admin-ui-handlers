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

// IdentityProvidersService defines an abstract backend to handle Roles related operations.
type IdentityProvidersService interface {

	// ListAvailableIdentityProviders returns the static list of supported identity providers.
	ListAvailableIdentityProviders(ctx context.Context, params *resources.GetAvailableIdentityProvidersParams) (*resources.PaginatedResponse[resources.AvailableIdentityProvider], error)

	// ListIdentityProviders returns a list of registered identity providers configurations.
	ListIdentityProviders(ctx context.Context, params *resources.GetIdentityProvidersParams) (*resources.PaginatedResponse[resources.IdentityProvider], error)

	// RegisterConfiguration register a new authentication provider configuration.
	RegisterConfiguration(ctx context.Context, provider *resources.IdentityProvider) (*resources.IdentityProvider, error)

	// DeleteConfiguration removes an authentication provider configuration identified by `id`.
	DeleteConfiguration(ctx context.Context, id string) (bool, error)

	// GetConfiguration returns the authentication provider configuration identified by `id`.
	GetConfiguration(ctx context.Context, id string) (*resources.IdentityProvider, error)

	// UpdateConfiguration update the authentication provider configuration identified by `id`.
	UpdateConfiguration(ctx context.Context, provider *resources.IdentityProvider) (*resources.IdentityProvider, error)
}
