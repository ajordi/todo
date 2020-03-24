# Project TODO

Golang Gin CRUD RESTful with Go Modules, Wire, Gorm and MySQL built with Hexagonal Architecture and DDD

## Tech Stack

1. Go Modules is a dependency management system introduced since Go 1.1+ https://blog.golang.org/using-go-modules
2. Wire is a code generation tool providing compile-time dependency injection for Go https://github.com/google/wire
3. Gin is a popular web framework written in Go https://github.com/gin-gonic/gin
4. Gorm is an ORM library https://gorm.io/

## Project structure 

```
├── cmd # Main applications for this project 
│   └── server # Directory name matching executable
├── internal # Private application and library code
│   ├── api 
│   ├── dao
│   ├── di # Dependency injection
│   └── service
└── pkg # Library code that's ok to use by external applications
    ├── adding # Adding context
    ├── authenticating # Authenticating context
    ├── deleting # Deleting context
    ├── listing # Listing context
    └── storage # Persistance storage implementation
        └── gorm # ORM Gorm
```

Generate `wire_gen.go` for dependency injection

```
$ cd internal/di
$ wire
```

Build the project

```
$ cd cmd/server
$ go build main.go
```

Run the built (mysql instance dependency)
```
$ PORT=8000 DB_URL="root:@tcp(127.0.0.1:3360)/platform?charset=utf8&parseTime=True&loc=Local" ./cmd/server/server
```

## TODO

    - Add Template build environment. Check => https://github.com/thockin/go-build-template

## References

* https://github.com/katzien/talks/blob/master/how-do-you-structure-your-apps/confoo-2019-03-13/slides.pdf 
* https://github.com/golang-standards/project-layout
