FROM golang:1-alpine AS builder

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/
RUN go build -o smkent-net ./cmd/smkent-net/

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /build/smkent-net ./
COPY static/ static/
COPY templates/ templates/

EXPOSE 8080
CMD ["./smkent-net"]
