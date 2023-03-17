package store

import "github.com/FFH255/CW-2023/internal/model"

type ExampleRepository struct {
	store *Store
}

func (r *ExampleRepository) Create(e *model.Example) (*model.Example, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (user_name, password) VALUES ($1, $2) RETURNING id",
		e.UserName,
		e.Password,
	).Scan(&e.Id); err != nil {
		return nil, err
	}

	return e, nil
}

func GetAll() {

}
