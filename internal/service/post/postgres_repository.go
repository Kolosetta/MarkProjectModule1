package post

import (
	"MarkProjectModule1/internal/models"
	"context"
	"github.com/Masterminds/squirrel"
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

func (rep *PostgresRepository) GetList() ([]models.Post, error) {
	query, args, err := rep.psql.
		Select("id", "author", "text", "likes").
		From("posts").
		OrderBy("id DESC").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := rep.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.Id, &p.Author, &p.Text, &p.Likes); err != nil {
			continue
		}
		posts = append(posts, p)
	}
	return posts, nil

}

func (rep *PostgresRepository) LikePost(postId int64, userId int64) error {
	//TODO: добавить транзакцию, оборачивающую это все
	query, args, err := rep.psql.
		Insert("post_likes").
		Columns("post_id", "user_id").
		Values(postId, userId).
		ToSql()
	if err != nil {
		return err
	}

	_, err = rep.db.Exec(context.Background(), query, args...)
	if err != nil {
		return err
	}
	//Обновление счетчика у самого поста
	_, err = rep.db.Exec(context.Background(), "UPDATE posts SET likes = likes + 1 WHERE id=$1", postId)
	return err
}

func (rep *PostgresRepository) Create(post models.Post) error {
	//TODO: добавить наружу метода возврат индекса присвоенного записи
	var postId int64
	query, args, err := rep.psql.
		Insert("posts").
		Columns("author", "text").
		Values(post.Author, post.Text).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return err
	}

	err = rep.db.QueryRow(context.Background(), query, args...).Scan(&postId) //scan возвращает id присвоенный записи
	if err != nil {
		return err
	}
	return nil
}
