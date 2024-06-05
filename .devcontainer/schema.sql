-- Create a table for Orders
CREATE TABLE orders (
    orderId VARCHAR(255) PRIMARY KEY,
    customerName VARCHAR(255),
    orderDate TIMESTAMP,
    status VARCHAR(50),
    total FLOAT
);

-- Create a table for OrderItems
CREATE TABLE orderItems (
    itemId SERIAL PRIMARY KEY,
    orderId VARCHAR(255),
    productId VARCHAR(255),
    quantity INTEGER,
    price FLOAT,
    FOREIGN KEY (orderId) REFERENCES orders(orderId)
);
