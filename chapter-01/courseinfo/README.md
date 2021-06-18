# Chapter 01
## Course Information

### Compile the course_info.proto
```bash
> protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/course_info.proto
```

### Extra go modules
```bash
> go get -u github.com/gofrs/uuid
```

### Server
```bash
> go run server/main.go
```

### Client
```bash
> go run client/main.go
```