# go-clean

Golang REST API service using Fiber framework and GORM with MySQL database by applying clean architecture and dependency injection by Wire.

## Template Structure

* [Fiber](https://github.com/gofiber/fiber) is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for fast development with zero memory allocation and performance in mind.
* [JWT](github.com/golang-jwt/jwt) A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens.
* [GORM](https://gorm.io/index.html) with [PostgresSQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)The fantastic ORM library for Golang aims to be developer friendly.
* [Wire](https://github.com/google/wire) is a code generation tool that automates connecting components using dependency injection.
* [validator](github.com/go-playground/validator) is a package validator implements value validations for structs and individual fields based on tags.
* [Viper](https://github.com/spf13/viper) is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.
* [mockery](https://github.com/vektra/mockery) provides the ability to easily generate mocks for Golang interfaces using the stretchr/testify/mock package. It removes the boilerplate coding required to use mocks.
* [UUID](https://github.com/google/uuid) generates and inspects UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.
* [zap](https://github.com/uber-go/zap) is a Blazing fast, structured, leveled logging in Go.

## Feature

* Basics CRUD REST API service with Fiber framework + extra advanced endpoint to partially search text match
* Using ORM library from GORM
* Apply clean architecture with dependency injection by Wire
* Using JWT for authentication
* Using Viper for env injection
* Support Prefork feature from Fiber to enable `SO_REUSEPORT` the load balancer at OS level to create multiple servers; for more information [fasthttp](https://pkg.go.dev/github.com/valyala/fasthttp/reuseport) , [nginx](https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/), [Fiber](https://github.com/gofiber/fiber/issues/180)
* Implement optimistic lock by versioning with GORM for avoid race conditions and concurrency
* Creating a persistent model at repository level to avoid the inhabit access from outsider
* Creating a response model at handler level to avoid the application model connect to the outsider
* Using Goolgle UUID for primary key in database table
* Using Zap as a logger for application
* Containerize with application and Postgres database with Docker
* Using mockery to generate mocks for usecase and repository interfaces
* Test coverage for all handler and usecase is 100%

## Using `go-fiber-clean-arch` project

### Prerequisite

1. Set up your local PostgresSQL database
2. Set up env value which required env key is

* DB_HOST=
* DB_NAME=
* DB_USER=
* DB_PORT=
* DB_PASSWORD=

3. Additional env key is

* RECOVER - to enable the reocery mode for Fiber framework
* TRACING - to enable the log tracing mode for Fiber framework
* PREFORK - to enable use of the SO_REUSEPORT socket option. This socket option allows multiple sockets to listen on the same IP address and port combination. The kernel then load balances incoming connections across the sockets.

### Run application

To use `go-fiber-clean-arch` project, follow these steps:

```bash
# Install dependencies
make deps

# Generate wire_gen.go for dependency injection
# Please make sure you are export the env for GOPATH
make wire

# Run the project in Development Mode
make run
```

### Run application with database via Docker container

To use `go-clean` project, follow these steps:

```bash
* Update .env file for `DB_HOST` key to match with database service name in docker compse

# Build and run application with Docker compose
make docker-compose-run
```

Additional commands:

```bash
➔ make help
build                          Compile the code, build Executable File
run                            Start application
test                           Run tests
test-coverage                  Run tests and generate coverage file
deps                           Install dependencies
deps-cleancache                Clear cache in Go module
wire                           Generate wire_gen.go
mockery                        Generate mock file
help                           Display this help screen
docker-build                   Build docker image with default setting and platform
docker-run                     Run docker image
docker-compose-run             Run docker image with postgres database in the contianer
```

## Available Endpoint

In the project directory, you can call:

### `GET /api`

* For getting status page

### `GET /api/v1/login`

* For generating a JWT

### `GET /api/v1/users`

* For getting list of users

### `GET /api/v1/users/:id`

* For getting user by ID

### `POST /api/v1/users`

* For creating new user

### `DELETE /api/v1/users/:id`

* For removing existing user

### `PUT /api/v1/users/:id`

* For updating existing user

### `GET /api/v1/users/name/:text`

* For retrieving a list of user information that their
name match or partially match with the specified text.

## Folder Structure

This project design by using clean architecture and hexagonal architecture so folder of project will organize base on
clean architecture below

Ref: <https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html>

Here below is how folder map to layer and component in clean architecture

* domain -> Entity
* usecase -> Usecase
* repository -> Repository
* api -> Handler
* driver -> remote call

## Trade-offs

* I'm prefer ORM library to provides an abstraction layer between the Go code and the database over the raw SQL. GORM makes less boilerplate code and fewer manual SQL queries. GORM provides official support for multiple database engines, and a simple CRUD interface for querying and manipulating database records. Furthermore, if the basic function from GORM is not enough, GORM still support query raw SQL too. On the other hand, GORM usually execute SQL queries under the hood, they can only hope to match the performance of an equivalent optimized SQL query; in practice, though, ORMs are often slower.
