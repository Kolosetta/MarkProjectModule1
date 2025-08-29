-- +goose Up
CREATE TABLE posts (
                       id SERIAL PRIMARY KEY,
                       author TEXT NOT NULL,
                       text TEXT NOT NULL,
                       likes INT DEFAULT 0,
                       created_at TIMESTAMPTZ DEFAULT now()
);

-- +goose Down
DROP TABLE posts;

