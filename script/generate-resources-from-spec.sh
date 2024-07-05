#!/usr/bin/env sh
# Copyright (C) 2024 Canonical Ltd.
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as
# published by the Free Software Foundation, either version 3 of the
# License, or (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.


# Generate Go server and types from an OpenAPI spec definition file.
# A local file path can be passed as the first parameter to the script.
#
# If no parameter is passed, the OpenAPI spec will be downloaded from
# the Github repository canonical/openfga-admin-openapi-spec.
#
# This script should be run from the Makefile's parent directory.

set -e

OPENAPI_SPEC_FILE="$1"

if [ -z "$OPENAPI_SPEC_FILE" ]; then
  _tmpdir=/tmp/openfga-admin-openapi-spec

  cleanup() {
    rm -rf "$_tmpdir"
  }
  trap cleanup 0

  cleanup

  mkdir -p "$_tmpdir"
  git clone -q --depth=1 git@github.com:canonical/openfga-admin-openapi-spec "$_tmpdir"
  OPENAPI_SPEC_FILE="${_tmpdir}/openapi.yaml"

elif ! [ -r "$OPENAPI_SPEC_FILE" ]; then
  echo "Can't read the '$OPENAPI_SPEC_FILE', exiting";
  exit 1
fi

go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
oapi-codegen -generate types,spec -package resources -o v1/resources/generated_types.go "$OPENAPI_SPEC_FILE"
oapi-codegen -generate chi-server -package resources -o v1/resources/generated_server.go "$OPENAPI_SPEC_FILE"
