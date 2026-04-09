-- name: CreateCategory :one
INSERT INTO category (
    name
) VALUES (
    $1
)
RETURNING *;

-- name: GetAllCategories :many
SELECT * FROM category;