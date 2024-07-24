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

import (
	"context"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// EntitlementsService defines an abstract backend to handle entitlement schema related operations.
type EntitlementsService interface {
	// ListEntitlements returns the list of entitlements in JSON format.
	ListEntitlements(ctx context.Context, params *resources.GetEntitlementsParams) ([]resources.EntityEntitlement, error)

	// RawEntitlements returns the list of entitlements as raw text.
	RawEntitlements(ctx context.Context) (string, error)
}
