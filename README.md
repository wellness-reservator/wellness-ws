# wellness-ws

Contains web services for booking lessons

## Running

Running `main.go` starts a web server on https://0.0.0.0:11000/. You can configure
the port used with the `$PORT` environment variable, and to serve on HTTP set
`$SERVE_HTTP=true`.

```
$ go run main.go
```

An OpenAPI UI is served on https://0.0.0.0:11000/.

### Running the standalone server

If you want to use a separate gRPC server, for example one written in Java or C++, you can run the
standalone web server instead:

```
$ go run ./cmd/standalone/ --server-address dns:///0.0.0.0:10000
```

## Requirements

Generating the files requires the `protoc` protobuf compiler.
Please install it according to the
[installation instructions](https://github.com/google/protobuf#protocol-compiler-installation)
for your specific platform.

## Getting started

After cloning the repo, there are a couple of initial steps;

1. Install the generate dependencies with `make install`.
   This will install `protoc-gen-go`, `protoc-gen-grpc-gateway`, `protoc-gen-openapiv2`, `protoc-gen-validate ` and `statik` which
   are necessary for us to generate the Go, swagger and static files.
1. Finally, generate the files with `make generate`.
   If you encounter an error here, make sure you've installed
   `protoc` and it is accessible in your `$PATH`, and make sure
   you've performed step 1.

Now you can run the web server with `go run main.go`.
