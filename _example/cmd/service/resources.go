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

	"github.com/canonical/rebac-admin-ui-handlers/v1/interfaces"
	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"

	"example/cmd/database"
)

// ResourcesService implements the `ResourcesService` interface.
type ResourcesService struct {
	Database *database.Database
}

// For doc/test sake, to hint that the struct needs to implement a specific interface.
var _ interfaces.ResourcesService = &ResourcesService{}

// ResourcesService defines an abstract backend to handle Resources related operations.
func (s *ResourcesService) ListResources(ctx context.Context, params *resources.GetResourcesParams) (*resources.PaginatedResponse[resources.Resource], error) {
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

	return Paginate(s.Database.ListUserResources(params.EntityType), params.Size, params.Page, params.NextToken, params.NextPageToken, false)
}
