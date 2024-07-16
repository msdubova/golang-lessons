CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name VARCHAR
);

-- SELECT *
-- FROM users
-- ;

-- INSERT INTO users(name) VALUES ('John');

-- UPDATE users
-- SET name = 'Zhenya'
-- WHERE id = 3
-- ;

CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    product_name VARCHAR NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id)
);

-- INSERT INTO orders(product_name, user_id) VALUES ('iPad', 3);

-- SELECT
--     o.id order_id,
--     o.product_name,
--     u.id user_id,
--     u.name user_name
-- FROM orders o
-- INNER JOIN users u ON o.user_id = u.id
-- ;
