package recipe

import "github.com/natnael-meresa/recipe/entity"

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Recipe, error)
	Search(query string) ([]*entity.Recipe, error)
	List() ([]*entity.Recipe, error)
}

//Writer Recipe writer
type Writer interface {
	Create(e *entity.Recipe) (entity.ID, error)
	Update(e *entity.Recipe) error
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetRecipe(id entity.ID) (*entity.Recipe, error)
	SearchRecipes(query string) ([]*entity.Recipe, error)
	ListRecipes() ([]*entity.Recipe, error)
	CreateRecipe(title string, author string, pages int, quantity int) (entity.ID, error)
	UpdateRecipe(e *entity.Recipe) error
	DeleteRecipe(id entity.ID) error
}
