BEGIN;

CREATE TABLE IF NOT EXISTS customer
(
    id         smallserial not null primary key,
    name       text        not null,
    email      text        not null unique,
    created_at timestamptz default now()
);

CREATE TABLE IF NOT EXISTS event
(
    id         smallserial not null primary key,
    "name"     text        not null,
    member     int references customer (id),
    date       timestamptz not null, -- Asia/Yekaterinburg
    created_at timestamptz default now()
);

COMMIT;