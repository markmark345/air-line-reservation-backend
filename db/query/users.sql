-- name: CreateUser :exec
INSERT INTO "users" (
    email, "password", phone, region, gender, title, first_name, last_name, age
) VALUES (
    'user01@gmail.com', 'password', '1234567890', 'thai', 'M', 'Mr.', 'testf', 'testl', 24
);
