CREATE TABLE item(
    chrt_id INT PRIMARY KEY,
    order_uid VARCHAR(50) REFERENCES orders(order_uid),
    track_number VARCHAR(50),
    price INT,
    rid VARCHAR(50),
    name VARCHAR(50),
    sale INT,
    size INT,
    total_price INT, 
    nm_id INT,
    brand VARCHAR(50),
    status INT
)