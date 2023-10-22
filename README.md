# go-grpc-sample

## Table of Contents

+ [Summary](#summary)
+ [References](#references)
+ [History](#history)
+ [How to use](#how-to-use)

## Summary

A repository demonstrating the use of the RPC protocol trough go gRPC

## References

- [Configure Visual Studio Code for Go development](https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code)
- [gRPC Quick Start](https://grpc.io/docs/languages/go/quickstart/)
- [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/)
- [helloworld sample](https://github.com/grpc/grpc-go/tree/master/examples/helloworld)

## History

Following steps shall not be executed. This section tracks all the preconditional steps taken in order to utilize go gRPC. In order to test the samples skip to the [How to use](#how-to-use) section. The steps from **1. Compiling the proto file** will be described for [samples/unary-prc/](./samples/unary-rpc/) and is applied similairly to any of the other existing samples.

**0. Setup**

Go trough the provided links in the [References section](#references).

**1. Compiling the proto file**

A  [hello.proto file](./samples/unary-rpc/proto/hello.proto) is manually created in the form of:

```sh
syntax = "proto3";

package sample;

option go_package = ".";

service Hello {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  string name = 1;
}

message HelloResponse {
  string greeting = 1;
}
```

Execute: 

```sh
cd proto
go mod init hello
protoc.exe --go_out=. --go-grpc_out=. hello.proto
go mod tidy
cd ..
```

This will regenerate the [hello_grpc.pb.go](./samples/unary-rpc/proto/hello_grpc.pb.go) and [hello.pb.go](./samples/unary-rpc/proto/hello.pb.go) files, which contain:
- Code for populating, serializing, and retrieving HelloRequest and HelloReply message types.
- Generated client and server code.

**2. Setting up client and server code utilizing outputs of proto file compilation**

For the client code following steps where executed:

```sh
mkdir -vp client
cd client
go mod init main
# create and implement the main.go (consider go package in proto folder as local dependency)
# update go.mod file
go mod tidy
```

For the server code following steps where executed:

```sh
mkdir -vp server
cd server
go mod init main
# create and implement the main.go (consider go package in proto folder as local dependency)
# update go.mod file
go mod tidy
```

## How to use

Following steps can be executed *repeatedly*.

Ramp up the gRPC server in one of the samples implementations in 1 terminal process, e.g. :

```sh
cd samples/unary-prc/server
go run main.go
```

Run the gRPC client in one of the samples implementations in another terminal process, e.g.:

```sh
cd samples/unary-prc/client
go run main.go SampleGrim
```

The result should look for [samples/unary-prc/](./samples/unary-rpc/) similair to:

![Result](./images/result.PNG)





