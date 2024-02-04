CREATE TABLE contact(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES chat_user(id) ON DELETE CASCADE,
    contact_id BIGINT NOT NULL REFERENCES chat_user(id) ON DELETE CASCADE,
    contact_name VARCHAR(128),
    UNIQUE(user_id, contact_id)
);