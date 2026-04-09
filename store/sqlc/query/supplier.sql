-- name: CreateSupplier :one
INSERT INTO supplier (
    name,
    address_id,
    phone_number
)
VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: DeleteSupplier :exec
DELETE FROM supplier
WHERE id = $1;

-- name: GetAllSuppliers :many
SELECT * from supplier
ORDER BY name ASC;

-- name: GetSupplierByID :one
SELECT * from supplier
WHERE id = $1;

-- name: UpdateSupplierAddress :one
UPDATE supplier
SET address_id = $2
WHERE id = $1
RETURNING *;