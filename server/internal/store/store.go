package store

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Store struct {
	сonfig            *Config
	db                *sql.DB
	exampleRepository *ExampleRepository
}

func New() *Store {
	return &Store{
		сonfig: NewConfig(),
	}
}

func (s *Store) Open() error {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s.сonfig.Host,
		s.сonfig.Port,
		s.сonfig.User,
		s.сonfig.Password,
		s.сonfig.DbName)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Can not open postgres.")
	}

	s.db = db

	return db.Ping()
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Example() *ExampleRepository {

	if s.exampleRepository != nil {
		return s.exampleRepository
	}

	s.exampleRepository = &ExampleRepository{
		store: s,
	}

	return s.exampleRepository
}
