# Example In-Memory ReBAC Admin

This simple web server implements the `v1` version of the ReBAC Admin OpenAPI [specification][spec]. It's meant to be used by:
- **Back-end** implementors (e.g., JAAS or Identity Platform teams), as an example on how to use the library. The [*Back-end implementation*](#Back-end-implementation) section below provides some useful information in this regard.
- **Front-end** implementors (e.g., the Web team) to test their code against a running server.

[spec]: https://github.com/canonical/openfga-admin-openapi-spec

## Setup and running

Note that you need to have Go (v1.21.3+) installed.

```sh
git clone github.com/canonical/rebac-admin-ui-handlers
# or for SSH-authenticated users:
#   git clone git@github.com:canonical/rebac-admin-ui-handlers

# (change working directory to the location where this `README.md` resides)
cd rebac-admin-ui-handlers/_example
```

you can then run the server by `make run`.

This spins up an HTTP server, listening on `localhost:9999`. You can test if the server is running via:

```sh
curl localhost:9999/rebac/v1/swagger.json
# or
curl localhost:9999/rebac/v1/entitlements/raw
```

## In-memory state

When the server starts, it'll read the `state.json` file and populate the in-memory database from that, and whenever the state changes (e.g., by adding some entity) it'll update `state.json`. Also, before attempting to access in-memory data, the server will reload the `state.json` file, which enables a semi hot-reload behavior.

At the beginning, when there's no `state.json` file, the server loads the initial data from `state.zero.json`. If you needed to reset the in-memory state to the initial state while the server is running, you can send a `GET` request to the `/reset` endpoint:

```sh
curl localhost:9999/reset
```

There's also a shell script, `test.sh` that invokes some API endpoints via `curl`. This is meant to be used as a test and also CLI example reference. You can also use it to populate some data into the running server.

## Testing

You can use `make test` to spin up the server and invoke various API endpoints. Note that when using `make test` the server is reset at the end (by providing `--cleanup` option to the `test.sh` script, which deletes the created entities/relationships) to make sure it's working as expected. If you want to keep the state, just call the underlying shell script `test.sh` with no arguments.


## Back-end implementation

To use the *ReBAC Admin UI Handler* library you should take care of the following:

1. Implement the `Authenticator` interface.
2. Implement `*Service` interfaces that match to your product service's needs.
3. **(Optional)** Implement the `CapabilitiesService` interface.
4. **(Optional)** Implement error response mapping.
5. Create a new instance of the library struct and register HTTP handlers.

These steps are explained as follows.

### 1. Implementing `Authenticator`

Authentication works by implementing the `Authenticator` interface defined in the `v1/interfaces` package. There's only one method, called `Authenticate` that you need to implement:

```go
type Authenticator interface {
    Authenticate(r *http.Request) (any, error) 
}
```

This method gets called on every HTTP request, and is expected to return something (preferably a struct pointer, e.g., `*User`) that will be used by other methods to find out who the calling *user* is. As you can see the HTTP request is provided to this method as a parameter, so it's possible to perform different kinds of authentications (e.g, Basic, token-based, or cookie-based).

If the authentication fails, the method can return an error like this:

```go
return nil, v1.NewAuthenticationError("invalid token")
```

It's best to use the error types defined in the `v1` package, so that the library could respond with the correct HTTP status code. Note that, if there's some other error (e.g., a database/IdP communication failure), the method should not use the `NewAuthenticationError`, because it's basically an internal server error:

```go
return nil, v1.NewUnknownError("database not reachable")
```

Not to mention, whenever authentication fails for a request, the library will bail out with an error response and will not proceed to calling service methods.

Finally, if everything is okay, the method should just return an object that represents the authenticated user:

```go
return &User{
    name: "john doe",
}, nil
```

This returned value, will be accessible to other implemented methods, via the `v1.GetIdentityFromContext` method. More on this in the next subsection.

### 2. Implement `*Service` interfaces

To handle requests, a service has to implement some of the `*Service` interfaces defined in the  `v1/interfaces` package:

- `IdentitiesService`
- `IdentityProvidersService`
- `GroupsService`
- `RolesService`
- `ResourcesService`
- `EntitlementsService`

> ❓ For example implementations of the above interfaces, check out the `cmd/service` package.

> ℹ️ Note that not all of these interfaces have to be implemented. It depends on the needs of your product service.

For example, let's say a product service needs to implement the `/identities/*` endpoints defined in the spec. For this purpose, the `IdentitiesService` interface needs to be implemented. The interface looks like this:

```go
// IdentitiesService defines an abstract backend to handle Identities related operations.
type IdentitiesService interface {
    // ListIdentities returns a page of Identity objects of at least `size` elements if available
    ListIdentities(ctx context.Context, params *resources.GetIdentitiesParams) (*resources.PaginatedResponse[resources.Identity], error)

    // ... More methods
}
```
So, you'll need a `struct` that implements these methods:

```go
type MyIdentitiesService struct {
    // Some fields
}
```

Let's see how the implementation of the `ListIdentities` method should look like:

```go
// ListIdentities returns a page of Identity objects of at least `size` elements if available
func (s *IdentitiesService) ListIdentities(ctx context.Context, params *resources.GetIdentitiesParams) (*resources.PaginatedResponse[resources.Identity], error) {
    raw, _ := v1.GetIdentityFromContext(ctx)
    user, _ := raw.(*User)

    if (!userHasPermission(user)) {
        return nil, v1.NewAuthorizationError("user requires `list-identities` permission")
    }

    // ...
    if (someWentWrong) {
        return nil, v1.NewUnknownError("something went wrong")
    }

    // ..
    return result, nil
}
```

Here, the authenticated `User` struct is taken via the `GetIdentityFromRequest` function. If the user doesn't have the required permissions, we should just return a pre-defined error by using the `NewAuthorizationError` method. Also, if something went wrong (like a database communication failure) you can use the `NewUnknownError` method to return that error.


### 3. Implement `CapabilitiesService` (optional)

The `CapabilitiesService` is meant to provide a self-identification mechanism for the API backend. When the client requests `GET /capabilities`, it'll receive a list of endpoints/methods that the backend implements. It helps with a fine-grained identification of the API capabilities. Also, the front-end application can then enable/disable various commands which improves the user experience.

Implementing the `CapabilitiesService` is optional. If you do not provide an implementation, the library will infer your backend capabilities by checking non-nil implementations of the `*Service` interfaces. However, this inference assumes that all interface methods are effectively implemented (i.e., they are not returning deliberate not-implemented errors). If you are not fully implementing an interface, it's best to provide an implementation of the `CapabilitiesService` that reflects this.

If you need to implement the `CapabilitiesService`, then there's only one method on this interface, called `ListCapabilities` that you need to implement. All that it needs to do is to return an array of endpoint/methods pairs:

```go
type MyCapabilitiesService struct {
}

func (s *MyCapabilitiesService) ListCapabilities(ctx context.Context) ([]resources.Capability, error) {
   return []resources.Capability{
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
    }, nil
}
```

Note that this is the entire list of endpoints. You should only return those that are handled/implemented by your product service.

You also need to register this implementation when you are creating a new instance of the `ReBACAdminBackend` struct:

```go
rebac, err := v1.NewReBACAdminBackend(v1.ReBACAdminBackendParams{
    Capabilities: &MyCapabilitiesService{/*...*/},
    // ...
})
```

### 4. Implement error response mapping (optional)

Optionally, you can have your own mapping to translate `error`s into HTTP responses. Although it's best to use the pre-defined error functions (e.g., `NewUnkownError`), you can also have your own custom error types and return them from the `*Service` interface implementations. In this case, you need to provide an implementation for the `ErrorResponseMapper` interface which is responsible to translate a given `error` into a `Response` struct:

```go
type ErrorResponseMapper interface {
    // MapError maps an error into a Response. If the method is unable to map the
    // error (e.g., the error is unknown), it must return nil.
    MapError(error) *resources.Response
}
```

Note that, you do not need to translate all `error`s passed to the `MapError` method. If it's not something that your product service is aware of, you can just return a `nil`, and the library will do the translation for you, of course, if possible.


### 5. Registering HTTP handlers

For this purpose, you need to create a new instance of the `ReBACAdminBackend` struct via its constructor and provide your implementation of `*Service` interfaces:

```go
rebac, err := v1.NewReBACAdminBackend(v1.ReBACAdminBackendParams{
    Authenticator:     &MyAuthenticator{/*...*/},
    Capabilities:      &MyCapabilitiesService{/*...*/},
    Groups:            &MyGroupsService{/*...*/},
    Identities:        &MyIdentitiesService{/*...*/},
    Roles:             &MyRolesService{/*...*/},
    Entitlements:      &MyEntitlementsService{/*...*/},
    Resources:         &MyResourcesService{/*...*/},
    IdentityProviders: &MyIdentityProvidersService{/*...*/},
})
```

> ❓ For an example on how to create an instance of the library, plug in interface implementations, and register the API endpoints, check out `main.go`.

Then you can use the returned `rebac` struct's `Handler` method to get the HTTP handlers and register them with your HTTP `ServeMux` like this:

```go
mux.Handle("/rebac/", rebac.Handler("/rebac/"))
```

If you use Chi serve mux, you should omit the base URL from the call to the `rebac.Handler`:
```go
mux := chi.NewMux()
mux.Handle("/rebac/", rebac.Handler(""))
```

If it's done correctly, you should be able to access the HTTP endpoints via a `curl` command like this:

```sh
curl <host>:<port>/rebac/v1/swagger.json
```
