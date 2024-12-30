INSERT INTO wallets (
    wallet_id, 
    user_id,
    created_at, 
    updated_at
) VALUES ($1, $2, now(), now()) 
RETURNING wallet_id, user_id, created_at, updated_at;

-- name: GetWalletByUserId :one
SELECT wallet_id,user_id, created_at, updated_at FROM wallets WHERE user_id = $1;