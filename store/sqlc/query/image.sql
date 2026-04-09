-- name: CreateImage :one
INSERT INTO image (
    image_data,
    product_id
) VALUES (
    $1, $2
)
RETURNING *;

-- name: UpdateImage :one
UPDATE image
SET image_data = $2
WHERE id = $1
RETURNING *;

-- name: DeleteImageByID :exec
DELETE FROM image WHERE id = $1;

-- name: GetImageByProductID :many
SELECT * FROM image
WHERE product_id = $1
ORDER BY created_at ASC;

-- name: GetImageByImageId :one
SELECT * FROM image
WHERE id = $1;