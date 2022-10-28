CREATE TABLE payment (
    transaction VARCHAR(50) PRIMARY KEY,
    order_uid INT UNIQUE FOREIGN KEY REFERENCES orders(order_uid),
    request_id VARCHAR(50),
    currency VARCHAR(8),
    provider VARCHAR(50),
    amount INT,
    payment_dt INT,
    bank VARCHAR(50),
    delivery_cost INT,
    goods_total INT,
    custom_fee INT,
) 