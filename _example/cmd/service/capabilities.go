package service

import (
	"context"
	"example/cmd/database"

	"github.com/canonical/rebac-admin-ui-handlers/v1/interfaces"
	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

type CapabilitiesService struct {
	Database *database.Database
}

// For doc/test sake, to hint that the struct needs to implement a specific interface.
var _ interfaces.CapabilitiesService = &CapabilitiesService{}

func (s *CapabilitiesService) ListCapabilities(ctx context.Context) ([]resources.Capability, error) {
	return s.Database.ListCapabilities(), nil
}
