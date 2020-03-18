# Todo project

Golang Gin CRUD RESTful with Go Modules, Wire, Gorm and MySQL built with Hexagonal Architecture and DDD

## Project structure 

see: https://github.com/katzien/talks/blob/master/how-do-you-structure-your-apps/confoo-2019-03-13/slides.pdf 
see: https://github.com/golang-standards/project-layout


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
    ├── deleting # Deleting context
    ├── listing # Listing context
    └── storage # Persistance storage implementation
        └── gorm # ORM Gorm
```

Generate wire_gen.go for dependency injection
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

TODO

    - Add Template build environment. Check => https://github.com/thockin/go-build-template
    - Add authentication
        
