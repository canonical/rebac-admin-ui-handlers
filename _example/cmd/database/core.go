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
	isDirty bool

	stateFilename     string
	zeroStateFilename string

	persistMutex sync.Mutex

	mutex sync.RWMutex

	Groups     map[string]resources.Group
	Identities map[string]resources.Identity
	Roles      map[string]resources.Role
	Idps       map[string]resources.IdentityProvider

	Group2Identity       *Relationship
	Group2Role           *Relationship
	Group2Entitlement    *Relationship
	Identity2Role        *Relationship
	Identity2Entitlement *Relationship
	Role2Entitlement     *Relationship

	// Constant data
	UserEntitlements           []resources.EntityEntitlement
	AvailableIdentityProviders []resources.AvailableIdentityProvider
	UserResources              []resources.Resource
	AuthModel                  string
	Capabilities               []resources.Capability
}

// NewDatabase creates a new in-memory database instance. The `stateFilename`
// argument should point to a JSON file that will be used to persist/load data.
// The `zeroStateFilename` should to a zero state file to be used when resetting
// the state (or when there the `stateFilename` does not exist).
func NewDatabase(stateFilename string, zeroStateFilename string) *Database {
	result := &Database{
		stateFilename:     stateFilename,
		zeroStateFilename: zeroStateFilename,
	}
	cleanupDatabase(result)
	return result
}

func cleanupDatabase(db *Database) {
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

// Load populates the state from the source.
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
	cleanupDatabase(db)
	if err := json.Unmarshal(raw, &db); err != nil {
		return fmt.Errorf("failed to unmarshal state data: %w", err)
	}
	db.isDirty = false
	return nil
}

// Persist writes the state to the source file, if there is any unpersisted changes.
func (db *Database) Persist() error {
	db.persistMutex.Lock()
	defer db.persistMutex.Unlock()
	if !db.isDirty {
		return nil
	}
	err := db.persist()
	if err != nil {
		return err
	}
	db.isDirty = false
	return nil
}

func (db *Database) persist() error {
	raw, _ := json.MarshalIndent(db, "", "    ")
	if err := os.WriteFile(db.stateFilename, raw, os.ModePerm); err != nil {
		return fmt.Errorf("failed to persist state: %w", err)
	}
	return nil
}

// Reset cleans up database and loads the zero-state data.
func (db *Database) Reset() error {
	db.persistMutex.Lock()
	defer db.persistMutex.Unlock()

	db.load(db.zeroStateFilename)
	return db.persist()
}
