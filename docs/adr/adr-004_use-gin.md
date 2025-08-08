# ADR 004: Use of Gin as the HTTP Web Framework

## Status

Accepted

## Context

The project requires a performant, flexible, and well-supported web framework to handle HTTP requests, middleware, 
routing, and input/output serialization. Several Go frameworks were evaluated for suitability:

* Standard `net/http`
* [Gin](https://github.com/gin-gonic/gin)
* [Chi](https://github.com/go-chi/chi)
* [Fiber](https://github.com/gofiber/fiber)

## Decision

Gin was chosen as the HTTP framework for the project.

### Reasons for Choosing Gin:

* **Developer Experience:** Clean and expressive routing syntax, middleware support, and built-in JSON binding.
* **Performance:** Near-native speed due to its internal use of `httprouter`.
* **Community and Ecosystem:** One of the most popular Go frameworks with a mature ecosystem and abundant documentation.
* **Feature Set:** Built-in middleware support (logging, recovery, CORS), form validation, grouping of routes, query parsing.
* **Flexibility:** While powerful, Gin allows full access to `net/http` primitives when needed.

## Consequences

### Positive:

* Faster development with less boilerplate for routing and binding
* High performance in both dev and production environments
* Easier onboarding for Go developers familiar with web APIs

### Negative:

* Slightly more opinionated than minimal frameworks like Chi
* Requires adapter layer (`ContextHTTP`) to maintain Clean Architecture boundaries

## Alternatives Considered

* **Chi:** Lightweight and idiomatic, but requires more boilerplate for binding and validation
* **net/http:** Minimal dependency, but too verbose for scalable route handling
* **Fiber:** Inspired by Express.js, but not as idiomatic or widely adopted in Go community

---

**Date:** 2025-08-05  
**Author:** Rodrigo Prata  
**Status:** Recorded  
