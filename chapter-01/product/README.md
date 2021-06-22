# Chapter 01
## product
A Simple RPC and a sever-to-client stream

### Extra package
```bash
> go get -u github.com/golang/protobuf/ptypes/wrappers
> go get -u google.golang.org/grpc/codes
> go get -u	google.golang.org/grpc/status
```

### Compile the helloworld.proto
```bash
> protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/product.proto
```

### Run server
```bash
> go run ./chapter-01/product/server/main.go
```

### Run client
```bash
> go run ./chapter-01/product/client/main.go
```

