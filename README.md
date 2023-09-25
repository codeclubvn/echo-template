# Trail backend

## Tech stack

- [Golang](https://golang.org/)
- [Gin](https://echo.labstack.com/)
- [Gorm](https://gorm.io/)
- [JWT](https://jwt.io/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)

## Error handling

Use github.com/pkg/errors to wrap errors and add stack trace to errors

Error in repository layer must be wrapped by errors.Wrap(err,message)

Service layer don't need to wrap errors returned from repository layer, because it already wrapped. Service layer only need to wrap errors returned from it's own logic

## Config

Copy file config.example.yml to config.yml to config for local environment

## Test

Write unit test for function in service layer has complex logic or should be tested
Unit test: run `go test ./test`

## Database
![Database design](deploy/database.png)

## Project structure

```
├── main.go
├── bootstrap
│ └── bootstrap.go // To initialize modules
├── config
│ ├── config.go // Define configs
│ └── config.yml // Define configurations for the local environment
├── route // To define routes
├── library // Setup external libraries
│── controllers // To handle requests from clients and return responses
│── middlewares // Middleware to handle requests before reaching the controller
├── api_errors
│ └── errors.go // Define errors code
├── dto // Define struct requests | response to clients
├── infrastructure // Modules to connect to external services
├── service // To handle project business logic
├── repository // To handle operations with the database
├── models // Define models for mapping with the database
└── utils // Support functions
├── constants // Define constants for the entire project
```
