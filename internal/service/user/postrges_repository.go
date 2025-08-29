package user

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

func (rep PostgresRepository) Create(user models.User) error {
	//TODO: добавить наружу метода возврат индекса присвоенного записи
	var userId int64
	query, args, err := rep.psql.
		Insert("users").
		Columns("username", "email").
		Values(user.Username, user.Email).
		ToSql()
	if err != nil {
		return err
	}

	err = rep.db.QueryRow(context.Background(), query, args...).Scan(&userId)
	if err != nil {
		return err
	}
	return nil
}

func (rep PostgresRepository) Get(username string) (models.User, error) {
	query, args, err := rep.psql.
		Select("id", "username", "email").
		From("users").
		Where(squirrel.Eq{"username": username}).
		ToSql()
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = rep.db.QueryRow(context.Background(), query, args...).Scan(&user.Id, &user.Username, &user.Email)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (rep PostgresRepository) GetList() ([]models.User, error) {
	query, args, err := rep.psql.
		Select("id", "username", "email").
		From("users").
		ToSql()
	if err != nil {
		return make([]models.User, 0), err
	}

	rows, err := rep.db.Query(context.Background(), query, args...)
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var currentUser models.User
		if err := rows.Scan(&currentUser.Id, &currentUser.Username, &currentUser.Email); err != nil {
			continue
		}
		users = append(users, currentUser)
	}
	return users, err
}
