# smkent.net web app

[![License](https://img.shields.io/github/license/smkent/net)](https://github.com/smkent/net/blob/main/COPYING)
[![CI](https://github.com/smkent/net/actions/workflows/ci.yaml/badge.svg)](https://github.com/smkent/net/actions/workflows/ci.yaml)
[![Renovate](https://img.shields.io/badge/renovate-enabled-brightgreen?logo=renovatebot)](https://renovatebot.com)
[![GitHub stars](https://img.shields.io/github/stars/smkent/net?style=social)](https://github.com/smkent/net)

## Usage notes

### Static asset overrides

Individual static assets may be overridden at runtime. To do so, set the
`STATIC_OVERRIDE_PATH` to a directory containing replacement files. If a file is
not found in the override path, the server will fall back to the corresponding
embedded asset.

## Development

* Serve locally on port 8080: `go run ./...`
* Run tests: `go test ./...`
* Build and run in Docker:
  `docker build -t net . && docker run --rm -p 8080:8080 net`

## License

Copyright (C) 2026 Stephen Kent and contributors

Licensed under the GNU General Public License v3.0 only
([`GPL-3.0-only`](https://github.com/smkent/net/blob/main/COPYING)).

## Project template

This project is based on and maintained with [copier-project][copier-project].

[copier-project]: https://smkent.github.io/copier-project
