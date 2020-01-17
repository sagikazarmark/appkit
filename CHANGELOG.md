# Change Log


All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).


## [Unreleased]

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


[Unreleased]: https://github.com/sagikazarmark/appkit/compare/v0.3.0...HEAD
[0.3.0]: https://github.com/sagikazarmark/appkit/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/sagikazarmark/appkit/compare/v0.1.0...v0.2.0
