CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE client (
    id uuid DEFAULT uuid_generate_v4 (),
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);