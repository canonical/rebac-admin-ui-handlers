package database

import (
	"context"
	"errors"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

func (db *Database) ListIdentityProviders() []resources.IdentityProvider {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	return GetMapValues(db.Idps)
}

func (db *Database) AddIdentityProvider(idp *resources.IdentityProvider) (*resources.IdentityProvider, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

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
	db.Persist()
	return &entry, nil
}

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

func (db *Database) UpdateIdentityProvider(ctx context.Context, idp *resources.IdentityProvider) *resources.IdentityProvider {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	if idp.Id == nil {
		return nil
	}

	if _, ok := db.Idps[*idp.Id]; !ok {
		return nil
	}

	db.Idps[*idp.Id] = *idp
	db.Persist()
	return idp
}

func (db *Database) DeleteIdentityProvider(id string) bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Load()

	if _, ok := db.Idps[id]; !ok {
		return false
	}
	defer db.Persist()
	delete(db.Idps, id)
	db.Persist()
	return true
}
