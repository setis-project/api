CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS admin_account (
    id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4 (),
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR NOT NULL,

    PRIMARY KEY (id)
);