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
	"net/http/httptest"
)

// newTestRequest returns a new HTTP request instance with the given body set in the
// corresponding context.
func newTestRequest[T any](method, path string, body *T) *http.Request {
	r := httptest.NewRequest(method, path, nil)
	return newRequestWithBodyInContext(r, body)
}

// stringPtr is a helper function that returns a pointer to the given string literal.
func stringPtr(s string) *string {
	return &s
}
