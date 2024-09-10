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

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// unimplementedEntitlementsService represents a *Null Object* implementation of the the `EntitlementsService` interface.
type unimplementedEntitlementsService struct{}

func (s unimplementedEntitlementsService) ListEntitlements(ctx context.Context, params *resources.GetEntitlementsParams) ([]resources.EntitlementSchema, error) {
	return nil, NewNotImplementedError("")
}

func (s unimplementedEntitlementsService) RawEntitlements(ctx context.Context) (string, error) {
	return "", NewNotImplementedError("")
}
