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
	"net/http"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// getAuthenticationMiddleware returns a middleware function that delegates the
// extraction of the caller identity to the provided authenticator backend, and
// store the returned identity in the request context.
// If no authenticator backend is provided, a no-op middleware is returned.
func (b *ReBACAdminBackend) authenticationMiddleware() resources.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if b.params.Authenticator == nil {
				// This should never happen, because the outmost constructor does not
				// allow nil for authenticator. But it's possible to miss this requirement
				// in manually created instances (like in tests), we should do the checking.
				writeErrorResponse(w, NewUnknownError("missing authenticator"))
				return
			}

			identity, err := b.params.Authenticator.Authenticate(r)
			if err != nil {
				writeServiceErrorResponse(w, b.params.AuthenticatorErrorMapper, err)
				return
			}
			if identity == nil {
				writeErrorResponse(w, NewAuthenticationError("nil identity"))
				return
			}
			next.ServeHTTP(w, newRequestWithIdentityInContext(r, identity))
		})
	}
}

type authenticatedIdentityContextKey struct{}

// GetIdentityFromContext fetches authenticated identity of the caller from the
// given request context. If the value was not found in the given context, this
// will return an error.
//
// The function is intended to be used by service backends.
func GetIdentityFromContext(ctx context.Context) (any, error) {
	identity := ctx.Value(authenticatedIdentityContextKey{})
	if identity == nil {
		return nil, NewAuthenticationError("missing caller identity")
	}
	return identity, nil
}

// newRequestWithIdentityInContext sets the given authenticated identity in a
// new request instance context and returns the new request.
func newRequestWithIdentityInContext(r *http.Request, identity any) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), authenticatedIdentityContextKey{}, identity))
}
