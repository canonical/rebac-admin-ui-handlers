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
	"github.com/canonical/rebac-admin-ui-handlers/v1/interfaces"
)

// handler is the innermost handler that calls the methods defined on the
// `*Service` interfaces.
type handler struct {
	Identities            interfaces.IdentitiesService
	IdentitiesErrorMapper ErrorResponseMapper

	Roles            interfaces.RolesService
	RolesErrorMapper ErrorResponseMapper

	IdentityProviders            interfaces.IdentityProvidersService
	IdentityProvidersErrorMapper ErrorResponseMapper

	Capabilities            interfaces.CapabilitiesService
	CapabilitiesErrorMapper ErrorResponseMapper

	Entitlements            interfaces.EntitlementsService
	EntitlementsErrorMapper ErrorResponseMapper

	Groups            interfaces.GroupsService
	GroupsErrorMapper ErrorResponseMapper

	Resources            interfaces.ResourcesService
	ResourcesErrorMapper ErrorResponseMapper
}
