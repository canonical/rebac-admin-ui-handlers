// Copyright (C) 2024 Canonical Ltd.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package v1

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	qt "github.com/frankban/quicktest"
	"go.uber.org/mock/gomock"

	"github.com/canonical/rebac-admin-ui-handlers/v1/interfaces"
	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

//go:generate mockgen -package interfaces -destination ./interfaces/mock_capabilities.go -source=./interfaces/capabilities.go
//go:generate mockgen -package v1 -destination ./mock_error_response.go -source=./error.go

func TestHandler_Capabilities_GetCapabilitiesSuccess(t *testing.T) {
	c := qt.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCapabilitiesService := interfaces.NewMockCapabilitiesService(ctrl)

	mockCapabilities := []resources.Capability{
		{
			Endpoint: "/capabilities",
			Methods:  []resources.CapabilityMethods{"GET"},
		},
		{
			Endpoint: "/identities",
			Methods:  []resources.CapabilityMethods{"GET", "POST", "DELETE", "PUT"},
		},
		{
			Endpoint: "/roles",
			Methods:  []resources.CapabilityMethods{"GET", "POST", "DELETE", "PUT"},
		},
	}

	mockCapabilitiesService.EXPECT().ListCapabilities(gomock.Any()).Return(mockCapabilities, nil)

	expectedResponse := resources.GetCapabilitiesResponse{
		Meta: resources.ResponseMeta{
			Size: len(mockCapabilities),
		},
		Data:   mockCapabilities,
		Status: http.StatusOK,
	}

	mockWriter := httptest.NewRecorder()
	mockRequest := httptest.NewRequest(http.MethodGet, "/capabilities", nil)

	sut := handler{Capabilities: mockCapabilitiesService}
	sut.GetCapabilities(mockWriter, mockRequest)

	result := mockWriter.Result()
	defer result.Body.Close()

	responseBody, err := io.ReadAll(result.Body)
	c.Assert(err, qt.IsNil)

	c.Assert(err, qt.IsNil, qt.Commentf("Unexpected err while unmarshaling resonse, got: %v", err))
	c.Assert(result.StatusCode, qt.Equals, http.StatusOK)
	c.Assert(string(responseBody), qt.JSONEquals, &expectedResponse)
}

