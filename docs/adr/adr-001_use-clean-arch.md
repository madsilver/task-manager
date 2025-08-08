# ADR 001: Use of Clean Architecture in the Project

## Status

Accepted

## Context

This Go project was designed to implement a modular, scalable, and testable web API. During the architectural planning 
phase, several approaches were considered:

* Traditional MVC
* Layered monolith (controller, service, repository)
* Hexagonal Architecture (Ports and Adapters)
* Clean Architecture

Clean Architecture was selected as the primary model, with pragmatic adaptations for the Go ecosystem.

## Decision

Adopt Clean Architecture with the following characteristics:

* **Domain layer (core)** containing entities and contracts, free of external dependencies
* **Use case layer** representing the business logic
* **Adapters for input/output** such as HTTP handlers, repositories, and presenters, organized under `internal/adapter`
* **Infrastructure layer** responsible for initializing and configuring external systems (e.g., server, database)
* Use of a `ContextHTTP` interface to decouple handlers from the Gin web framework

This architecture allows:

* Swapping Gin for another web framework without impacting domain logic
* Testing handlers using mocks of the HTTP context interface
* Easier maintenance and clearer separation of concerns

## Consequences

### Positive:

* Highly testable and loosely coupled code
* Scalable structure supporting multiple modules
* Better readability and clearer responsibilities

### Negative:

* Slightly more verbose initial structure
* Higher learning curve for newcomers
* Additional organizational overhead

## Alternatives Considered

* Using plain `Gin` with a flat MVC structure (faster but tightly coupled)
* Hexagonal Architecture (no significant advantage for this use case)

---

**Date:** 2025-08-05  
**Author:** Rodrigo Prata  
**Status:** Recorded  
