# ADR 002: Use of GORM as the ORM Layer

## Status

Accepted

## Context

The project requires interaction with a PostgreSQL database using a relational data model. While building the data access layer, the following options were considered:

* Using the native `database/sql` package with manual queries
* Employing lightweight SQL builders like `sqlx` or `squirrel`
* Adopting a full-featured ORM: GORM

## Decision

GORM was chosen as the Object Relational Mapper for the following reasons:

* **Developer productivity:** GORM provides a rich API for CRUD operations, associations, and transactions, reducing boilerplate code.
* **Mature ecosystem:** It is one of the most widely used ORMs in the Go ecosystem, with good documentation and active maintenance.
* **Extensibility:** GORM supports hooks, migrations, custom data types, and query chaining, which are useful for real-world applications.
* **Integration:** It integrates seamlessly with PostgreSQL and other drivers, and is compatible with the project’s modular structure.
* **Testability:** Despite being a full-featured ORM, the usage is isolated within repository implementations, ensuring clean architectural boundaries.

## Consequences

### Positive:

* Faster development speed due to reduced query boilerplate
* Better maintainability of data access logic
* Readable and expressive query syntax

### Negative:

* Some performance overhead compared to raw SQL
* Additional abstraction may hide subtle query behavior if not used carefully
* Learning curve for advanced features and GORM-specific behavior

## Alternatives Considered

* `database/sql` for maximum control and performance
* `sqlx` for SQL + struct mapping with less abstraction
* `ent` for strict schema modeling

---

**Date:** 2025-08-05  
**Author:** Rodrigo Prata  
**Status:** Recorded  
