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

type handlerDispatcherParams struct {
	ImplementsGroups            bool
	ImplementsIdentities        bool
	ImplementsIdentityProviders bool
	ImplementsRoles             bool
	ImplementsResources         bool
	ImplementsEntitlements      bool
}

type handlerDispatcher struct {
	// Wrapped/decorated handler
	resources.ServerInterface

	params handlerDispatcherParams
}

// newHandlerDispatcher returns a new instance of the newHandlerDispatcher struct.
func newHandlerDispatcher(handler resources.ServerInterface, params handlerDispatcherParams) *handlerDispatcher {
	return &handlerDispatcher{
		ServerInterface: handler,
		params:          params,
	}
}

// GetCapabilities delegates the call to the wrapped handler's `GetCapabilities` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetCapabilities(w http.ResponseWriter, r *http.Request) {
	// This endpoint should always work, regardless of any user implementations provided.
	h.ServerInterface.GetCapabilities(w, r)
}

// SwaggerJson delegates the call to the wrapped handler's `SwaggerJson` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) SwaggerJson(w http.ResponseWriter, r *http.Request) {
	// This endpoint is not handled by the user implementation.
	h.ServerInterface.SwaggerJson(w, r)
}

// GetIdentityProviders delegates the call to the wrapped handler's `GetIdentityProviders` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetIdentityProviders(w http.ResponseWriter, r *http.Request, params resources.GetIdentityProvidersParams) {
	if !h.params.ImplementsIdentityProviders {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetIdentityProviders(w, r, params)
}

// PostIdentityProviders delegates the call to the wrapped handler's `PostIdentityProviders` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) PostIdentityProviders(w http.ResponseWriter, r *http.Request) {
	if !h.params.ImplementsIdentityProviders {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.PostIdentityProviders(w, r)
}

// GetAvailableIdentityProviders delegates the call to the wrapped handler's `GetAvailableIdentityProviders` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetAvailableIdentityProviders(w http.ResponseWriter, r *http.Request, params resources.GetAvailableIdentityProvidersParams) {
	if !h.params.ImplementsIdentityProviders {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetAvailableIdentityProviders(w, r, params)
}

// DeleteIdentityProvidersItem delegates the call to the wrapped handler's `DeleteIdentityProvidersItem` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) DeleteIdentityProvidersItem(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsIdentityProviders {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.DeleteIdentityProvidersItem(w, r, id)
}

// GetIdentityProvidersItem delegates the call to the wrapped handler's `GetIdentityProvidersItem` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetIdentityProvidersItem(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsIdentityProviders {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetIdentityProvidersItem(w, r, id)
}

// PutIdentityProvidersItem delegates the call to the wrapped handler's `PutIdentityProvidersItem` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) PutIdentityProvidersItem(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsIdentityProviders {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.PutIdentityProvidersItem(w, r, id)
}

// GetEntitlements delegates the call to the wrapped handler's `GetEntitlements` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetEntitlements(w http.ResponseWriter, r *http.Request, params resources.GetEntitlementsParams) {
	if !h.params.ImplementsEntitlements {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetEntitlements(w, r, params)
}

// GetRawEntitlements delegates the call to the wrapped handler's `GetRawEntitlements` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetRawEntitlements(w http.ResponseWriter, r *http.Request) {
	if !h.params.ImplementsEntitlements {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetRawEntitlements(w, r)
}

// GetGroups delegates the call to the wrapped handler's `GetGroups` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetGroups(w http.ResponseWriter, r *http.Request, params resources.GetGroupsParams) {
	if !h.params.ImplementsGroups {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetGroups(w, r, params)
}

// PostGroups delegates the call to the wrapped handler's `PostGroups` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) PostGroups(w http.ResponseWriter, r *http.Request) {
	if !h.params.ImplementsGroups {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.PostGroups(w, r)
}

// DeleteGroupsItem delegates the call to the wrapped handler's `DeleteGroupsItem` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) DeleteGroupsItem(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsGroups {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.DeleteGroupsItem(w, r, id)
}

// GetGroupsItem delegates the call to the wrapped handler's `GetGroupsItem` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetGroupsItem(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsGroups {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetGroupsItem(w, r, id)
}

// PutGroupsItem delegates the call to the wrapped handler's `PutGroupsItem` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) PutGroupsItem(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsGroups {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.PutGroupsItem(w, r, id)
}

// GetGroupsItemEntitlements delegates the call to the wrapped handler's `GetGroupsItemEntitlements` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetGroupsItemEntitlements(w http.ResponseWriter, r *http.Request, id string, params resources.GetGroupsItemEntitlementsParams) {
	if !h.params.ImplementsGroups {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetGroupsItemEntitlements(w, r, id, params)
}

// PatchGroupsItemEntitlements delegates the call to the wrapped handler's `PatchGroupsItemEntitlements` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) PatchGroupsItemEntitlements(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsGroups {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.PatchGroupsItemEntitlements(w, r, id)
}

// GetGroupsItemIdentities delegates the call to the wrapped handler's `GetGroupsItemIdentities` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetGroupsItemIdentities(w http.ResponseWriter, r *http.Request, id string, params resources.GetGroupsItemIdentitiesParams) {
	if !h.params.ImplementsGroups {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetGroupsItemIdentities(w, r, id, params)
}

// PatchGroupsItemIdentities delegates the call to the wrapped handler's `PatchGroupsItemIdentities` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) PatchGroupsItemIdentities(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsGroups {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.PatchGroupsItemIdentities(w, r, id)
}

