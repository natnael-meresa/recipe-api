package recipe

import "github.com/natnael-meresa/recipe/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetRecipe(id entity.ID) (*entity.Recipe, error) {
	r, err := s.repo.Get(id)

	if r == nil {
		return nil, entity.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return r, nil
}