func TestHandler_Capabilities_GetCapabilitiesSuccess_WithInference(t *testing.T) {
	c := qt.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedCapabilities := []resources.Capability{
		{Endpoint: "/swagger.json", Methods: []resources.CapabilityMethods{"GET"}},
		{Endpoint: "/capabilities", Methods: []resources.CapabilityMethods{"GET"}},
		{Endpoint: "/authentication/providers", Methods: []resources.CapabilityMethods{"GET"}},
		{Endpoint: "/authentication", Methods: []resources.CapabilityMethods{"GET", "POST"}},
		{Endpoint: "/authentication/{id}", Methods: []resources.CapabilityMethods{"GET", "PUT", "DELETE"}},
		{Endpoint: "/identities", Methods: []resources.CapabilityMethods{"GET", "POST"}},
		{Endpoint: "/identities/{id}", Methods: []resources.CapabilityMethods{"GET", "PUT", "DELETE"}},
		{Endpoint: "/identities/{id}/groups", Methods: []resources.CapabilityMethods{"GET", "PATCH"}},
		{Endpoint: "/identities/{id}/roles", Methods: []resources.CapabilityMethods{"GET", "PATCH"}},
		{Endpoint: "/identities/{id}/entitlements", Methods: []resources.CapabilityMethods{"GET", "PATCH"}},
		{Endpoint: "/groups", Methods: []resources.CapabilityMethods{"GET", "POST"}},
		{Endpoint: "/groups/{id}", Methods: []resources.CapabilityMethods{"GET", "PUT", "DELETE"}},
		{Endpoint: "/groups/{id}/identities", Methods: []resources.CapabilityMethods{"GET", "PATCH"}},
		{Endpoint: "/groups/{id}/roles", Methods: []resources.CapabilityMethods{"GET", "PATCH"}},
		{Endpoint: "/groups/{id}/entitlements", Methods: []resources.CapabilityMethods{"GET", "PATCH"}},
		{Endpoint: "/roles", Methods: []resources.CapabilityMethods{"GET", "POST"}},
		{Endpoint: "/roles/{id}", Methods: []resources.CapabilityMethods{"GET", "PUT", "DELETE"}},
		{Endpoint: "/roles/{id}/entitlements", Methods: []resources.CapabilityMethods{"GET", "PATCH"}},
		{Endpoint: "/entitlements", Methods: []resources.CapabilityMethods{"GET"}},
		{Endpoint: "/entitlements/raw", Methods: []resources.CapabilityMethods{"GET"}},
		{Endpoint: "/resources", Methods: []resources.CapabilityMethods{"GET"}},
	}

	expectedResponse := resources.GetCapabilitiesResponse{
		Meta: resources.ResponseMeta{
			Size: len(expectedCapabilities),
		},
		Data:   expectedCapabilities,
		Status: http.StatusOK,
	}

	mockWriter := httptest.NewRecorder()
	mockRequest := httptest.NewRequest(http.MethodGet, "/capabilities", nil)

	// Provide non-nil implementation for all interfaces, expect for the
	// `CapabilitiesService`, to enforce inference.
	sut := handler{
		Identities:        interfaces.NewMockIdentitiesService(ctrl),
		Groups:            interfaces.NewMockGroupsService(ctrl),
		IdentityProviders: interfaces.NewMockIdentityProvidersService(ctrl),
		Entitlements:      interfaces.NewMockEntitlementsService(ctrl),
		Roles:             interfaces.NewMockRolesService(ctrl),
		Resources:         interfaces.NewMockResourcesService(ctrl),
	}
	sut.GetCapabilities(mockWriter, mockRequest)

	result := mockWriter.Result()
	defer result.Body.Close()

	responseBody, err := io.ReadAll(result.Body)
	c.Assert(err, qt.IsNil)

	c.Assert(err, qt.IsNil, qt.Commentf("Unexpected err while unmarshaling response, got: %v", err))
	c.Assert(result.StatusCode, qt.Equals, http.StatusOK)
	c.Assert(string(responseBody), qt.JSONEquals, &expectedResponse)
}

func TestHandler_Capabilities_GetCapabilitiesFailure(t *testing.T) {
	c := qt.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCapabilitiesService := interfaces.NewMockCapabilitiesService(ctrl)
	mockErrorResponseMapper := NewMockErrorResponseMapper(ctrl)

	mockErrorResponse := resources.Response{
		Message: "mock-error",
		Status:  http.StatusInternalServerError,
	}

	mockError := errors.New("test-error")

	mockCapabilitiesService.EXPECT().ListCapabilities(gomock.Any()).Return(nil, mockError)
	mockErrorResponseMapper.EXPECT().MapError(gomock.Eq(mockError)).Return(&mockErrorResponse)

	sut := handler{
		Capabilities:            mockCapabilitiesService,
		CapabilitiesErrorMapper: mockErrorResponseMapper,
	}

	mockWriter := httptest.NewRecorder()
	mockRequest := httptest.NewRequest(http.MethodGet, "/capabilities", nil)

	sut.GetCapabilities(mockWriter, mockRequest)

	result := mockWriter.Result()
	defer result.Body.Close()

	c.Assert(result.StatusCode, qt.Equals, http.StatusInternalServerError)

	data, err := io.ReadAll(result.Body)
	c.Assert(err, qt.IsNil)

	c.Assert(err, qt.IsNil, qt.Commentf("Unexpected err while unmarshaling resonse, got: %v", err))
	c.Assert(result.StatusCode, qt.Equals, http.StatusInternalServerError)
	c.Assert(data, qt.JSONEquals, mockErrorResponse)
}
