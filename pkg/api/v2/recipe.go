package v2

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/NoUmlautsAllowed/gocook/pkg/api"
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
	// this is how the format should look like: crop-480x600, example
	// https://img.chefkoch-cdn.de/rezepte/2812481433250378/bilder/1185849/<format>/die-ultimative-vegane-lasagne.jpg
	// https://img.chefkoch-cdn.de/rezepte/2812481433250378/bilder/1185849/crop-480x600/die-ultimative-vegane-lasagne.jpg
	// this would be a query for pagination
	// https://api.chefkoch.de/v2/search-frontend/recipes?query=Lasagne+Vegan&limit=41&offset=41&analyticsTags=user,user_logged_out&enableClickAnalytics=true
	u, _ := url.Parse(a.baseSearchURL)
	u.Path = path.Join(u.Path, "recipe")
	query := make(url.Values)
	query.Set("query", s.Query)
	query.Set("limit", s.Limit)
	query.Set("offset", s.Offset)
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
