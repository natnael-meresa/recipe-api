package entity

import "time"

type Recipe struct {
	ID        ID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewRecipe(name string) (*Recipe, error) {
	r := &Recipe{
		ID:   NewID(),
		Name: name,
	}

	err := r.Validate()

	if err != nil {
		return nil, ErrInvalidEntity
	}

	return r, nil
}

func (r *Recipe) Validate() error {
	if r.Name == "" {
		return ErrInvalidEntity
	}

	return nil
}
