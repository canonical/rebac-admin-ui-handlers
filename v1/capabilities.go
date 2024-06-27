// Copyright 2024 Canonical Ltd.
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"net/http"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// GetCapabilities returns the list of endpoints implemented by this API.
// (GET /capabilities)
func (h handler) GetCapabilities(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	var capabilities []resources.Capability
	var err error
	if h.Capabilities != nil {
		capabilities, err = h.Capabilities.ListCapabilities(ctx)
		if err != nil {
			writeServiceErrorResponse(w, h.CapabilitiesErrorMapper, err)
			return
		}
	} else {
		capabilities = h.inferCapabilities()
	}

	response := resources.GetCapabilitiesResponse{
		Meta: resources.ResponseMeta{
			Size: len(capabilities),
		},
		Data:   capabilities,
		Status: http.StatusOK,
	}

	writeResponse(w, http.StatusOK, response)
}

// inferCapabilities infers the handler capabilities based on the provided
// service backends.
func (h handler) inferCapabilities() []resources.Capability {
	result := []resources.Capability{
		{Endpoint: "/swagger.json", Methods: []resources.CapabilityMethods{"GET"}},
		{Endpoint: "/capabilities", Methods: []resources.CapabilityMethods{"GET"}},
	}

	if h.IdentityProviders != nil {
		result = append(result, []resources.Capability{
			{Endpoint: "/authentication/providers", Methods: []resources.CapabilityMethods{"GET"}},
			{Endpoint: "/authentication", Methods: []resources.CapabilityMethods{"GET", "POST"}},
			{Endpoint: "/authentication/{id}", Methods: []resources.CapabilityMethods{"GET", "PUT", "DELETE"}},
		}...)
	}

	if h.Identities != nil {
		result = append(result, []resources.Capability{
			{Endpoint: "/identities", Methods: []resources.CapabilityMethods{"GET", "POST"}},
			{Endpoint: "/identities/{id}", Methods: []resources.CapabilityMethods{"GET", "PUT", "DELETE"}},
			{Endpoint: "/identities/{id}/groups", Methods: []resources.CapabilityMethods{"GET", "PATCH"}},
			{Endpoint: "/identities/{id}/roles", Methods: []resources.CapabilityMethods{"GET", "PATCH"}},
			{Endpoint: "/identities/{id}/entitlements", Methods: []resources.CapabilityMethods{"GET", "PATCH"}},
		}...)
	}

	if h.Groups != nil {
		result = append(result, []resources.Capability{
			{Endpoint: "/groups", Methods: []resources.CapabilityMethods{"GET", "POST"}},
			{Endpoint: "/groups/{id}", Methods: []resources.CapabilityMethods{"GET", "PUT", "DELETE"}},
			{Endpoint: "/groups/{id}/identities", Methods: []resources.CapabilityMethods{"GET", "PATCH"}},
			{Endpoint: "/groups/{id}/roles", Methods: []resources.CapabilityMethods{"GET", "PATCH"}},
			{Endpoint: "/groups/{id}/entitlements", Methods: []resources.CapabilityMethods{"GET", "PATCH"}},
		}...)
	}

	if h.Roles != nil {
		result = append(result, []resources.Capability{
			{Endpoint: "/roles", Methods: []resources.CapabilityMethods{"GET", "POST"}},
			{Endpoint: "/roles/{id}", Methods: []resources.CapabilityMethods{"GET", "PUT", "DELETE"}},
			{Endpoint: "/roles/{id}/entitlements", Methods: []resources.CapabilityMethods{"GET", "PATCH"}},
		}...)
	}

	if h.Entitlements != nil {
		result = append(result, []resources.Capability{
			{Endpoint: "/entitlements", Methods: []resources.CapabilityMethods{"GET"}},
			{Endpoint: "/entitlements/raw", Methods: []resources.CapabilityMethods{"GET"}},
		}...)
	}

	if h.Resources != nil {
		result = append(result, []resources.Capability{
			{Endpoint: "/resources", Methods: []resources.CapabilityMethods{"GET"}},
		}...)
	}

	return result
}
