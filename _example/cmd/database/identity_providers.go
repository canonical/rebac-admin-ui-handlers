package database

import (
	"context"
	"errors"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// ListIdentityProviders returns the list of identity providers.
func (db *Database) ListIdentityProviders() []resources.IdentityProvider {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	return getMapValues(db.Idps)
}

// AddIdentityProvider adds a new identity provider.
func (db *Database) AddIdentityProvider(idp *resources.IdentityProvider) (*resources.IdentityProvider, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	// Note that this is not a required field (as per the OpenAPI spec), but
	// since we want to use the same value as the ID, we have to enforce it to
	// be non-nil.
	if idp.Name == nil {
		return nil, errors.New("missing name")
	}

	// Names should be unique.
	_, ok := db.Idps[*idp.Name]
	if ok {
		return nil, errors.New("already exists")
	}

	id := *idp.Name
	entry := *idp
	entry.Id = &id
	db.Idps[id] = entry
	db.isDirty = true
	return &entry, nil
}

// GetIdentityProvider returns an identity provider identified by given ID. If
// nothing found, the method returns nil.
func (db *Database) GetIdentityProvider(id string) *resources.IdentityProvider {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	db.Load()

	result, ok := db.Idps[id]
	if !ok {
		return nil
	}
	return &result
}

// UpdateIdentityProvider updates an identity provider to the given value. If
// nothing found or the ID field is not set, the method returns nil.
func (db *Database) UpdateIdentityProvider(ctx context.Context, idp *resources.IdentityProvider) *resources.IdentityProvider {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	// Note that this field is optional (as per OpenAPI spec), but we need it to
	// be non-nil to perform the operation.
	if idp.Id == nil {
		return nil
	}

	if _, ok := db.Idps[*idp.Id]; !ok {
		return nil
	}

	db.Idps[*idp.Id] = *idp
	db.isDirty = true
	return idp
}

// DeleteIdentityProvider deletes an identity provider identified by given ID.
// If nothing found, the method returns false. On a successful deletion, the
// method returns true.
func (db *Database) DeleteIdentityProvider(id string) bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()
	defer db.Persist()

	if _, ok := db.Idps[id]; !ok {
		return false
	}
	delete(db.Idps, id)
	db.isDirty = true
	return true
}
