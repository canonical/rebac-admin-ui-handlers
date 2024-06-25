package database

import (
	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

func (db *Database) GetAuthModel() string {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return db.AuthModel
}

func (db *Database) ListUserEntitlements() []resources.EntityEntitlement {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return db.UserEntitlements
}

func (db *Database) ListAvailableIdentityProviders() []resources.AvailableIdentityProvider {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return db.AvailableIdentityProviders
}

func (db *Database) ListUserResources() []resources.Resource {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return db.UserResources
}

func (db *Database) ListCapabilities() []resources.Capability {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	return db.Capabilities
}
