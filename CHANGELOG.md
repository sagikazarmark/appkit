# Change Log


All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).


## [Unreleased]


## [0.13.0] - 2021-11-13

### Changed

- Updated dependencies


## [0.12.0] - 2021-09-25

### Changed

- Go kit updated to [v0.12.0](https://github.com/go-kit/kit/releases/tag/v0.12.0)
- Updated dependencies


## [0.11.0] - 2021-07-04

### Changed

- Go kit updated to [v0.11.0](https://github.com/go-kit/kit/releases/tag/v0.11.0)
- Updated dependencies


## [0.10.0] - 2020-05-31

### Changed

- Updated dependencies


## [0.9.0] - 2020-04-09

### Added:

- `run`: `TemporalWorkerActor` function
- `buildinfo` package

### Deprecated

- `run`: `CadenceWorkerRun` function. Use `CadenceWorkerActor` instead.


## [0.8.0] - 2020-02-24

### Added

- `run`: `CadenceWorker` run function


## [0.7.0] - 2020-02-17

### Changed

- Updated dependencies


## [0.6.0] - 2020-02-14

### Added

- `IsServiceError` checker function
- `ServiceErrorMiddleware` endpoint middleware

### Deprecated

- `IsClientError` checker function. Use `IsServiceError` instead.
- `ClientErrorMiddleware` endpoint middleware. Use `ServiceErrorMiddleware` instead.


## [0.5.0] - 2020-01-22

### Added

- HTTP server run actor
- Server log run actor


## [0.4.0] - 2020-01-17

### Added

- `transport/grpc`: `SetStatusMatchers` to override existing matchers
- `transport/http`: `SetProblemMatchers` to override existing matchers

### Changed

- `transport/grpc`: Default matchers are appended to existing ones
- `transport/grpc`: `WithStatusMatchers` appends matchers to existing ones
- `transport/http`: Default matchers are appended to existing ones
- `transport/http`: `WithProblemMatchers` appends matchers to existing ones


## [0.3.0] - 2020-01-13

### Added

- GracefulRestart actor
- gRPC server actor


## [0.2.0] - 2020-01-13

### Added

- Default problem and status converters


## [0.1.0] - 2020-01-12

- Initial release


[Unreleased]: https://github.com/sagikazarmark/appkit/compare/v0.13.0...HEAD
[0.13.0]: https://github.com/sagikazarmark/appkit/compare/v0.12.0...v0.13.0
[0.12.0]: https://github.com/sagikazarmark/appkit/compare/v0.11.0...v0.12.0
[0.11.0]: https://github.com/sagikazarmark/appkit/compare/v0.10.0...v0.11.0
[0.10.0]: https://github.com/sagikazarmark/appkit/compare/v0.9.0...v0.10.0
[0.9.0]: https://github.com/sagikazarmark/appkit/compare/v0.8.0...v0.9.0
[0.8.0]: https://github.com/sagikazarmark/appkit/compare/v0.7.0...v0.8.0
[0.7.0]: https://github.com/sagikazarmark/appkit/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/sagikazarmark/appkit/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/sagikazarmark/appkit/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/sagikazarmark/appkit/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/sagikazarmark/appkit/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/sagikazarmark/appkit/compare/v0.1.0...v0.2.0
