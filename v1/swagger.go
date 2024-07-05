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
	"log"
	"net/http"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// SwaggerJson Returns the OpenAPI spec as a JSON file.
// (GET /swagger.json)
func (h handler) SwaggerJson(w http.ResponseWriter, req *http.Request) {
	swagger, err := resources.GetSwagger()
	if err != nil {
		writeErrorResponse(w, NewUnknownError("cannot retrieve swagger data"))
		return
	}

	body, err := swagger.MarshalJSON()
	if err != nil {
		writeErrorResponse(w, NewUnknownError("cannot marshal spec as JSON"))
		return
	}

	if _, err := w.Write(body); err != nil {
		log.Printf("failed to write response body: %v", err)
	}
}
