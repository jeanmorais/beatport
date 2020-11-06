# Base build image
FROM golang:1.14-stretch AS base
WORKDIR $GOPATH/src/github.com/jeanmorais/beatport

# Dependencies
FROM base AS dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

#  Test
FROM dependencies AS test
COPY . .
RUN go test -v -cpu 1 -failfast -coverprofile=coverage.out -covermode=set ./...
RUN grep -v "_mock" coverage.out >> /filtered_coverage.out
RUN go tool cover -func /filtered_coverage.out

# Build
FROM dependencies AS build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/beatport ./cmd/beatport

# Package
FROM alpine:3.12.0 AS image
COPY --from=build /go/bin/beatport beatport
ENTRYPOINT ["/beatport"]

