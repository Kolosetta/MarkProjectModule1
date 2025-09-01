#!/bin/sh
set -e

echo "Waiting for database at $DB_HOST:$DB_PORT..."

# Ждём готовность БД через psql
while ! PGPASSWORD=$POSTGRES_PASSWORD psql -h "$DB_HOST" -U "$POSTGRES_USER" -d postgres -c '\q' 2>/dev/null; do
  echo "Database not ready yet, retrying..."
  sleep 2
done
  #Прогоняем миграции
echo "Database is ready. Running migrations..."
goose -dir ./db/migrations postgres "$DATABASE_URL" up

echo "Starting application..."
exec ./my-go-app
