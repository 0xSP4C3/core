-- Add UUID extension
--CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="GMT";

-- Create users table
CREATE TABLE users (
  -- id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP NULL,
  email VARCHAR (255) NOT NULL UNIQUE,
  password_hash VARCHAR (255) NOT NULL,
  user_status INT NOT NULL,
  user_role VARCHAR (25) NOT NULL
);

-- Create exchange table
CREATE TABLE exchanges (
  -- id            UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
  id            UUID DEFAULT gen_random_uuid() PRIMARY KEY,
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
  id            UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  created_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at    TIMESTAMP NULL,
  name          VARCHAR(25) NOT NULL,
  code          VARCHAR(10) NOT NULL,
  description   VARCHAR(1000) NULL,
  is_deleted    BOOLEAN DEFAULT FALSE,
  exchange_id   UUID NOT NULL,
  CONSTRAINT fk_coin_exchange_id FOREIGN KEY (exchange_id) REFERENCES exchanges (id) ON DELETE CASCADE
);

-- Create coin_uri table
CREATE TABLE coin_uri (
  coin_id       UUID NOT NULL UNIQUE PRIMARY KEY,
  created_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at    TIMESTAMP NULL,
  uri           VARCHAR(2000) NOT NULL,
  CONSTRAINT fk_coin_uri_coin_id FOREIGN KEY (coin_id) REFERENCES coins (id) ON DELETE CASCADE
);

-- Create feed_time table
CREATE TABLE feed_time (
  id                UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  created_at        TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at        TIMESTAMP NULL,
  start_at          TIMESTAMP NOT NULL,
  ended_at          TIMESTAMP NOT NULL
);

-- Create time_frame table
CREATE TABLE time_frame (
  id            UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  created_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at    TIMESTAMP NULL,
  name          VARCHAR(25) NOT NULL,
  description   VARCHAR(1000) NULL,
  is_enabled    BOOLEAN DEFAULT TRUE
);

-- Create feeds table
CREATE TABLE feeds (
  id                UUID DEFAULT gen_random_uuid(),
  created_at        TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at        TIMESTAMP NULL,
  open_bid          DECIMAL NOT NULL,
  close_bid         DECIMAL NOT NULL,
  highest_bid       DECIMAL NOT NULL,
  lowest_bid        DECIMAL NOT NULL,
  total_trade       DECIMAL NOT NULL,
  base_volume       DECIMAL NOT NULL,
  quote_volume      DECIMAL NOT NULL,
  coin_id           UUID NOT NULL,
  feed_time_id      UUID NOT NULL,
  time_frame_id     UUID NOT NULL,
  CONSTRAINT fk_feeds_coin_id FOREIGN KEY (coin_id) REFERENCES coins (id) ON DELETE CASCADE,
  CONSTRAINT fk_feeds_time_id FOREIGN KEY (feed_time_id) REFERENCES feed_time (id) ON DELETE CASCADE,
  CONSTRAINT fk_feeds_range_id FOREIGN KEY (time_frame_id) REFERENCES time_frame (id) ON DELETE CASCADE
);

-- Add indexes
-- users
CREATE INDEX ix_active_users ON users (id) WHERE user_status = 1;
-- exchanges
CREATE INDEX ix_active_exchanges ON exchanges (id) WHERE is_enabled = TRUE;
-- coins
CREATE INDEX ix_active_coins ON coins (id) WHERE is_deleted = FALSE;
-- coin_uri
CREATE INDEX ix_coin_uri_coin_id ON coin_uri (coin_id);
-- feed_time
CREATE INDEX ix_feed_time_start_at ON feed_time (start_at);
-- time_frame
CREATE INDEX ix_active_time_frame ON time_frame (id) WHERE is_enabled = TRUE;
-- feed
CREATE INDEX ix_feed ON feeds (time_frame_id, feed_time_id);
