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

// unimplementedIdentityProvidersService represents a *Null Object* implementation of the the `IdentityProvidersService` interface.
type unimplementedIdentityProvidersService struct{}

func (s unimplementedIdentityProvidersService) ListAvailableIdentityProviders(ctx context.Context, params *resources.GetAvailableIdentityProvidersParams) (*resources.PaginatedResponse[resources.AvailableIdentityProvider], error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedIdentityProvidersService) ListIdentityProviders(ctx context.Context, params *resources.GetIdentityProvidersParams) (*resources.PaginatedResponse[resources.IdentityProvider], error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedIdentityProvidersService) RegisterConfiguration(ctx context.Context, provider *resources.IdentityProvider) (*resources.IdentityProvider, error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedIdentityProvidersService) DeleteConfiguration(ctx context.Context, id string) (bool, error) {
	return false, NewNotImplementedError("")
}

func (s unimplementedIdentityProvidersService) GetConfiguration(ctx context.Context, id string) (*resources.IdentityProvider, error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedIdentityProvidersService) UpdateConfiguration(ctx context.Context, provider *resources.IdentityProvider) (*resources.IdentityProvider, error) {
	return nil, NewNotImplementedError("")
}
