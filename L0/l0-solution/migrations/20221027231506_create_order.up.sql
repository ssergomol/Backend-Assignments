CREATE TABLE order(
    order_uid VARCHAR(50) PRIMARY KEY,
    track_number VARCHAR(50),
    entry VARCHAR(50),
    locale VARCHAR(3),
    internal_signature VARCHAR(50),
    customer_id VARCHAR(50),
    delivery_service VARCHAR(50),
    shardkey VARCHAR(50),
    sm_id INT,
    date_created CHAR(20),
    oof_shard VARCHAR(50)
);