# smkent.net web app

## Development

* Serve locally on port 8080: `go run ./...`
* Run tests: `go test ./...`
* Build and run in Docker:
  `docker build -t net . && docker run --rm -p 8080:8080 net`
