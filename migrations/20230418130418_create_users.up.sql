CREATE TABLE users(
    id bigserial primary key,
    email varchar not null unique,
    encrypted_password varchar not null unique
)