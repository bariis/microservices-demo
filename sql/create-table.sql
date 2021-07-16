CREATE TABLE IF NOT EXISTS product (
    product_id INT NOT NULL,
    product_name varchar(20) NOT NULL,
    product_description varchar(100) NOT NULL,
    product_price FLOAT NOT NULL,
    PRIMARY KEY (product_id)
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE client (
    id uuid DEFAULT uuid_generate_v4 (),
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);