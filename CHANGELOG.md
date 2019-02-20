# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased
### Changed
- **BREAKING** It is now configurable whether to emit non-runtime metrics with `host` tag or not. Runtime config now expects the boolean field `metrics.m3.includeHost` to be set. Runtime metrics are always emitted with `host` tag.
- Removed logger metrics since it is barely useful.

## 0.2.0 - 2019-01-17
### Added
- Application configuration (e.g. `config/base.json`) can now be specified in YAML in addition to JSON (#504). Zanzibar will look for the `yaml` file first, and fall back to `json` file if it does not exist. JSON static configuration support may be removed in future major releases.
- Module configuration (`services/<name>/service-config.json`) can now be specified as YAML ina ddition to JSON (#468).
- Panics in endpoints are now caught (#458). HTTP endpoints return `502` status code with a body `"Unexpected workflow panic, recovered at endpoint."`. TChannel endpoints will return `ErrCodeUnexpected`.
- Transport specific client config structs added (`HTTPClientConfig`, `TChannelClientConfig`, `CustomClientConfig`) that match the JSON serialized objects in `client-config.json` for the supported client transports.
- Client calls are now protected with circuit breaker (https://github.com/uber/zanzibar/pull/539). Circuit breaker is enabled by default for each client, it can be disabled or fine tuned with proper configurations. It also emits appropriate metrics for monitoring/alerting.

### Changed
- **BREAKING** All [`metrics`](https://godoc.org/github.com/uber/zanzibar/runtime#call_metrics.go) counter and timer name has been changed and using RootScope instead of AllHostScope and PerHostScope since all parameter at name (e.g. host, env and etc) is already moved to tags.(e.g. fetch name:$service.$env.per-workers.inbound.calls.recvd is changed to fetch name:endpoint.request env:$env service:$service)
- **BREAKING** Application packages must now export a global variable named `AppOptions` of type [`*zanzibar.Options`](https://godoc.org/github.com/uber/zanzibar/runtime#Options) to be located at package root (the package defined in `build.json`/`build.yaml`).
- **BREAKING** `codegen.NewHTTPClientSpec`, `codegen.NewTChannelClientSpec`, `codegen.NewCustomClientSpec` and `codegen.ClientClassConfig` removed ([#515](https://github.com/uber/zanzibar/pull/515)).
- **BREAKING** HTTP router [`runtime.HTTPRouter`](https://godoc.org/github.com/uber/zanzibar/runtime#HTTPRouter) method `Register` renamed to `Handle` to better unify with the `net/http` standard library.
- **BREAKING** HTTP router type `runtime.HTTPRouter` switched from exposed concrete type to an interface type to allow changing the implementation.
- **BREAKING** `ServerHTTPRequest.Params` type changed from `julienschmidt/httprouter.Params` to `url.Values` from the standard library.
- Application logs should use the context logger in DefaultDeps.
- Added [`ContextExtractor`](https://godoc.org/github.com/uber/zanzibar/runtime#ContextExtractor) interface. It is part of the API for defining "extractors" or functions to pull out dynamic fields like trace ID, request headers, etc. out of the context to be used in log fields and metric tags. These can be used to pull out fields that are application-specific without adding code to zanzibar.
- Zanzibar now requires `yq` to be installed, whereas it previously required `jq` to be installed. `yq` is available over [PyPI](https://pypi.org/project/yq/) (`pip install yq`) and homebrew (`brew install python-yq`).
- Integrated with [Fx](http://go.uber.org/fx) in the main loop of the generated application.

### Deprecated
- JSON static configuration support is now deprecated.
- `JSONFileRaw` and `JSONFileName` fields of [`ModuleInstance`](https://godoc.org/github.com/uber/zanzibar/codegen#ModuleInstance) are now deprecated. When YAML configuration is used, `JSONFileRaw` and `JSONFileName` will be zero-valued.
- Exported types like [`ClientClassConfig`](https://godoc.org/github.com/uber/zanzibar/codegen#ModuleInstance) will have their JSON tags removed in the future.
- `gateway.Logger` is deprecated. Applications should get `ContextLogger` from `DefaultDeps` instead. Internal libraries can use the unexported `gateway.logger`.

### Fixed
- HTTP `DELETE` methods on clients can now send a JSON payload. Previously it was silently discarded.
- Fixed typo in metrics scope tag for `protocol` (was `protocal`).

## 0.1.2 - 2018-08-28
### Added
- **Context logger**: Added [`ContextLogger`](https://godoc.org/github.com/uber/zanzibar/runtime#ContextLogger) interface. This new logger interface automatically has log fields for Zanzibar-defined per-request fields like `requestUUID` to allow correlating all log messages from a given request. [`WithLogFields`](https://godoc.org/github.com/uber/zanzibar/runtime#WithLogFields) method added to `runtime` package to allow users to add their own log fields to be used by subsequent logs that use the context. The context logger is added as a new field to [`DefaultDependencies`](https://godoc.org/github.com/uber/zanzibar/runtime#DefaultDependencies).

### Changed
- Removed support for Go 1.9 and added support for Go 1.11.
- Some request fields like `endpointID` will no longer be present on messages using the default logger.

### Deprecated
- The default logger (`DefaultDependencies.Logger`) is now deprecated along with multiple related functions and public variables. The preferred way to log is now using the context logger.

## 0.1.1 - 2018-08-21
### Added
- Upgraded thriftrw to v1.12.0 from v1.8.0. This adds, among other things, `Ptr()` helper methods and getters for some thriftrw-go defined structs.

## 0.1.0 - 2018-08-17
### Added
- Initial release.
