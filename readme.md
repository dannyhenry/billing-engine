# Billing Engine

Billing Engine is a loan repayment service responsible for:

 - Loan repayment schedule generation 
 - Payment processing 
 - Outstanding balance calculation 
 - Delinquency determination

A simple Loan Billing Engine built with:

 - Go 
 - Echo Framework 
 - GORM 
 - PostgreSQL 
 - Clean Architecture 
 - EnvConfig (Kelsey Hightower)

---

## Architecture

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

## API Endpoints

### Get Outstanding
```
GET /loans/{loan_id}/outstanding
```

### Check Delinquency
```
GET /loans/{loan_id}/delinquent
```

### Pay Loans
```
POST /loans/payments
```

## Authentication

All endpoints require:

```http
Authorization: Bearer <token>
```

Example:

```http
Authorization: Bearer secret-token
```

---

## Assumptions

1. Weekly installment amount is fixed.
2. Payments must be a multiple of weekly installment amount.
3. Payments are allocated to the oldest unpaid schedules first.
4. Partial installment payments are not supported.
5. Overpayments are not supported.
6. Loan is automatically closed when all schedules are paid.
7. Delinquency is determined by the number of due unpaid schedules.

---