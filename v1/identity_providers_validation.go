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
	"net/http"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// PostIdentityProviders validates request body for the PostIdentityProviders method and delegates to the underlying handler.
func (v handlerWithValidation) PostIdentityProviders(w http.ResponseWriter, r *http.Request) {
	body := &resources.IdentityProvider{}
	v.validateRequestBody(body, w, r, func(w http.ResponseWriter, r *http.Request) {
		v.ServerInterface.PostIdentityProviders(w, r)
	})
}

// PutIdentityProvidersItem validates request body for the PutIdentityProvidersItem method and delegates to the underlying handler.
func (v handlerWithValidation) PutIdentityProvidersItem(w http.ResponseWriter, r *http.Request, id string) {
	body := &resources.IdentityProvider{}
	v.validateRequestBody(body, w, r, func(w http.ResponseWriter, r *http.Request) {
		if body.Id == nil || id != *body.Id {
			writeErrorResponse(w, NewRequestBodyValidationError("identity provider ID from path does not match the IdentityProvider object"))
			return
		}
		v.ServerInterface.PutIdentityProvidersItem(w, r, id)
	})
}
