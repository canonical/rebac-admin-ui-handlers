package main

import (
	"fmt"
	"net/http"
	"os"

	v1 "github.com/canonical/rebac-admin-ui-handlers/v1"

	"example/cmd/database"
	"example/cmd/service"
)

const stateFilename = "state.json"
const zeroStateFilename = "state.zero.json"

func main() {
	db := database.NewDatabase(stateFilename, zeroStateFilename)
	_ = db.Load()

	rebac, err := v1.NewReBACAdminBackend(v1.ReBACAdminBackendParams{
		Authenticator:     &service.HappyAuthenticator{},
		Capabilities:      &service.CapabilitiesService{Database: db},
		Groups:            &service.GroupsService{Database: db},
		Identities:        &service.IdentitiesService{Database: db},
		Roles:             &service.RolesService{Database: db},
		Entitlements:      &service.EntitlementsService{Database: db},
		Resources:         &service.ResourcesService{Database: db},
		IdentityProviders: &service.IdentityProvidersService{Database: db},
	})
	if err != nil {
		panic(err.Error())
	}

	mux := http.NewServeMux()

	// NOTE: When using the standard Go ServeMux, make sure you provide the same
	// base URL for both arguments below. The endpoints will be accessible via
	// `/rebac/v1/*` (note the `/v1/` prefix).
	//
	// For testing you can try:
	//    curl 0:9999/rebac/v1/swagger.json
	mux.Handle("/rebac/", rebac.Handler("/rebac/"))

	// NOTE: When using Chi, you should omit the base URL for the latter; like
	// this:
	//   mux := chi.NewMux()
	//   mux.Handle("/rebac/", rebac.Handler(""))

	// These endpoints are just for the sake of this in-memory server. So, you
	// don't need to implement them in your project.
	mux.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Reset(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
	exit := make(chan bool, 1)
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Shutting down\n"))
		exit <- true
	})

	go func() {
		<-exit
		os.Exit(0)
	}()

	fmt.Println("Running on :9999")
	err = http.ListenAndServe(":9999", mux)
	panic(err.Error())
}
