-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="GMT";

-- Create users table
CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    email VARCHAR (255) NOT NULL UNIQUE,
    password_hash VARCHAR (255) NOT NULL,
    user_status INT NOT NULL,
    user_role VARCHAR (25) NOT NULL
);

-- Create books table
CREATE TABLE books (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    title VARCHAR (255) NOT NULL,
    author VARCHAR (255) NOT NULL,
    book_status INT NOT NULL,
    book_attrs JSONB NOT NULL
);

-- Create exchange table
CREATE TABLE exchanges (
  id            UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
  created_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at    TIMESTAMP NULL,
  name          VARCHAR(25) NOT NULL,
  description   VARCHAR(1000) NULL,
  uri           VARCHAR(2000) NULL,
  is_enabled    BOOLEAN DEFAULT TRUE,
  is_blocked    BOOLEAN DEFAULT FALSE,
  is_deleted    BOOLEAN DEFAULT FALSE
);

-- Create coins table
CREATE TABLE coins (
  id            UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
  created_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at    TIMESTAMP NULL,
  exchange_id   UUID NOT NULL REFERENCES exchanges (id) ON DELETE CASCADE,
  name          VARCHAR(25) NOT NULL,
  code          VARCHAR(10) NOT NULL,
  description   VARCHAR(1000) NULL,
  is_deleted    BOOLEAN DEFAULT FALSE
);

CREATE TABLE coin_uri (
  coin_id       UUID NOT NULL REFERENCES coins (id) on DELETE CASCADE,
  uri           VARCHAR(2000) NOT NULL,
  created_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at    TIMESTAMP NULL
);

-- Add indexes
CREATE INDEX active_users ON users (id) WHERE user_status = 1;
CREATE INDEX active_books ON books (title) WHERE book_status = 1;
CREATE INDEX active_exchanges ON exchanges (id) WHERE is_enabled = TRUE;
CREATE INDEX active_coins ON coins (id) WHERE is_deleted = FALSE;
