CREATE TABLE chat_user(
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(128) NOT NULL,
    last_name VARCHAR(128) NOT NULL,
    phone CHAR(11) NOT NULL UNIQUE,
    user_name VARCHAR(128) NOT NULL UNIQUE,
    password_c VARCHAR(128) NOT NULL,
    salt CHAR(8) NOT NULL,
    image_path VARCHAR(256) UNIQUE,
    bio VARCHAR(300),
    created_at TIMESTAMPTZ NOT NULL
);