-- +goose Up
CREATE TABLE posts (
                       id BIGSERIAL,
                       author TEXT NOT NULL,
                       text TEXT NOT NULL,
                       likes INT DEFAULT 0,
                       created_at TIMESTAMPTZ DEFAULT now(),
                       PRIMARY KEY (id, created_at) --Ключ составной, чтобы на него можно было ссылаться снаружи. ПРосто на PK id в партиционированной таблице ссылатсья нельзя
) PARTITION BY RANGE (created_at);

-- Пример партиций по годам
CREATE TABLE posts_2024 PARTITION OF posts
    FOR VALUES FROM ('2024-01-01') TO ('2025-01-01');

CREATE TABLE posts_2025 PARTITION OF posts
    FOR VALUES FROM ('2025-01-01') TO ('2026-01-01');

-- +goose Down
DROP TABLE posts;

