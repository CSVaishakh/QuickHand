CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TYPE user_role AS ENUM (
    'client',
    'handyman'
);

CREATE TABLE users (
    user_id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name       VARCHAR(25) NOT NULL,
    last_name        VARCHAR(25) NOT NULL,
    password_hash    TEXT NOT NULL,
    email            VARCHAR(255) UNIQUE NOT NULL,
    role             user_role NOT NULL,
    phone_number     TEXT NOT NULL,
    img              TEXT,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);