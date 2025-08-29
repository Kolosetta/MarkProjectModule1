-- +goose Up
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT NOT NULL,
                       email TEXT NOT NULL,
                       created_at TIMESTAMPTZ DEFAULT now()
);

-- +goose Down
DROP TABLE users;
