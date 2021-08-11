CREATE TABLE IF NOT EXISTS product (
    product_id INT NOT NULL,
    product_name varchar(20) NOT NULL,
    product_description varchar(100) NOT NULL,
    product_price FLOAT NOT NULL,
    PRIMARY KEY (product_id)
);