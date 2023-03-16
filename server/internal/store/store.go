package store

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Store struct {
	Config *Config
}

func New() *Store {
	return &Store{
		Config: NewConfig(),
	}
}

func (s *Store) Open() error {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s.Config.Host,
		s.Config.Port,
		s.Config.User,
		s.Config.Password,
		s.Config.DbName)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Can not open postgres.")
	}

	return db.Ping()
}

func (s *Store) Close() {

}
