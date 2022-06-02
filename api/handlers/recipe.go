package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/natnael-meresa/recipe/usecase/recipe"
)

type RecipeHandler struct {
	recipeUseCase *recipe.Service
}

func NewRecipeHandler(recipeUseCase *recipe.Service) *RecipeHandler {
	return &RecipeHandler{
		recipeUseCase: recipeUseCase,
	}
}

func (rh *RecipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	// rh.GetRecipe()
}

func (rh *RecipeHandler) MakeRecipeHandler(r *mux.Router) {
	r.HandleFunc("/recipes", rh.GetRecipe).Methods(http.MethodGet)
	// r.Handle("/v1/recipe", GetRecipe).Methods("GET", "OPTIONS")
}
