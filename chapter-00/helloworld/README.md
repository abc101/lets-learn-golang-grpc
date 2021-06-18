# Chapter 00
## Hello World

### Compile the helloworld.proto
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/helloworld.proto
```

### Run server
```bash
go run ./chapter-00/server/main.go
```

### Run client
```bash
go run ./chapter-00/client/main.go
```

### Results (server-side)
```bash
Received: World
```

### Results (client-side)
```bash
Say: Hello World
```

### With a name
```bash
go run ./chapter-00/server/main.go Alice
```
