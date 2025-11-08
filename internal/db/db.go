package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// подключение к БД через стандартную библиотеку GO, возвращает err если возникла ошибка
// Getenv - получит значение переменной окружения из контейнера
func Connect() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// подключение
	//sslmode=disable — отключает SSL, чтобы локально не требовался сертификат
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// проверим, что соединение работает
	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the DB!")
	return db, nil
}

func Migrate(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS candidates(
	  id SERIAL PRIMARY KEY,
	  name TEXT NOT NULL,
	  position TEXT NOT NULL,
	  email TEXT NOT NULL,
	  phone TEXT NOT NULL,
	  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("Migration is completed: candidates table is ready")
	return nil
}
