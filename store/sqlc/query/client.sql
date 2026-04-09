-- name: CreateClient :one
INSERT INTO client (
    client_name,
    client_surname,
    birthday,
    gender,
    address_id
)
VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetClientByID :one
SELECT * from client
WHERE id = $1;

-- name: DeleteClient :exec
DELETE FROM client
WHERE id = $1;

-- name: GetClientsByFullName :many
SELECT * FROM client
WHERE client_name ILIKE $1 AND client_surname ILIKE $2;

-- name: GetClientsWithPagination :many
SELECT * FROM client
ORDER BY registration_date
LIMIT $1 OFFSET $2;

-- name: GetAllClients :many
SELECT * FROM client ORDER BY registration_date;

-- name: UpdateClientAddress :one
UPDATE client
SET address_id = $2
WHERE id = $1
RETURNING *;