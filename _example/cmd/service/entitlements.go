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

package service

import (
	"context"
	"example/cmd/database"

	"github.com/canonical/rebac-admin-ui-handlers/v1/interfaces"
	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// EntitlementsService implements the `EntitlementsService` interface.
type EntitlementsService struct {
	Database *database.Database
}

// For doc/test sake, to hint that the struct needs to implement a specific interface.
var _ interfaces.EntitlementsService = &EntitlementsService{}

// ListEntitlements returns the list of entitlements in JSON format.
func (s *EntitlementsService) ListEntitlements(ctx context.Context, params *resources.GetEntitlementsParams) ([]resources.EntitlementSchema, error) {
	// For the sake of this example we allow everyone to call this method. If it's not
	// the case, you can do the following to get the user:
	//
	//    raw, _ := v1.GetIdentityFromContext(ctx)
	//    user, _ := raw.(*User)
	//
	// And return this error if the user is not authorized:
	//
	//    return nil, v1.NewAuthorizationError("user cannot add group")
	//

	return s.Database.ListUserEntitlements(), nil
}

// RawEntitlements returns the list of entitlements as raw text.
func (s *EntitlementsService) RawEntitlements(ctx context.Context) (string, error) {
	// This should return the raw OpenFGA authorization model, so we just return
	// a fake one:

	return s.Database.GetAuthModel(), nil
}
