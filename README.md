# smkent.net web app

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
