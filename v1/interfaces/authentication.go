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

package interfaces

import "net/http"

// Authenticator defines an abstract backend to perform authentication on HTTP requests.
type Authenticator interface {
	// Authenticate receives an HTTP request and returns the identity of the caller.
	// The same identity will be available to the service backend through the request
	// context. To avoid issues with value types, it's best to return a pointer type.
	//
	// Note that the implementations of this method should not alter the state of the
	// received request instance.
	//
	// If the returned identity is nil it will be regarded as authentication failure.
	//
	// To return an error, the implementations should use the provided error functions
	// (e.g., `NewAuthenticationError`) and avoid creating ad-hoc errors.
	Authenticate(r *http.Request) (any, error)
}
