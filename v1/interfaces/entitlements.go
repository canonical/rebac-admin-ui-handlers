// Copyright 2024 Canonical Ltd.
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

type EntitlementsService interface {
	// ListEntitlements returns the list of entitlements in JSON format.
	ListEntitlements(ctx context.Context, params *resources.GetEntitlementsParams) (*resources.PaginatedResponse[resources.EntityEntitlement], error)

	// RawEntitlements returns the list of entitlements as raw text.
	RawEntitlements(ctx context.Context) (string, error)
}
