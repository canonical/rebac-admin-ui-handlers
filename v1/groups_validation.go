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

// PostGroups validates request body for the PostGroups method and delegates to the underlying handler.
func (v handlerWithValidation) PostGroups(w http.ResponseWriter, r *http.Request) {
	body := &resources.Group{}
	v.validateRequestBody(body, w, r, func(w http.ResponseWriter, r *http.Request) {
		v.ServerInterface.PostGroups(w, r)
	})
}

// PutGroupsItem validates request body for the PutGroupsItem method and delegates to the underlying handler.
func (v handlerWithValidation) PutGroupsItem(w http.ResponseWriter, r *http.Request, id string) {
	body := &resources.Group{}
	v.validateRequestBody(body, w, r, func(w http.ResponseWriter, r *http.Request) {
		if body.Id == nil || id != *body.Id {
			writeErrorResponse(w, NewRequestBodyValidationError("group ID from path does not match the Group object"))
			return
		}
		v.ServerInterface.PutGroupsItem(w, r, id)
	})
}

// PatchGroupsItemEntitlements validates request body for the PatchGroupsItemEntitlements method and delegates to the underlying handler.
func (v handlerWithValidation) PatchGroupsItemEntitlements(w http.ResponseWriter, r *http.Request, id string) {
	body := &resources.GroupEntitlementsPatchRequestBody{}
	v.validateRequestBody(body, w, r, func(w http.ResponseWriter, r *http.Request) {
		v.ServerInterface.PatchGroupsItemEntitlements(w, r, id)
	})
}

// PatchGroupsItemIdentities validates request body for the PatchGroupsItemIdentities method and delegates to the underlying handler.
func (v handlerWithValidation) PatchGroupsItemIdentities(w http.ResponseWriter, r *http.Request, id string) {
	body := &resources.GroupIdentitiesPatchRequestBody{}
	v.validateRequestBody(body, w, r, func(w http.ResponseWriter, r *http.Request) {
		v.ServerInterface.PatchGroupsItemIdentities(w, r, id)
	})
}

// PatchGroupsItemRoles validates request body for the PatchGroupsItemRoles method and delegates to the underlying handler.
func (v handlerWithValidation) PatchGroupsItemRoles(w http.ResponseWriter, r *http.Request, id string) {
	body := &resources.GroupRolesPatchRequestBody{}
	v.validateRequestBody(body, w, r, func(w http.ResponseWriter, r *http.Request) {
		v.ServerInterface.PatchGroupsItemRoles(w, r, id)
	})
}
