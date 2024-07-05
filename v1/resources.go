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

// GetResources returns the list of known resources.
// (GET /resources)
func (h handler) GetResources(w http.ResponseWriter, req *http.Request, params resources.GetResourcesParams) {
	ctx := req.Context()

	res, err := h.Resources.ListResources(ctx, &params)
	if err != nil {
		writeServiceErrorResponse(w, h.ResourcesErrorMapper, err)
		return
	}

	response := resources.GetResourcesResponse{
		Links:  resources.NewResponseLinks[resources.Resource](req.URL, res),
		Meta:   res.Meta,
		Data:   res.Data,
		Status: http.StatusOK,
	}

	writeResponse(w, http.StatusOK, response)
}
