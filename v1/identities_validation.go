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

// PostIdentities validates request body for the PostIdentities method and delegates to the underlying handler.
func (v handlerWithValidation) PostIdentities(w http.ResponseWriter, r *http.Request) {
	body := &resources.Identity{}
	v.validateRequestBody(body, w, r, func(w http.ResponseWriter, r *http.Request) {
		v.ServerInterface.PostIdentities(w, r)
	})
}

// PutIdentitiesItem validates request body for the PutIdentitiesItem method and delegates to the underlying handler.
func (v handlerWithValidation) PutIdentitiesItem(w http.ResponseWriter, r *http.Request, id string) {
	body := &resources.Identity{}
	v.validateRequestBody(body, w, r, func(w http.ResponseWriter, r *http.Request) {
		if body.Id == nil || id != *body.Id {
			writeErrorResponse(w, NewRequestBodyValidationError("identity ID from path does not match the Identity object"))
			return
		}
		v.ServerInterface.PutIdentitiesItem(w, r, id)
	})
}

// PatchIdentitiesItemEntitlements validates request body for the PatchIdentitiesItemEntitlements method and delegates to the underlying handler.
func (v handlerWithValidation) PatchIdentitiesItemEntitlements(w http.ResponseWriter, r *http.Request, id string) {
	body := &resources.IdentityEntitlementsPatchRequestBody{}
	v.validateRequestBody(body, w, r, func(w http.ResponseWriter, r *http.Request) {
		v.ServerInterface.PatchIdentitiesItemEntitlements(w, r, id)
	})
}

// PatchIdentitiesItemGroups validates request body for the PatchIdentitiesItemGroups method and delegates to the underlying handler.
func (v handlerWithValidation) PatchIdentitiesItemGroups(w http.ResponseWriter, r *http.Request, id string) {
	body := &resources.IdentityGroupsPatchRequestBody{}
	v.validateRequestBody(body, w, r, func(w http.ResponseWriter, r *http.Request) {
		v.ServerInterface.PatchIdentitiesItemGroups(w, r, id)
	})
}

// PatchIdentitiesItemRoles validates request body for the PatchIdentitiesItemRoles method and delegates to the underlying handler.
func (v handlerWithValidation) PatchIdentitiesItemRoles(w http.ResponseWriter, r *http.Request, id string) {
	body := &resources.IdentityRolesPatchRequestBody{}
	v.validateRequestBody(body, w, r, func(w http.ResponseWriter, r *http.Request) {
		v.ServerInterface.PatchIdentitiesItemRoles(w, r, id)
	})
}
