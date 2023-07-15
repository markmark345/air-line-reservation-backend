-- name: CreateUser :one
INSERT INTO "users" (
    email, "password", phone, region, gender, title, first_name, last_name, age
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- 'user01@gmail.com', 'password', '1234567890', 'thai', 'M', 'Mr.', 'testf', 'testl', 24
