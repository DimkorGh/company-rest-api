# Company RESTful API

This repo includes a RESTful API microservice, written in golang, in order to expose simple CRUD operations for a company entity.

### Makefile Commands

he project provides the following features:

- RESTful endpoints
- CRUD operations
- JWT-based authentication
- Data validation
- Unit tests
- Integration tests
- Swagger api documentation
- Postman collection and environment
- Git Actions for linting and unit tests
- Multi-Stage Dockerfile

The project uses the following Go packages:

- Routing: **[gorilla-mux](https://github.com/gorilla/mux)**
- Logging: **[zap](https://github.com/uber-go/zap)**
- Database: **[mongo](https://github.com/mongodb/mongo-go-driver)**
- Event Producing: **[confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go)**
- Configuration: **[viper](https://github.com/spf13/viper)** 
- Env: **[godotenv](https://github.com/joho/godotenv)**
- Migration: **[golang-migrate](https://github.com/golang-migrate/migrate)**
- Validation: **[go-playground](https://github.com/go-playground/validator)**
- Jwt: **[jwt-go](https://github.com/dgrijalva/jwt-go)**
- Mocking **[go-mock](https://github.com/golang/mock)**
- Assertions **[testify](https://github.com/stretchr/testify)**

### Makefile Commands

| Command                               | Usage                                                      |
|---------------------------------------|------------------------------------------------------------|
| docker.app.start                      | `Start all services`                                       |
| docker.format                         | `Reformats go source code via docker`                      |
| docker.linter.run                     | `Run linter`                                               |
| docker.test.unit                      | `Run unit tests via docker`                                |
| docker.test.all                       | `Run both unit and integration tests via docker`           |
| docker.test.all.coverage.withView     | `Run both unit and integration tests via docker with view` |
| docker.mock.generate FILE={FILE_PATH} | `Generate mock for a specific file via docker`             |
| docker.swagger.generate               | `Generate swagger yaml and json files`                     |


* In order to execute makefile commands type **make** plus a command from the table above

  make {command}

---

### Notes

* *mongo-init.js* is used in order to add unique constraint to company's collection name field
* .env is pushed to the repository for testing purposes ONLY. In a production environment it should never be tracked
