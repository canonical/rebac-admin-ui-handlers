package database

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// Database represents an in-memory relational storage which keeps track of
// various kinds of entities.
type Database struct {
	stateFilename     string
	zeroStateFilename string

	persistMutex sync.Mutex

	mutex sync.RWMutex

	Groups     map[string]resources.Group
	Identities map[string]resources.Identity
	Roles      map[string]resources.Role
	Idps       map[string]resources.IdentityProvider

	Group2Identity       *relationship
	Group2Role           *relationship
	Group2Entitlement    *relationship
	Identity2Role        *relationship
	Identity2Entitlement *relationship
	Role2Entitlement     *relationship

	// Constant data
	UserEntitlements           []resources.EntityEntitlement
	AvailableIdentityProviders []resources.AvailableIdentityProvider
	UserResources              []resources.Resource
	AuthModel                  string
	Capabilities               []resources.Capability
}

func NewDatabase(stateFilename string, zeroStateFilename string) *Database {
	result := &Database{
		stateFilename:     stateFilename,
		zeroStateFilename: zeroStateFilename,
	}
	CleanupDatabase(result)
	return result
}

func CleanupDatabase(db *Database) {
	db.Groups = map[string]resources.Group{}
	db.Identities = map[string]resources.Identity{}
	db.Roles = map[string]resources.Role{}
	db.Idps = map[string]resources.IdentityProvider{}

	db.Group2Identity = NewRelationship()
	db.Group2Role = NewRelationship()
	db.Group2Entitlement = NewRelationship()
	db.Identity2Role = NewRelationship()
	db.Identity2Entitlement = NewRelationship()
	db.Role2Entitlement = NewRelationship()

	db.UserEntitlements = []resources.EntityEntitlement{}
	db.AvailableIdentityProviders = []resources.AvailableIdentityProvider{}
	db.UserResources = []resources.Resource{}
	db.AuthModel = ""
	db.Capabilities = []resources.Capability{}
}

func (db *Database) Load() error {
	db.persistMutex.Lock()
	defer db.persistMutex.Unlock()

	if _, err := os.Stat(db.stateFilename); err != nil {
		return db.load(db.zeroStateFilename)
	}
	return db.load(db.stateFilename)
}

func (db *Database) load(filename string) error {
	var raw []byte
	if _, err := os.Stat(filename); err != nil {
		raw, _ = json.Marshal(NewDatabase("", ""))
	} else {
		raw, err = os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("failed to read state file %s: %w", filename, err)
		}
	}
	CleanupDatabase(db)
	if err := json.Unmarshal(raw, &db); err != nil {
		return fmt.Errorf("failed to unmarshal state data: %w", err)
	}
	return nil
}

func (db *Database) Persist() error {
	db.persistMutex.Lock()
	defer db.persistMutex.Unlock()
	return db.persist()
}

func (db *Database) persist() error {
	raw, _ := json.MarshalIndent(db, "", "    ")
	if err := os.WriteFile(db.stateFilename, raw, os.ModePerm); err != nil {
		return fmt.Errorf("failed to persist state: %w", err)
	}
	return nil
}

func (db *Database) Reset() error {
	db.persistMutex.Lock()
	defer db.persistMutex.Unlock()

	db.load(db.zeroStateFilename)
	return db.persist()
}