// GetGroupsItemRoles delegates the call to the wrapped handler's `GetGroupsItemRoles` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetGroupsItemRoles(w http.ResponseWriter, r *http.Request, id string, params resources.GetGroupsItemRolesParams) {
	if !h.params.ImplementsGroups {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetGroupsItemRoles(w, r, id, params)
}

// PatchGroupsItemRoles delegates the call to the wrapped handler's `PatchGroupsItemRoles` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) PatchGroupsItemRoles(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsGroups {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.PatchGroupsItemRoles(w, r, id)
}

// GetIdentities delegates the call to the wrapped handler's `GetIdentities` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetIdentities(w http.ResponseWriter, r *http.Request, params resources.GetIdentitiesParams) {
	if !h.params.ImplementsIdentities {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetIdentities(w, r, params)
}

// PostIdentities delegates the call to the wrapped handler's `PostIdentities` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) PostIdentities(w http.ResponseWriter, r *http.Request) {
	if !h.params.ImplementsIdentities {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.PostIdentities(w, r)
}

// DeleteIdentitiesItem delegates the call to the wrapped handler's `DeleteIdentitiesItem` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) DeleteIdentitiesItem(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsIdentities {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.DeleteIdentitiesItem(w, r, id)
}

// GetIdentitiesItem delegates the call to the wrapped handler's `GetIdentitiesItem` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetIdentitiesItem(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsIdentities {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetIdentitiesItem(w, r, id)
}

// PutIdentitiesItem delegates the call to the wrapped handler's `PutIdentitiesItem` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) PutIdentitiesItem(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsIdentities {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.PutIdentitiesItem(w, r, id)
}

// GetIdentitiesItemEntitlements delegates the call to the wrapped handler's `GetIdentitiesItemEntitlements` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetIdentitiesItemEntitlements(w http.ResponseWriter, r *http.Request, id string, params resources.GetIdentitiesItemEntitlementsParams) {
	if !h.params.ImplementsIdentities {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetIdentitiesItemEntitlements(w, r, id, params)
}

// PatchIdentitiesItemEntitlements delegates the call to the wrapped handler's `PatchIdentitiesItemEntitlements` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) PatchIdentitiesItemEntitlements(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsIdentities {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.PatchIdentitiesItemEntitlements(w, r, id)
}

// GetIdentitiesItemGroups delegates the call to the wrapped handler's `GetIdentitiesItemGroups` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetIdentitiesItemGroups(w http.ResponseWriter, r *http.Request, id string, params resources.GetIdentitiesItemGroupsParams) {
	if !h.params.ImplementsIdentities {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetIdentitiesItemGroups(w, r, id, params)
}

// PatchIdentitiesItemGroups delegates the call to the wrapped handler's `PatchIdentitiesItemGroups` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) PatchIdentitiesItemGroups(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsIdentities {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.PatchIdentitiesItemGroups(w, r, id)
}

// GetIdentitiesItemRoles delegates the call to the wrapped handler's `GetIdentitiesItemRoles` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetIdentitiesItemRoles(w http.ResponseWriter, r *http.Request, id string, params resources.GetIdentitiesItemRolesParams) {
	if !h.params.ImplementsIdentities {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetIdentitiesItemRoles(w, r, id, params)
}

// PatchIdentitiesItemRoles delegates the call to the wrapped handler's `PatchIdentitiesItemRoles` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) PatchIdentitiesItemRoles(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsIdentities {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.PatchIdentitiesItemRoles(w, r, id)
}

// GetResources delegates the call to the wrapped handler's `GetResources` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetResources(w http.ResponseWriter, r *http.Request, params resources.GetResourcesParams) {
	if !h.params.ImplementsResources {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetResources(w, r, params)
}

// GetRoles delegates the call to the wrapped handler's `GetRoles` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetRoles(w http.ResponseWriter, r *http.Request, params resources.GetRolesParams) {
	if !h.params.ImplementsRoles {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetRoles(w, r, params)
}

// PostRoles delegates the call to the wrapped handler's `PostRoles` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) PostRoles(w http.ResponseWriter, r *http.Request) {
	if !h.params.ImplementsRoles {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.PostRoles(w, r)
}

// DeleteRolesItem delegates the call to the wrapped handler's `DeleteRolesItem` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) DeleteRolesItem(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsRoles {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.DeleteRolesItem(w, r, id)
}

// GetRolesItem delegates the call to the wrapped handler's `GetRolesItem` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetRolesItem(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsRoles {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetRolesItem(w, r, id)
}

// PutRolesItem delegates the call to the wrapped handler's `PutRolesItem` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) PutRolesItem(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsRoles {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.PutRolesItem(w, r, id)
}

// GetRolesItemEntitlements delegates the call to the wrapped handler's `GetRolesItemEntitlements` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) GetRolesItemEntitlements(w http.ResponseWriter, r *http.Request, id string, params resources.GetRolesItemEntitlementsParams) {
	if !h.params.ImplementsRoles {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.GetRolesItemEntitlements(w, r, id, params)
}

// PatchRolesItemEntitlements delegates the call to the wrapped handler's `PatchRolesItemEntitlements` method, if it is allowed; otherwise returns a `501 Unimplemented` status code.
func (h handlerDispatcher) PatchRolesItemEntitlements(w http.ResponseWriter, r *http.Request, id string) {
	if !h.params.ImplementsRoles {
		writeErrorResponse(w, NewNotImplementedError(""))
		return
	}
	h.ServerInterface.PatchRolesItemEntitlements(w, r, id)
}
