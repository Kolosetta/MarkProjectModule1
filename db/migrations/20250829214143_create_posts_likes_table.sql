-- +goose Up
CREATE TABLE post_likes (
                            post_id BIGINT NOT NULL,
                            created_at TIMESTAMPTZ NOT NULL,
                            user_id BIGINT NOT NULL,
                            PRIMARY KEY (post_id, user_id),
                            FOREIGN KEY (post_id, created_at) REFERENCES posts(id, created_at)
                            --TIMESTAMPT добавлен, чтобы корректно ссылаться на таблицу posts
);
-- +goose Down
DROP TABLE post_likes;
