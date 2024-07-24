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
	"strings"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// GetAuthModel returns the raw OpenFGA authorization model.
func (db *Database) GetAuthModel() string {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return db.AuthModel
}

// ListUserEntitlements returns the list of entitlements.
func (db *Database) ListUserEntitlements() []resources.EntitlementSchema {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return db.Entitlements
}

// ListAvailableIdentityProviders returns the list of available identity providers.
func (db *Database) ListAvailableIdentityProviders() []resources.AvailableIdentityProvider {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return db.AvailableIdentityProviders
}

// ListUserResources returns the list of resources.
// If `entityType` is nil then resource type filtering (exact match) is not applied.
// If `entityName` is nil then resource name filtering (prefix match) is not applied.
func (db *Database) ListUserResources(entityType *string, entityName *string) []resources.Resource {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	result := []resources.Resource{}
	for _, resource := range db.Resources {
		if entityType != nil && resource.Entity.Type != *entityType {
			continue
		}
		if entityName != nil && !strings.HasPrefix(resource.Entity.Name, *entityName) {
			continue
		}
		result = append(result, resource)
	}
	return result
}

// ListCapabilities returns the list of capabilities. Note that, normally,
// capabilities are known at the handler/service level (not at the database
// level as in here). But, here, to avoid hardcoding things and be able to
// modify the API with different responses, we need to fetch the data from
// database.
func (db *Database) ListCapabilities() []resources.Capability {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return db.Capabilities
}
