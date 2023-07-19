-- name: CreateUser :one
INSERT INTO "users" (
    email, "password", phone, region, gender, title, first_name, last_name, age
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: GetUsers :one
SELECT * FROM "users"
WHERE email = $1 AND password = $2;

-- name: DeleteUser :exec
DELETE FROM "users" WHERE user_id = $1::uuid;
