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

// GetEntitlements returns the list of known entitlements in JSON format.
// (GET /entitlements)
func (h handler) GetEntitlements(w http.ResponseWriter, req *http.Request, params resources.GetEntitlementsParams) {
	ctx := req.Context()

	entitlements, err := h.Entitlements.ListEntitlements(ctx, &params)
	if err != nil {
		writeServiceErrorResponse(w, h.EntitlementsErrorMapper, err)
		return
	}

	response := resources.GetEntitlementsResponse{
		Links:  resources.NewResponseLinks[resources.EntityEntitlement](req.URL, entitlements),
		Meta:   entitlements.Meta,
		Data:   entitlements.Data,
		Status: http.StatusOK,
	}

	writeResponse(w, http.StatusOK, response)

}

// GetRawEntitlements returns the list of known entitlements as raw text.
// (GET /entitlements/raw)
func (h handler) GetRawEntitlements(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	entitlementsRawString, err := h.Entitlements.RawEntitlements(ctx)
	if err != nil {
		writeServiceErrorResponse(w, h.EntitlementsErrorMapper, err)
		return
	}

	w.Header().Add("Content-Type", "text/plain")
	writeResponse(w, http.StatusOK, entitlementsRawString)

}
