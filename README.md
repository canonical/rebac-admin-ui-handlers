# ReBAC Admin UI Handlers

This package is a thin library that implements the *ReBAC Admin UI* OpenAPI [spec][openapi-spec]. With this library comes a set of abstractions that product services, like JAAS or Identity Platform Admin UI, need to implement and plug in. The library itself does not directly communicate with the underlying authorization provider (i.e., OpenFGA).

[openapi-spec]: https://github.com/canonical/openfga-admin-openapi-spec

## Development

To setup development environment follow these steps:

```sh
git clone github.com/canonical/rebac-admin-ui-handlers
# or for SSH-authenticated users:
#   git clone git@github.com:canonical/rebac-admin-ui-handlers
cd rebac-admin-ui-handlers
make mocks
```

To run the unit tests or get a test coverage you can use the following Makefile targets:

```sh
make test
make test-coverage
```

To pull the latest stable OpenAPI spec and generate the request/response types accordingly, run:

```sh
make pull-spec
```
