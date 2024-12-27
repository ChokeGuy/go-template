INSERT INTO wallets (
    wallet_id, 
    user_id,
    balance,
    created_at, 
    updated_at
) VALUES ($1, $2,$3, now(), now()) 
RETURNING wallet_id, user_id, balance, created_at, updated_at;

-- name: GetWalletById :one
SELECT wallet_id,user_id,created_at,updated_at FROM wallets WHERE wallet_id = $1;