package v2

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"sync"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/api"
)

func (a *API) Get(id string) (*api.Recipe, error) {
	u, _ := url.Parse(a.baseRecipeURL)
	u.Path = path.Join(u.Path, id)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", a.userAgent)

	resp, err := a.defaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Println(resp.StatusCode, u)
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)

	var recipe api.Recipe
	if err = json.Unmarshal(data, &recipe); err != nil {
		return nil, err
	}

	recipe.PreviewImageURLTemplate = a.replaceImageCdnURL(recipe.PreviewImageURLTemplate)
	return &recipe, nil
}

func (a *API) Search(s api.Search) (*api.RecipeSearch, error) {
	// this is how the api call looks like
	// https://api.chefkoch.de/v2/search/recipe?query=lasagne%20vegan
	// this is how the format should look like crop-480x600
	// https://img.chefkoch-cdn.de/rezepte/2812481433250378/bilder/1185849/<format>/die-ultimative-vegane-lasagne.jpg
	// https://img.chefkoch-cdn.de/rezepte/2812481433250378/bilder/1185849/crop-480x600/die-ultimative-vegane-lasagne.jpg
	// this would be a query for pagination
	// https://api.chefkoch.de/v2/search-frontend/recipes?query=Lasagne+Vegan&limit=41&offset=41&analyticsTags=user,user_logged_out&enableClickAnalytics=true

	// and there is another API one can use to find recipes, that also allows the use of search filters, named as tags
	// a call from the chefkoch.de web UI looks like this:
	// https://api.chefkoch.de/v2/search-gateway/recipes?additionalFields=results.recipe.isSaved,results.recipe.campaign,recipeIntegrationSets&limit=42&offset=0&enableClickAnalytics=true&analyticsTags=web,search-recipe,user,user_logged_out&orderBy=2&query=Lasagne&plusCount=6&tags=57
	// with all unused filters, tracking parameters etc. it can look like this:
	// https://api.chefkoch.de/v2/search-gateway/recipes?limit=42&offset=0&orderBy=2&query=Lasagne&tags=57,50
	// this returns results for an easy, vegan lasagna

	u, _ := url.Parse(a.baseSearchURL)
	u.Path = path.Join(u.Path, "recipes")
	query := make(url.Values)
	query.Set("query", s.Query)
	query.Set("limit", s.Limit)
	query.Set("offset", s.Offset)
	query.Set("tags", s.Tags)
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", a.userAgent)

	resp, err := a.defaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Println(resp.StatusCode, u)
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	var recipeSearch api.RecipeSearch
	err = json.Unmarshal(data, &recipeSearch)
	if err != nil {
		return nil, err
	}

	for i := range recipeSearch.Results {
		r := &recipeSearch.Results[i]
		r.Recipe.PreviewImageURLTemplate = a.replaceImageCdnURL(r.Recipe.PreviewImageURLTemplate)
	}

	return &recipeSearch, nil
}

func (a *API) Comments(c api.CommentQuery) (*api.Comments, error) {
	// An example request URL looks like https://api.chefkoch.de/v2/recipes/876401193058553/comments?offset=20&limit=20&order=1&orderBy=1
	// whereas the number after /recipes/ represents the recipe id
	u, _ := url.Parse(a.baseRecipeURL)
	u.Path = path.Join(path.Join(u.Path, c.RecipeID), "comments")
	query := make(url.Values)
	query.Set("limit", strconv.Itoa(c.Limit))
	query.Set("offset", strconv.Itoa(c.Offset))
	query.Set("order", "0") // sort by date descending (newest first)
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", a.userAgent)

	resp, err := a.defaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Println(resp.StatusCode, u)
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)

	var comments api.Comments
	if err = json.Unmarshal(data, &comments); err != nil {
		return nil, err
	}

	for i := range comments.Results {
		comment := &comments.Results[i]
		comment.Owner.AvatarImageURLTemplate = a.replaceImageCdnURL(comment.Owner.AvatarImageURLTemplate)
	}
	return &comments, nil
}

func (a *API) getInspirationsByType(inspirationsPath string) (*api.RecipeInspirations, error) {
	u, _ := url.Parse(a.baseInspirationsURL)
	u.Path = path.Join(u.Path, inspirationsPath)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", a.userAgent)

	resp, err := a.defaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Println(resp.StatusCode, u)
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)

	var recipeInspirations api.RecipeInspirations
	if err = json.Unmarshal(data, &recipeInspirations); err != nil {
		return nil, err
	}

	for idx := range recipeInspirations.Recipes {
		recipeInspirations.Recipes[idx].PreviewImageURLTemplate = a.replaceImageCdnURL(recipeInspirations.Recipes[idx].PreviewImageURLTemplate)
	}

	return &recipeInspirations, nil
}

func (a *API) Inspirations() (*api.RecipeInspirationsMixed, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	var cookInspirations *api.RecipeInspirations
	var bakeInspirations *api.RecipeInspirations
	var cookInspirationsErr error
	var bakeInspirationsErr error

	go func() {
		cookInspirations, cookInspirationsErr = a.getInspirationsByType("what-to-cook-today")
		wg.Done()
	}()

	go func() {
		bakeInspirations, bakeInspirationsErr = a.getInspirationsByType("what-to-bake-today")
		wg.Done()
	}()

	wg.Wait()

	if cookInspirationsErr != nil {
		return nil, cookInspirationsErr
	}
	if bakeInspirationsErr != nil {
		return nil, bakeInspirationsErr
	}

	return &api.RecipeInspirationsMixed{CookingRecipes: cookInspirations.Recipes, BakingRecipes: bakeInspirations.Recipes}, nil
}
