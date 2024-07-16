package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Info().Msg("Imported main package")
}

func main() {
	storage, err := newStorage("postgres://admin:donotcrackplease@localhost:5432/university?sslmode=disable")
	if err != nil {
		log.Fatal().Msgf("Creating storage: %v", err)
	}

	users, err := storage.getAllUsers()
	if err != nil {
		log.Fatal().Msgf("Getting all users: %v", err)
	}

	log.Info().Msgf("Got users: %v", users)

	err = storage.insertUser("Taras")
	if err != nil {
		log.Fatal().Msgf("Insertig Taras: %v", err)
	}

	users, err = storage.getAllUsers()
	if err != nil {
		log.Fatal().Msgf("Getting all users: %v", err)
	}

	log.Info().Msgf("Got users: %v", users)
}

type storage struct {
	db *sql.DB
}

func newStorage(connString string) (*storage, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	return &storage{db: db}, nil
}

type user struct {
	id   int
	name string
}

func (s *storage) getAllUsers() ([]user, error) {
	rows, err := s.db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, fmt.Errorf("selecting users: %w", err)
	}

	var users []user

	for rows.Next() {
		var u user

		err := rows.Scan(&u.id, &u.name)
		if err != nil {
			return nil, fmt.Errorf("scanning rows: %w", err)
		}

		users = append(users, u)
	}

	return users, nil
}

func (s *storage) insertUser(name string) error {
	_, err := s.db.Exec("INSERT INTO users(name) VALUES ($1)", name)
	if err != nil {
		return fmt.Errorf("inserting user: %w", err)
	}

	return nil
}
