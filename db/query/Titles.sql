-- name: CreateTitles :one
INSERT INTO "Titles" (
    title
) VALUES (
    'MR.'
)
RETURNING *;