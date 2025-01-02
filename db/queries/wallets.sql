-- name: GetWalletByUserId :one
SELECT wallet_id,user_id, created_at, updated_at FROM wallets WHERE user_id = $1;