SELECT * FROM pg_extension;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE gender AS ENUM ('M', 'F', 'N');

CREATE TABLE "users" (
    user_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(60) NOT NULL,
    phone VARCHAR(13),
    region VARCHAR(100),
    "gender" gender DEFAULT 'N',
    title VARCHAR(50) NOT NULL DEFAULT 'Not Specified',
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    create_at timestamptz NOT NULL DEFAULT (now()),
    update_at timestamptz NOT NULL DEFAULT (now()),
    age SMALLINT NOT NULL
);

CREATE TABLE "roles" (
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL,
    create_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_roles" (
    user_id UUID PRIMARY KEY,
    role INT NOT NULL DEFAULT 0,
    role_description VARCHAR(100),
    create_at timestamptz NOT NULL DEFAULT (now()),
    update_at timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (role) REFERENCES roles(role_id)
);

CREATE TABLE "user_addresses" (
    address_id SERIAL PRIMARY KEY,
    user_id UUID,
    address VARCHAR(255),
    create_at timestamptz NOT NULL DEFAULT (now()),
    update_at timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(user_id)
);