package api

type RecipeAPI interface {
	Get(id string) (*Recipe, error)
	Search(s Search) (*RecipeSearch, error)
}
