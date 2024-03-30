# Task Manager
![Badge](https://img.shields.io/badge/Go-v1.21-blue)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=madsilver_silver-clean-code&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=madsilver_task-manager)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=madsilver_silver-clean-code&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=madsilver_task-manager)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=madsilver_silver-clean-code&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=madsilver_task-manager)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=madsilver_silver-clean-code&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=madsilver_task-manager)

The Task Manager is a software for maintenance tasks performed during a
working day. This application has two types of users (Manager, Technician).
The technician performs tasks and is only able to see, create or update his own
performed tasks.
The manager can see tasks from all the technicians, delete them, and should be
notified when some tech performs a task.
A task has a summary and a date when it was performed, the
summary from the task can contain personal information.

## Configuration
### Environment variables
| variable          | default    | description                             |
|-------------------|------------|-----------------------------------------|
| SERVER_PORT       | 8000       | server port                             |
| MYSQL_USER        | silver     |                                         |
| MYSQL_PASSWORD    | silver     |                                         |
| MYSQL_DATABASE    | silverlabs |                                         |
| MYSQL_HOST        | 127.0.0.1  |                                         |
| MYSQL_PORT        | 3306       |                                         |
| RABBITMQ_USER     | silver     |                                         |
| RABBITMQ_PASSWORD | silver     |                                         |
| RABBITMQ_HOST     | 127.0.0.1  |                                         |
| RABBITMQ_PORT     | 5672       |                                         |

### Roles
| user       | role       |
|------------|------------|
| Manager    | manager    |
| Technician | technician |

## Usage
### Start using it
```shell
make run
```

### Makefile
Use ``make help`` or only ``make`` to check all the available commands.

### HTTP Requests
The x-user-id and x-role headers are required. The API must have a gateway responsible for user authentication.
The gateway is expected to fill the headers with user data.

| header    | required | description    |
|-----------|----------|----------------|
| x-user-id | yes      | user ID        |
| x-role    | yes      | user role      |

#### Example:
```shell
curl --location 'http://localhost:8000/v1/tasks' \
    --header 'x-user-id: 1' \
    --header 'x-role: manager'
```

### Documentation
1. [Docs](docs)
2. [Swagger](docs/swagger.json)
3. [Swagger Panel](http://localhost:8000/swagger/index.html)