-- name: CreateProduct :one
INSERT INTO product (
    name,
    category_id,
    price,
    available_stock,
    supplier_id
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetProductByID :one
SELECT * FROM product WHERE id = $1;

-- name: GetAvailableProducts :many
SELECT * FROM product
WHERE available_stock > 0
ORDER BY available_stock DESC;

-- name: DeleteProductByID :exec
DELETE FROM product WHERE id = $1;

-- name: DecreaseProductStock :one
UPDATE product
SET available_stock = available_stock - sqlc.arg(decreaseValue)::SMALLINT
WHERE id = $1 
  AND available_stock >= sqlc.arg(decreaseValue)::SMALLINT
RETURNING *;

-- name: IncreaseProductStock :one
UPDATE product
SET available_stock = available_stock + sqlc.arg(increaseValue)::SMALLINT
WHERE id = $1 
RETURNING *;