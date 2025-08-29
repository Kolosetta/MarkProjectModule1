-- +goose Up

CREATE TABLE post_likes (
                            post_id BIGINT NOT NULL REFERENCES posts(id),
                            user_id BIGINT NOT NULL,
                            PRIMARY KEY (post_id, user_id)
);

-- +goose Down
DROP TABLE post_likes;
