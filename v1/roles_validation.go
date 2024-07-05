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

// PostRoles validates request body for the PostRoles method and delegates to the underlying handler.
func (v handlerWithValidation) PostRoles(w http.ResponseWriter, r *http.Request) {
	body := &resources.Role{}
	v.validateRequestBody(body, w, r, func(w http.ResponseWriter, r *http.Request) {
		v.ServerInterface.PostRoles(w, r)
	})
}

// PutRolesItem validates request body for the PutRolesItem method and delegates to the underlying handler.
func (v handlerWithValidation) PutRolesItem(w http.ResponseWriter, r *http.Request, id string) {
	body := &resources.Role{}
	v.validateRequestBody(body, w, r, func(w http.ResponseWriter, r *http.Request) {
		if body.Id == nil || id != *body.Id {
			writeErrorResponse(w, NewRequestBodyValidationError("role ID from path does not match the Role object"))
			return
		}
		v.ServerInterface.PutRolesItem(w, r, id)
	})
}

// PatchRolesItemEntitlements validates request body for the PatchRolesItemEntitlements method and delegates to the underlying handler.
func (v handlerWithValidation) PatchRolesItemEntitlements(w http.ResponseWriter, r *http.Request, id string) {
	body := &resources.RoleEntitlementsPatchRequestBody{}
	v.validateRequestBody(body, w, r, func(w http.ResponseWriter, r *http.Request) {
		v.ServerInterface.PatchRolesItemEntitlements(w, r, id)
	})
}
