package repository

import (
	"gorm.io/gorm"
)

type Recipe struct {
	db *gorm.DB
}

func NewRecipe(db *gorm.DB) *Recipe {
	return &Recipe{db}
}

func (r *Recipe) Get(id string) (*Recipe, error) {
	return nil, nil
}

// Create
func (r *Recipe) Create(recipe *Recipe) error {
	return nil
}

// Update
func (r *Recipe) Update(recipe *Recipe) error {
	return nil
}

// Delete
func (r *Recipe) Delete(id string) error {
	return nil
}

// List
func (r *Recipe) List() ([]*Recipe, error) {
	return nil, nil
}

// ListByUser
func (r *Recipe) ListByUser(userID string) ([]*Recipe, error) {
	return nil, nil
}
