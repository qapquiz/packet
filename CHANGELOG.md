# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.5.5] - 2019-12-09
### Fixed
- Wrong len when update content-length in BytesWithHeader()

## [2.5.4] - 2019-12-04
### Fixed
- Increase index more than needed in WriteFloat32 and WriteFloat64

## [2.5.3] - 2019-11-27
### Changed
- Use multiplier instead of continuous grow buffer in growBufferCap

## [2.5.2] - 2019-11-27
### Fixed
- Use len() instead of cap() when NewWriter because it more accurate.
