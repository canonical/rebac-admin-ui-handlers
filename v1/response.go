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
	"encoding/json"
	"net/http"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// writeErrorResponse writes the given err in the response with format defined
// by the OpenAPI spec.
func writeErrorResponse(w http.ResponseWriter, err error) {
	resp := mapErrorResponse(err)

	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("unexpected marshalling error")); err != nil {
			// TODO(CSS-7642): we should log the error.
			return
		}
		return
	}

	setJSONContentTypeHeader(w)
	w.WriteHeader(int(resp.Status))
	if _, err := w.Write(body); err != nil {
		// TODO(CSS-7642): we should log the error.
		return
	}
}

// mapErrorResponse returns a Response instance filled with the given error.
func mapErrorResponse(err error) *resources.Response {
	var asErrorWithStatus *errorWithStatus
	if err == nil {
		// Theoretically, this should never happen, but we anyway have to check for
		// a nil argument.
		asErrorWithStatus = &errorWithStatus{status: http.StatusOK}
	} else if e, ok := err.(*errorWithStatus); ok {
		asErrorWithStatus = e
	} else if e := mapHandlerBadRequestError(err); e != nil {
		asErrorWithStatus = e
	} else {
		asErrorWithStatus = &errorWithStatus{
			status:  http.StatusInternalServerError,
			message: err.Error(),
		}
	}

	return &resources.Response{
		Message: asErrorWithStatus.Error(),
		Status:  asErrorWithStatus.status,
	}
}

// writeResponse is a helper method to avoid verbose repetition of very common instructions
func writeResponse(w http.ResponseWriter, status int, responseObject interface{}) {
	data, err := json.Marshal(responseObject)
	if err != nil {
		writeErrorResponse(w, err)
		return
	}

	setJSONContentTypeHeader(w)
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		// TODO(CSS-7642): we should log the error.
		return
	}
}

// mapServiceErrorResponse maps errors thrown by services to the designated
// response type. If the given mapper is nil, the method uses the default
// mapping strategy.
//
// This method should never return nil response.
func mapServiceErrorResponse(mapper ErrorResponseMapper, err error) *resources.Response {
	var response *resources.Response
	if mapper != nil {
		response = mapper.MapError(err)
	}

	if response == nil {
		response = mapErrorResponse(err)
	}
	return response
}

// writeServiceErrorResponse is a helper method that maps errors thrown by
// services and writes them to the HTTP response stream.
func writeServiceErrorResponse(w http.ResponseWriter, mapper ErrorResponseMapper, err error) {
	response := mapServiceErrorResponse(mapper, err)
	writeResponse(w, response.Status, response)
}

func setJSONContentTypeHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
