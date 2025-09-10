create table if not exists payments (
    correlation_id uuid primary key,
    amount numeric(10,2) not null,
    inserted_at timestamptz(6) not null,
    type varchar(20) not null
);

create index if not exists idx_payments_on_inserted_at_and_type on payments (inserted_at, type);