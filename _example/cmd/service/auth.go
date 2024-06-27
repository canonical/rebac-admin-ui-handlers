// Copyright 2024 Canonical Ltd.

package service

import (
	"net/http"

	"github.com/canonical/rebac-admin-ui-handlers/v1/interfaces"
)

// User is the data struct that we use to share the authenticated user with the
// API endpoint handlers.
type User struct {
	name string
}

// HappyAuthenticator implements a happy (all-granted) authenticator.
type HappyAuthenticator struct{}

// For doc/test sake, to hint that the struct needs to implement a specific interface.
var _ interfaces.Authenticator = &HappyAuthenticator{}

// Authenticate extracts the calling user information from the the given HTTP
// request. See the `Authenticator` interface for more.
func (a *HappyAuthenticator) Authenticate(r *http.Request) (any, error) {
	// This method is going to be called on every HTTP request. It can use the
	// provided HTTP request to perform user authentication (e.g., by checking
	// cookies or the bearer token).

	// If the authorization fails, you should just return like this (note the
	// usage of predefined custom errors):
	//   return nil, v1.NewAuthenticationError("some reason")

	// If there is some error unrelated to authentication (like a network or
	// database issue), you can just return like this:
	//   return nil, v1.NewUnknownError("some error")

	return &User{
		name: "john doe",
	}, nil
}
