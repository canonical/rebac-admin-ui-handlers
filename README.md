# ReBAC Admin UI Handlers

This package is a thin library that implements the *ReBAC Admin UI* OpenAPI [spec][openapi-spec]. With this library comes a set of abstractions that product services, like JAAS or Identity Platform Admin UI, need to implement and plug in. The library itself does not directly communicate with the underlying authorization provider (i.e., OpenFGA).

[openapi-spec]: https://github.com/canonical/openfga-admin-openapi-spec

There is an example application of the library, under the `_example` directory, which implements an in-memory server.

## Development

To setup your development environment run these commands:

```sh
git clone github.com/canonical/rebac-admin-ui-handlers
# or for SSH-authenticated users:
#   git clone git@github.com:canonical/rebac-admin-ui-handlers
cd rebac-admin-ui-handlers
make mocks
```

Development-related actions are done via various Makefile targets. You can always run `make help` to find out about the available commands.

Below are some useful Makefile targets you can use.

#### Run tests/coverage

```sh
make test
make test-coverage
```

#### Pull latest OpenAPI spec
To pull the latest stable OpenAPI spec and generate the request/response types accordingly, run:

```sh
make pull-spec
```

#### Clean working directory

```sh
make clean
```
