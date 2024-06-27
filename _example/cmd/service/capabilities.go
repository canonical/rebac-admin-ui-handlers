// Copyright 2024 Canonical Ltd.

package service

import (
	"context"
	"example/cmd/database"

	"github.com/canonical/rebac-admin-ui-handlers/v1/interfaces"
	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// CapabilitiesService implements the `CapabilitiesService` interface.
type CapabilitiesService struct {
	Database *database.Database
}

// For doc/test sake, to hint that the struct needs to implement a specific interface.
var _ interfaces.CapabilitiesService = &CapabilitiesService{}

// ListCapabilities returns a list of capabilities supported by this service.
func (s *CapabilitiesService) ListCapabilities(ctx context.Context) ([]resources.Capability, error) {
	// Note that, normally, capabilities are known (as static information). But,
	// here, to avoid hardcoding things and be able to modify the API with
	// different responses, we need to fetch the data from database.
	return s.Database.ListCapabilities(), nil
}
