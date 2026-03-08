# smkent.net web app

## Development

Build and run in Docker:

```sh
docker build -t net . && docker run --rm -p 8080:8080 net
```

Build and run tests locally:

```
go test ./...
```
