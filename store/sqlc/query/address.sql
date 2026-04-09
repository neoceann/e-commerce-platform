-- name: CreateAddress :one
INSERT INTO address (
    country,
    city,
    street
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetAddressByID :one
SELECT * FROM address
WHERE id = $1;