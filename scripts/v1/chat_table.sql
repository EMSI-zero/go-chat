CREATE TABLE conversation(
    id BIGSERIAL PRIMARY KEY,
    is_private BOOLEAN NOT NULL,
    name VARCHAR(128),
    description VARCHAR(300),
    private_members VARCHAR(10) UNIQUE,
    image_path VARCHAR(128),
    created_at TIMESTAMPTZ NOT NULL,
    created_by BIGINT NOT NULL
);


CREATE TABLE user_conversation(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES chat_user(id) ON DELETE CASCADE,
    conversation_id BIGINT NOT NULL REFERENCES conversation(id) ON DELETE CASCADE,
    role_c INT ,
    last_read TIMESTAMPTZ NOT NULL,
    joined_since TIMESTAMPTZ NOT NULL,
    UNIQUE(user_id, conversation_id)
);