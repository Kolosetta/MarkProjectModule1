package post

import (
	"MarkProjectModule1/internal/models"
	"context"
	squirrel "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db   *pgxpool.Pool
	psql squirrel.StatementBuilderType
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (rep *PostgresRepository) GetList() []models.Post {
	return []models.Post{}
}

func (rep *PostgresRepository) LikePost(postId int64, userId int64) error {
	return nil
}

func (rep *PostgresRepository) Create(post models.Post) (int64, error) {
	var postId int64
	query, args, err := rep.psql.
		Insert("posts").
		Columns("author", "text").
		Values(post.Author, post.Text).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return -1, err
	}

	err = rep.db.QueryRow(context.Background(), query, args...).Scan(&postId) //scan возвращает id присвоенный записи
	if err != nil {
		return -1, err
	}
	return postId, nil
}
