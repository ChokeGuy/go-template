CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP TABLE IF EXISTS products CASCADE;

CREATE TABLE wallets (
    wallet_id UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    balance FLOAT NOT NULL DEFAULT 0,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);