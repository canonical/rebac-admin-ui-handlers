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

// CapabilitiesService implements the `CapabilitiesService` interface.
type CapabilitiesService struct {
	Database *database.Database
}

// For doc/test sake, to hint that the struct needs to implement a specific interface.
var _ interfaces.CapabilitiesService = &CapabilitiesService{}

// ListCapabilities returns a list of capabilities supported by this service.
func (s *CapabilitiesService) ListCapabilities(ctx context.Context) ([]resources.Capability, error) {
	// Note that, normally, capabilities are known (as static information). But,
	// here, to avoid hardcoding things and be able to modify the API with
	// different responses, we need to fetch the data from database.
	return s.Database.ListCapabilities(), nil
}
