CREATE TABLE kline (
    id              UUID UNIQUE DEFAULT gen_random_uuid(),
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    open_at         TIMESTAMP NOT NULL,
    open_bid        DECIMAL NOT NULL,
    highest_bid     DECIMAL NOT NULL,
    lowest_bid      DECIMAL NOT NULL,
    close_bid       DECIMAL NOT NULL,
    close_at        TIMESTAMP NULL,
    timeframe_id    UUID NOT NULL,
    coin_id         UUID NOT NULL,
    CONSTRAINT fk_kline_timeframe_id FOREIGN KEY (timeframe_id) REFERENCES timeframe (id),
    CONSTRAINT fk_kline_coin_id FOREIGN KEY (coin_id) REFERENCES coin (id)
)

CREATE TABLE ema_type (
    id              UUID UNIQUE DEFAULT gen_random_uuid(),
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    name            VARCHAR(30) NOT NULL
)

CREATE TABLE ema (
    id              UUID UNIQUE DEFAULT gen_random_uuid(),
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    value           DECIMAL NOT NULL,
    ema_type_id     UUID NOT NULL,
    kline_id        UUID NOT NULL,
    CONSTRAINT fk_ema_kline_id FOREIGN KEY (kline_id) REFERENCES kline (id),
    CONSTRAINT fk_ema_ema_type_id FOREIGN KEY (ema_type_id) REFERENCES ema_type (id)
);

CREATE TABLE timeframe (
    id              UUID UNIQUE DEFAULT gen_random_uuid(),
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    name            VARCHAR(20) NOT NULL,
)

CREATE TABLE coin (
    id              UUID UNIQUE DEFAULT gen_random_uuid(),
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    name            VARCHAR(500) NOT NULL,
    exchange_id     UUID NOT NULL,
    CONSTRAINT fk_coin_endpoint_exchange_id FOREIGN KEY (exchange_id) REFERENCES exchange (id)
)

CREATE TABLE exchange (
    id              UUID UNIQUE DEFAULT gen_random_uuid(),
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    name            VARCHAR(100),
    enpoint     VARCHAR(500) NOT NULL,
)

