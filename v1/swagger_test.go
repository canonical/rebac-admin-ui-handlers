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
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestHandler_SwaggerJson(t *testing.T) {
	c := qt.New(t)

	sut := handler{}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/swagger.json", nil)
	sut.SwaggerJson(w, req)

	result := w.Result()
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	c.Assert(err, qt.IsNil)

	parsedSpec := map[string]any{}
	err = json.Unmarshal(body, &parsedSpec)
	c.Assert(err, qt.IsNil)
	c.Assert(len(parsedSpec) > 0, qt.IsTrue)
}
