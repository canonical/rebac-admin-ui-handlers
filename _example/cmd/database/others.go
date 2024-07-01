// Copyright 2024 Canonical Ltd.

package database

import (
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
func (db *Database) ListUserEntitlements() []resources.EntityEntitlement {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return db.UserEntitlements
}

// ListAvailableIdentityProviders returns the list of available identity providers.
func (db *Database) ListAvailableIdentityProviders() []resources.AvailableIdentityProvider {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return db.AvailableIdentityProviders
}

// ListUserResources returns the list of resources.
func (db *Database) ListUserResources() []resources.Resource {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return db.Resources
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
