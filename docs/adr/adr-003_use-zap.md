# ADR 003: Use of Uber Zap for Structured Logging

## Status

Accepted

## Context

The project requires consistent and performant logging to aid in debugging, monitoring, and observability. The logs should support both development (human-readable) and production (structured/JSON) formats.

Go offers the standard `log` package, but it lacks features such as log levels, structured fields, and output customization. Alternatives like `logrus`, `zerolog`, and `zap` were considered.

## Decision

[Uber Zap](https://github.com/uber-go/zap) was selected as the logging library for this project.

### Reasons for Choosing Zap:

* **Performance:** Zap is one of the fastest structured loggers in the Go ecosystem, optimized for high-throughput systems.
* **Structured Logging:** It supports strongly typed, structured log fields, which integrate well with observability tools like ELK/EFK stacks.
* **Flexibility:** Provides both `SugaredLogger` (ergonomic) and `Logger` (performance-optimized) APIs.
* **Production Readiness:** Widely used in production-grade services by companies like Uber.
* **Customizability:** Supports easy configuration of log levels, formats (console/JSON), and output destinations.

## Consequences

### Positive:

* High-performance, structured, and scalable logging
* Easy integration with metrics and log pipelines
* Consistent logging across modules and layers

### Negative:

* Slightly steeper learning curve compared to simpler libraries
* Requires wrapper or factory to abstract usage in domain/usecase layers if needed

## Alternatives Considered

* **logrus**: Easy to use but slower performance and no longer actively evolving
* **zerolog**: Very fast and efficient, but with a less conventional API
* **standard log**: Too limited for structured production logging

---

**Date:** 2025-08-05  
**Author:** Rodrigo Prata  
**Status:** Recorded  
