# Echo-Template

## Tech stack

- [Golang](https://golang.org/): version 1.21
- [Echo](https://echo.labstack.com/): version 4
- [Gorm](https://gorm.io/): version 1.21
- [JWT](https://jwt.io/): version 3.21
- [PostgreSQL](https://www.postgresql.org/): version 16
- [Cloudinary](https://cloudinary.com/)
- [Docker](https://www.docker.com/): version 20.10
- [Docker-compose](https://docs.docker.com/compose/): version 1.29
- [Makefile](https://www.gnu.org/software/make/manual/make.html): version 4.3
- [Swagger](https://swagger.io/): version 2.0
- [mockery]()
- [assert]()
## Architecture
- Source code base on Domain Driven Design & Clean Architecture
- Base on [Clean Architecture](https://github.com/amitshekhariitbhu/go-backend-clean-architecture) source code
  ![Architecture](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg) by Uncle Bob

Clean Architecture is a software development architecture that has become a standard for building easily maintainable, platform-independent, and source code-reusable applications. Below are some important benefits of using Clean Architecture
-  Use [radix tree](https://en.wikipedia.org/wiki/Radix_tree) to handle save file assets

## Project structure

```

├── makefile # run makefile to run project
├── cmd
│   ├── main.go # entry point
├── presenter
│   ├── controller # To handle requests from clients and return responses
│   ├── middleware # Middleware to handle requests before reaching the controller
│   ├── request # Request struct
│   ├── router # Router
│   ├── docs # Swagger docs
├── usecase # business logic
├── domain
│   ├── assets # save file uploaded
│   ├── entity # database entity
│       ├──  # response.go # struct mapping model to response
│   ├── repo # repository interface
│       ├── model # struct mapping database table
├── pkg
│   ├── api_errors # error response
│   ├── lib # Set up external libraries.
│   ├── constants # constants
│   ├── utils # utils
├── infra # config connecting to external services
├── boostrap # dependencies injection modules
├── config # config file
│   ├── app.env # environment variables
│   ├── config.yaml # config file
│   ├── config.go # load config file
├── migration # Database migrations
```

## Database
![Database design](https://res.cloudinary.com/dsr2xnaj7/image/upload/v1695718620/database_uai7ty.png)

## How to run
- Install [Postgres](https://www.postgresql.org/download/)
- Install [Golang](https://golang.org/doc/install)
- Get KeyAPI of [Cloudinary](https://cloudinary.com/)
- Clone this project
- Create database
- Create .env & config.yml file in config folder
- Run ``` make ``` to run project
- Run ```mockery --dir=domain --output=domain/mocks --outpkg=mocks --all``` to generate mock file


