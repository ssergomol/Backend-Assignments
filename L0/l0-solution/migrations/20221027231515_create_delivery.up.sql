CREATE TABLE delivery(
    delivery_id SERIAL PRIMARY KEY,
    order_uid INT UNIQUE FOREIGN KEY REFERENCES order(order_uid),
    name VARCHAR(50) NOT NULL, 
    phone VARCHAR(16),
    zip VARCHAR(16),
    city VARCHAR(50),
    address VARCHAR(100),
    region VARCHAR(50),
    email VARCHAR(300)
);