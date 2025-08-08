# Task Manager API

![Go](https://img.shields.io/badge/Go-1.24.5-blue)
![Docker](https://img.shields.io/badge/docker-available-blue)
![Build](https://github.com/madsilver/task-manager/actions/workflows/go.yml/badge.svg)

This project is a Restful API for a maintenance task management system.

---

## How to run the application

### Running with Docker Compose

```bash
docker-compose up -d --build
```

This will start both the MySQL database and the Go API container.

Access the API at: `http://localhost:8000/api/v1`

---

### Running locally via Go

```bash
make run
```
This will start the Go API server on port 8000.

---

## How to run tests

```sh
make test
```

---

## Database setup

⚠️ Important: Run the database migrations before starting the app. This is required to initialize the necessary
tables and seed essential data.
```sh
make migrate
```

---

## Usage Example

- [Example usage of the API](docs/http/examples.http)

---

## Documentation

- [ADRs](docs/adr/)
- [HTTP Examples](docs/http/)
- [Swagger documentation](docs/swagger/swagger.json)
- [Swagger UI](http://localhost:8000/swagger/index.html)
- [Database Migrations](migrations)

---

## Environment Variables

This project uses [`godotenv`](https://github.com/joho/godotenv) to load environment variables from a `.env` file.

Copy the example file and adjust values as needed:

```bash
cp .env.example .env
```

These variables are loaded both locally and by Docker Compose using `env_file`.

⚠️ Important: You must run the following command to generate a secure JWT secret before starting the application:
```sh
make jwt-sec
```

This ensures the `JWT_SECRET` in .env is a strong and valid secret used for signing JWTs.

Additionally, the field `summary` in tasks is encrypted using **AES-256-GCM**. To support this, you must define a secure key:

```sh
make aes-sec
```

This updates the `.env` file with a valid `TASK_ENCRYPTION_KEY` (Base64-encoded 32-byte key).

Ensure the following is present in your `.env` file.

Special attention should be given to the environment variables `DB_HOST` and `RABBITMQ_HOST`. When running the 
application locally, these should be set to `localhost` or the appropriate host.

---

## Endpoints

| Method  | Route              | Description                            |
|---------|--------------------|----------------------------------------|
| POST    | `/api/v1/login`    | Authenticate user and return JWT       |
| POST    | `/api/v1/users`    | Create a new user (manager only)       |
| POST    | `/api/v1/tasks`    | Create a new task (technician only)    |
| GET     | `/api/v1/tasks`    | Get all tasks (manager only)           |
| GET     | `/api/v1/tasks/:id`| Get task by ID (manager or technician) |
| PUT     | `/api/v1/tasks/:id`| Update task (manager only)             |
| DELETE  | `/api/v1/tasks/:id`| Delete task (manager only)             |



Swagger UI is available at: http://localhost:8000/swagger/index.html

---

## Project Structure
This folder structure is based on [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) 
principles and follows [Go's official module layout guidelines](https://go.dev/doc/modules/layout) for scalability and 
maintainability.

```
internal/
├── adapter/
│   ├── handler/        # HTTP handlers (input)
│   ├── presenter/      # ViewModels / output
│   └── repository/     # Database adapters
├── domain/             # Entities + interfaces
├── infrastructure/
│   ├── db/             # Database connection and migrations
│   ├── env/            # Environment variable management
│   ├── logger/         # Logging setup
│   └── server/         # Server and routing setup
├── usecase/            # Business rules
cmd/
└── api/                # Application entry point (main.go)
```

---

## Requirements

* [Go 1.21+](https://golang.org/dl/)
* [Docker](https://www.docker.com/products/docker-desktop)

---

## Makefile Commands

Use the provided `Makefile` to simplify common tasks:
```
Usage:
        make [COMMAND]
Commands:
help            Display help screen
secrets         Generate all secrets and update .env
jwt-sec         Generate a random JWT_SECRET and update .env
aes-sec         Generate a random AES-256 key and update .env
run             Run the api
build           Build program
image           Build docker image
migrate         Run database migrations using Docker
openapi         Build open-api docs
mock            Build mocks
test            Run the tests of the project
html-c          Generate HTML coverage report
lint            Run linter
sec             Run vulnerability scan
```

---

## Notes

* The `domain` layer is free from external dependencies
* The `ContextHTTP` interface decouples HTTP handlers from any specific web framework
* GORM usage is isolated in the repository adapter layer

---

## License

MIT © 2025
