# go-grpc-sample

## Table of Contents

+ [Summary](#summary)
+ [References](#references)
+ [How to use](#how-to-use)

## Summary

A repository demonstrating the use of the RPC protocol trough go gRPC

## References

- [Configure Visual Studio Code for Go development](https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code)
- [gRPC Quick Start](https://grpc.io/docs/languages/go/quickstart/)
- [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/)

## How to use

**0. Setup**

Go trough the provided links in the [References section](#references):
- **Configure Visual Studio Code for Go development**
- **gRPC Quick Start**
- **Protocol Buffer Compiler Installation**

**1. Compiling the proto file**

A  [sample.proto file](./proto/sample.proto) is manually created in the form of:

```go
syntax = "proto3";

package sample;

option go_package = ".";

service HelloService {
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
cd sample
protoc.exe --go_out=. --go-grpc_out=. sample.proto
cd ..
```

This will regenerate the `sample.pb.go` and `sample.pb.go` files, which contain:
- Code for populating, serializing, and retrieving HelloRequest and HelloReply message types.
- Generated client and server code.






