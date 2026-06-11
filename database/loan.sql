CREATE TABLE loans (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    principal_amount BIGINT NOT NULL,
    interest_rate NUMERIC(5,2) NOT NULL,
    total_amount BIGINT NOT NULL,
    tenure_weeks INT NOT NULL,
    weekly_amount BIGINT NOT NULL,
    start_date DATE NOT NULL,
    status VARCHAR(20) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR NOT NULL,

    is_deleted BOOLEAN
);

CREATE TABLE loan_schedules
(
    id          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    loan_id     BIGINT      NOT NULL REFERENCES loans (id),
    week_number INT         NOT NULL,
    due_date    DATE        NOT NULL,
    amount      BIGINT      NOT NULL,
    status      VARCHAR(20) NOT NULL,
    paid_at     TIMESTAMP NULL,

    created_at  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by  VARCHAR     NOT NULL,
    updated_at  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by  VARCHAR     NOT NULL,

    is_deleted  BOOLEAN,

    UNIQUE (loan_id, week_number)
);

CREATE TABLE payments (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    loan_id BIGINT NOT NULL REFERENCES loans(id),
    amount BIGINT NOT NULL,
    payment_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR DEFAULT 'SYSTEM'
);

CREATE TABLE payment_details (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    payment_id BIGINT NOT NULL REFERENCES payments(id),
    schedule_id BIGINT NOT NULL REFERENCES loan_schedules(id),
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR DEFAULT 'SYSTEM'
);

CREATE INDEX idx_loans_customer_id
    ON "assignment".loans(customer_id);

CREATE INDEX idx_loan_schedules_loan_id
    ON "assignment".loan_schedules(loan_id);

CREATE INDEX idx_loan_schedules_due_date
    ON "assignment".loan_schedules(due_date);

CREATE INDEX idx_payment_allocations_payment_id
    ON assignment.payment_allocations(payment_id);

CREATE INDEX idx_payment_allocations_schedule_id
    ON assignment.payment_allocations(schedule_id);