# Changelog

## [0.1.0] - 2024-07-23

- Initial release: go_openshowvar library is published.
- A Go library that uses the OpenShowVar protocol to connect to Kuka robots and perform data read/write operations.

## [0.1.1] - 2024-07-23

### Added

- Integrated GitHub Actions CI workflow to automatically run unit tests on every push and pull request to the `main` branch.
- Enhanced continuous integration by ensuring code quality through automated testing.

## [1.0.0] - 2024-12-04

### Added

- Introduced `IsConnected` method to the `OpenShowVar` struct, which allows checking if the connection to the server is active.

## [1.0.1] - 2024-12-05

### Update

- Changed the method name `IsConnect` to `IsConnected` in the OpenShowVar library for consistency and clarity.

## [1.0.2] - 2025-05-22

### Update

- Added connection timeout in `Connect()` method to limit the duration of establishing TCP connection.
- Added read/write timeouts in `Send()` method to prevent indefinite blocking during TCP operations in the OpenShowVar library.
