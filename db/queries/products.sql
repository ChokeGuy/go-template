-- name: CreateProduct :one
INSERT INTO products (
    product_id, 
    name, 
    description, 
    price, 
    created_at, 
    updated_at
) VALUES ($1, $2, $3, $4, now(), now()) 
RETURNING product_id, name, description, price, created_at, updated_at;

-- name: GetProductById :one
SELECT * FROM products WHERE product_id = $1;

-- name: GetProductByName :one
SELECT * FROM products WHERE name = $1;