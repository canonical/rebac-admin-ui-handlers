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

// Package v1 provides HTTP handlers that implement the ReBAC Admin OpenAPI spec.
// This package delegates authorization and data manipulation to user-defined
// "backend"s that implement the designated abstractions.
package v1

import (
	"net/http"
	"strings"

	"github.com/canonical/rebac-admin-ui-handlers/v1/interfaces"
	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// ReBACAdminBackendParams contains references to user-defined implementation
// of required abstractions, called "backend"s.
type ReBACAdminBackendParams struct {
	// Authenticator is required.
	Authenticator            interfaces.Authenticator
	AuthenticatorErrorMapper ErrorResponseMapper

	Identities            interfaces.IdentitiesService
	IdentitiesErrorMapper ErrorResponseMapper

	Roles            interfaces.RolesService
	RolesErrorMapper ErrorResponseMapper

	IdentityProviders            interfaces.IdentityProvidersService
	IdentityProvidersErrorMapper ErrorResponseMapper

	Capabilities            interfaces.CapabilitiesService
	CapabilitiesErrorMapper ErrorResponseMapper

	Entitlements            interfaces.EntitlementsService
	EntitlementsErrorMapper ErrorResponseMapper

	Groups            interfaces.GroupsService
	GroupsErrorMapper ErrorResponseMapper

	Resources            interfaces.ResourcesService
	ResourcesErrorMapper ErrorResponseMapper
}

// ReBACAdminBackend represents the ReBAC admin backend as a whole package.
type ReBACAdminBackend struct {
	params  ReBACAdminBackendParams
	handler resources.ServerInterface
}

// NewReBACAdminBackend returns a new ReBACAdminBackend instance, configured
// with given backends.
func NewReBACAdminBackend(params ReBACAdminBackendParams) (*ReBACAdminBackend, error) {
	return newReBACAdminBackendWithService(
		params,
		newHandlerWithValidation(&handler{
			Identities:            params.Identities,
			IdentitiesErrorMapper: params.IdentitiesErrorMapper,

			Roles:            params.Roles,
			RolesErrorMapper: params.RolesErrorMapper,

			IdentityProviders:            params.IdentityProviders,
			IdentityProvidersErrorMapper: params.IdentityProvidersErrorMapper,

			Capabilities:            params.Capabilities,
			CapabilitiesErrorMapper: params.CapabilitiesErrorMapper,

			Entitlements:            params.Entitlements,
			EntitlementsErrorMapper: params.EntitlementsErrorMapper,

			Groups:            params.Groups,
			GroupsErrorMapper: params.GroupsErrorMapper,

			Resources:            params.Resources,
			ResourcesErrorMapper: params.ResourcesErrorMapper,
		})), nil
}

// newReBACAdminBackendWithService returns a new ReBACAdminBackend instance, configured
// with given backends and service implementation.
//
// This is intended for internal/test use cases.
func newReBACAdminBackendWithService(params ReBACAdminBackendParams, handler resources.ServerInterface) *ReBACAdminBackend {
	return &ReBACAdminBackend{
		params:  params,
		handler: handler,
	}
}

// Handler returns HTTP handlers implementing the ReBAC Admin OpenAPI spec.
func (b *ReBACAdminBackend) Handler(baseURL string) http.Handler {
	var middlewares []resources.MiddlewareFunc
	if b.params.Authenticator != nil {
		middlewares = append(middlewares, b.authenticationMiddleware())
	}

	baseURL, _ = strings.CutSuffix(baseURL, "/")
	return resources.HandlerWithOptions(b.handler, resources.ChiServerOptions{
		BaseURL:     baseURL + "/v1",
		Middlewares: middlewares,
		ErrorHandlerFunc: func(w http.ResponseWriter, _ *http.Request, err error) {
			writeErrorResponse(w, err)
		},
	})
}
