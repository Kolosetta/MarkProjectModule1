package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	godotenv "github.com/joho/godotenv"
	"os"
)

func NewPostgresPool() (*pgxpool.Pool, error) {
	err := godotenv.Load() //Загружает .env файл в программу
	//TODO: переписать получение кредов на нормальное
	dsn := os.Getenv("DATABASE_URL")
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	return pool, nil
}
