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
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/go-chi/chi/v5"
)

//go:generate mockgen -package interfaces -destination ./interfaces/mock_authentication.go -source=./interfaces/authentication.go

// TestHandlerWorksWithStandardMux this test ensures that the returned Handler
// can be used with the Golang standard library multiplexers, and it's not tied
// to the underlying router library.
func TestHandlerWorksWithStandardMux(t *testing.T) {
	c := qt.New(t)

	sut, _ := NewReBACAdminBackend(ReBACAdminBackendParams{
		Authenticator: &noopAuthenticator{},
	})
	handler := sut.Handler("/some/base/path/")

	mux := http.NewServeMux()
	mux.Handle("/some/base/path/", handler)

	server := httptest.NewServer(mux)
	defer server.Close()

	println(server.URL)

	res, err := http.Get(server.URL + "/some/base/path/v1/swagger.json")
	c.Assert(err, qt.IsNil)
	c.Assert(res.StatusCode, qt.Equals, http.StatusOK)
	defer res.Body.Close()

	out, err := io.ReadAll(res.Body)
	c.Assert(err, qt.IsNil)
	c.Assert(len(out) > 0, qt.IsTrue)
}

// TestHandlerWorksWithChiMux this test ensures that the returned Handler
// can be used with the Chi multiplexers.
func TestHandlerWorksWithChiMux(t *testing.T) {
	c := qt.New(t)

	sut, _ := NewReBACAdminBackend(ReBACAdminBackendParams{
		Authenticator: &noopAuthenticator{},
	})
	handler := sut.Handler("")

	mux := chi.NewMux()
	mux.Mount("/some/base/path", handler)

	server := httptest.NewServer(mux)
	defer server.Close()

	println(server.URL)

	res, err := http.Get(server.URL + "/some/base/path/v1/swagger.json")
	c.Assert(err, qt.IsNil)
	c.Assert(res.StatusCode, qt.Equals, http.StatusOK)
	defer res.Body.Close()

	out, err := io.ReadAll(res.Body)
	c.Assert(err, qt.IsNil)
	c.Assert(len(out) > 0, qt.IsTrue)
}
