-- name: CreateUserRoles :one
INSERT INTO "user_roles" (
    role, role_description
) VALUES (
    $1, $2
)

RETURNING *;
