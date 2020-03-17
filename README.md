Generate wire_gen.go for dependency injection
```
$ cd PROJECT_PATH/todo/internal
$ wire
```

Build the project
```
$ cd PROJECT_PATH/cmd/server
$ go build main.go
```

Run the built (mysql instance dependency)
```
$ PORT=8000 DB_URL="root:@tcp(127.0.0.1:3360)/platform?charset=utf8&parseTime=True&loc=Local" ./cmd/server/server
```