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

package database

import (
	"fmt"
	"strings"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// EntitlementToString marshals a given entitlement as a string.
func EntitlementToString(e resources.EntityEntitlement) string {
	// For example: "can_read::controller:foo"
	return fmt.Sprintf("%s::%s:%s", e.Entitlement, e.EntityType, e.EntityId)
}

// EntitlementFromString unmarshals a entitlement from a given string.
func EntitlementFromString(s string) resources.EntityEntitlement {
	parts := strings.SplitN(s, ":", 4)
	result := resources.EntityEntitlement{}
	if len(parts) > 0 {
		result.Entitlement = parts[0]
	}
	if len(parts) > 1 {
		result.EntityType = parts[1]
	}
	if len(parts) > 2 {
		result.EntityId = parts[2]
	}
	return result
}

func getMapValues[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
