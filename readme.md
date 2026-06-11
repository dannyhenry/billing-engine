# Billing Engine

A simple Loan Billing Engine built with:

 - Go 
 - Echo Framework 
 - GORM 
 - PostgreSQL 
 - Clean Architecture 
 - EnvConfig (Kelsey Hightower)

---

# Architecture

The project follows Clean Architecture principles to separate business logic from infrastructure and delivery concerns.

```text
HTTP Request
    │
    ▼
Echo Handler
    │
    ▼
Service Layer
    │
    ▼
Repository Interface
    │
    ▼
GORM Repository
    │
    ▼
PostgreSQL
```