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
	"fmt"
	"net/http"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// errorWithStatus is an internal error representation that holds the corresponding
// HTTP status code along with the error message.
type errorWithStatus struct {
	// status is the HTTP standard equivalent status. Acceptable
	// values are `http.Status*` constants.
	status  int
	message string
}

// Error implements the error interface.
func (e *errorWithStatus) Error() string {
	statusText := http.StatusText(e.status)
	if statusText == "" {
		statusText = "[Unknown error]"
	}
	if e.message == "" {
		return statusText
	}
	return fmt.Sprintf("%s: %s", statusText, e.message)
}

// NewAuthenticationError returns an error instance that represents an authentication error.
func NewAuthenticationError(message string) error {
	return &errorWithStatus{
		status:  http.StatusUnauthorized,
		message: fmt.Sprintf("authentication failed: %s", message),
	}
}

// NewAuthorizationError returns an error instance that represents an unauthorized access error.
func NewAuthorizationError(message string) error {
	return &errorWithStatus{
		status:  http.StatusUnauthorized,
		message: fmt.Sprintf("authorization failed: %s", message),
	}
}

// NewNotFoundError returns an error instance that represents a not-found error.
func NewNotFoundError(message string) error {
	return &errorWithStatus{
		status:  http.StatusNotFound,
		message: message,
	}
}

// NewMissingRequestBodyError returns an error instance that represents a missing request body error.
func NewMissingRequestBodyError(message string) error {
	return &errorWithStatus{
		status:  http.StatusBadRequest,
		message: fmt.Sprintf("missing request body: %s", message),
	}
}

// NewValidationError returns an error instance that represents an input validation error.
func NewValidationError(message string) error {
	return &errorWithStatus{
		status:  http.StatusBadRequest,
		message: message,
	}
}

// NewRequestBodyValidationError returns an error instance that represents a request body validation error.
func NewRequestBodyValidationError(message string) error {
	return &errorWithStatus{
		status:  http.StatusBadRequest,
		message: fmt.Sprintf("invalid request body: %s", message),
	}
}

// NewInvalidRequestError returns an error instance that represents a problem with the input (e.g., when trying to add
// an entry which already exists).
func NewInvalidRequestError(message string) error {
	return &errorWithStatus{
		status:  http.StatusBadRequest,
		message: fmt.Sprintf("invalid request: %s", message),
	}
}

// NewNotImplementedError returns an error instance that reports the requested operation is not implemented by the
// backend.
func NewNotImplementedError(message string) error {
	return &errorWithStatus{
		status:  http.StatusNotImplemented,
		message: fmt.Sprintf("not implemented: %s", message),
	}
}

// NewUnknownError returns an error instance that represents an unknown internal error.
func NewUnknownError(message string) error {
	return &errorWithStatus{
		status:  http.StatusInternalServerError,
		message: message,
	}
}

// ErrorResponseMapper is the basic interface to allow for error -> http response mapping
type ErrorResponseMapper interface {
	// MapError maps an error into a Response. If the method is unable to map the
	// error (e.g., the error is unknown), it must return nil.
	MapError(error) *resources.Response
}

// mapHandlerBadRequestError checks if the given error is an "Bad Request" error
// thrown at the handler root (i.e., an auto-generated error type) and return the
// equivalent errorWithStatus instance. If the given error is not an internal
// handler error, this function will return nil.
func mapHandlerBadRequestError(err error) *errorWithStatus {
	if !isHandlerBadRequestError(err) {
		return nil
	}
	return &errorWithStatus{
		status:  http.StatusBadRequest,
		message: err.Error(),
	}
}

func isHandlerBadRequestError(err error) bool {
	switch err.(type) {
	case *resources.UnmarshalingParamError:
		return true
	case *resources.RequiredParamError:
		return true
	case *resources.RequiredHeaderError:
		return true
	case *resources.InvalidParamFormatError:
		return true
	case *resources.TooManyValuesForParamError:
		return true
	}
	return false
}
