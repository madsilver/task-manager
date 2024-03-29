# Task Manager

![Badge](https://img.shields.io/badge/Go-v1.21-blue)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=madsilver_silver-clean-code&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=madsilver_task-manager)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=madsilver_silver-clean-code&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=madsilver_task-manager)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=madsilver_silver-clean-code&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=madsilver_task-manager)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=madsilver_silver-clean-code&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=madsilver_task-manager)

## Configuration
### Environment variables
| variable          | default    | description                             |
|-------------------|------------|-----------------------------------------|
| SERVER_PORT       | 8000       | server port                             |
| MYSQL_USER        | silver     |                                         |
| MYSQL_PASSWORD    | silver     |                                         |
| MYSQL_DATABASE    | silverlabs |                                         |
| MYSQL_HOST        | localhost  |                                         |
| MYSQL_PORT        | 3306       |                                         |

## Usage
### Start using it
```shell
make run
```

### Makefile
Use ``make help`` or only ``make`` to check all the available commands.