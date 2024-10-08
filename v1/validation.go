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
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// handlerWithValidation decorates a given handler with validation logic. The
// request body is parsed into a safely-typed value and passed to the handler
// via context.
type handlerWithValidation struct {
	// Wrapped/decorated handler
	resources.ServerInterface

	validate *validator.Validate
}

// newHandlerWithValidation returns a new instance of the validationHandlerDecorator struct.
func newHandlerWithValidation(handler resources.ServerInterface) *handlerWithValidation {
	return &handlerWithValidation{
		ServerInterface: handler,
		validate:        validator.New(),
	}
}

// requestBodyContextKey is the context key to retrieve the parsed request body struct instance.
type requestBodyContextKey struct{}

// getRequestBodyFromContext fetches request body from given context. If the value
// was not found in the given context, this will return an error.
func getRequestBodyFromContext(ctx context.Context) (any, error) {
	body := ctx.Value(requestBodyContextKey{})
	if body == nil {
		return nil, NewMissingRequestBodyError("request body is not available")
	}
	return body, nil
}

// newRequestWithBodyInContext sets the given body in a new request instance context
// and returns the new request.
func newRequestWithBodyInContext(r *http.Request, body any) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), requestBodyContextKey{}, body))
}

// parseRequestBody parses request body as JSON and populates the given body instance.
func parseRequestBody(body any, r *http.Request) error {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		return NewMissingRequestBodyError("request body is not a valid JSON")
	}
	return nil
}

// validateRequestBody is a helper method to avoid repetition. It parses
// request body, validates it against the given validator instance and if it's
// okay, will delegate to the provided callback with a new HTTP request instance
// with the parse body in the context.
func (v handlerWithValidation) validateRequestBody(body any, w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request)) {
	err := parseRequestBody(body, r)
	if err != nil {
		writeErrorResponse(w, err)
		return
	}
	if err := v.validate.Struct(body); err != nil {
		writeErrorResponse(w, NewRequestBodyValidationError(err.Error()))
		return
	}
	f(w, newRequestWithBodyInContext(r, body))
}
