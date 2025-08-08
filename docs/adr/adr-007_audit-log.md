# ADR 007: Implementing Audit Table for HIPAA Compliance

## Status

Accepted

## Date

2025-08-07

## Context

The application handles sensitive user data, potentially including **Protected Health Information (PHI)**. To comply with the **Health Insurance Portability and Accountability Act (HIPAA)**, specifically section **45 CFR § 164.312(b)**, the system must:

> *"Implement technical procedures to record and examine activity in information systems that contain or use electronic protected health information (ePHI)."*

This mandates **audit controls** to track and review activities involving PHI, especially create, update, and delete operations.

## Decision

We will implement a dedicated audit table named `task_audit_logs` to record actions performed on the `task` entity.

The audit table will include:

- The type of operation: `READ`, `CREATE`, `UPDATE`, or `DELETE`
- The ID of the affected task (`task_id`)
- The ID of the user who performed the operation (`performed_by`), if available
- The timestamp of the operation (`performed_at`)

## Consequences

### Positive

- Ensures compliance with **HIPAA Security Rule §164.312(b)**
- Improves traceability and accountability of critical data changes
- Enables audit reporting and unauthorized access detection
- Provides a foundation to audit other entities in the future

### Negative

- Slight increase in storage due to historical logs
- Added complexity in persistence layer

## Alternatives Considered

- **Database triggers**: Rejected due to reduced testability and flexibility

## Future Work

- Restrict access to `task_audit_logs` table
- Enforce **secure retention** and **monitoring policies** for audit logs

---

**Date:** 2025-08-07  
**Author:** Rodrigo Prata    
**Status:** Recorded  