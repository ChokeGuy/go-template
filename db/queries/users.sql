INSERT INTO users (
    user_id, 
    name, 
    created_at, 
    updated_at
) VALUES ($1, $2, now(), now()) 
RETURNING user_id, name, created_at, updated_at;

-- name: GetUserById :one
SELECT user_id FROM users WHERE user_id = $1;