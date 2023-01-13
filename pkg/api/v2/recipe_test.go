package v2

import (
	"encoding/json"
	"github.com/NoUmlautsAllowed/gocook/pkg/api"
	"github.com/NoUmlautsAllowed/gocook/pkg/utils"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestV2Api_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)

	//go s.Start()

	a := V2Api{
		baseRecipeUrl: s.URL + "/r",
		baseSearchUrl: s.URL + "/s",
	}

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		e := json.NewEncoder(w)
		e.Encode(api.Recipe{})
		if r.URL.Path != "/r/123456" {
			t.Error("expected 123456")
		}
	})

	r, err := a.Get("123456")
	if err != nil {
		t.Error("did not expect error")
	}

	if r == nil {
		t.Error("recipe expected")
	}

	s.Close()
}

func TestV2Api_Get2(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)

	//go s.Start()

	a := V2Api{
		baseRecipeUrl: s.URL + "/r",
		baseSearchUrl: s.URL + "/s",
	}

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		if r.URL.Path != "/r/123456" {
			t.Error("expected 123456")
		}
	})

	r, err := a.Get("123456")
	if err == nil {
		t.Error("expected error")
	}

	if r != nil {
		t.Error("no recipe expected")
	}

	s.Close()
}

func TestV2Api_Get3(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)

	a := V2Api{
		baseRecipeUrl: s.URL + "/r",
		baseSearchUrl: s.URL + "/s",
		defaultClient: http.Client{
			Timeout: 50 * time.Millisecond,
		},
	}

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
	})

	r, err := a.Get("123456")
	if err == nil {
		t.Error("expected error")
	}

	if r != nil {
		t.Error("no recipe expected")
	}

	s.Close()
}

func TestV2Api_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)

	//go s.Start()

	a := V2Api{
		baseRecipeUrl: s.URL + "/r",
		baseSearchUrl: s.URL + "/s",
	}

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		e := json.NewEncoder(w)
		e.Encode(api.RecipeSearch{
			Results: []api.RecipeSearchResult{
				{},
				{},
				{},
			},
		})
		if r.URL.Path != "/s/recipe" {
			t.Error("expected recipe search")
		}
		if r.URL.RawQuery != "query=q" {
			t.Error("expected query q")
		}
	})

	r, err := a.Search(api.Search{Query: "q"})

	if err != nil {
		t.Error("did not expect error")
	}
	if r == nil {
		t.Error("result expected")
	}

	s.Close()
}

func TestV2Api_Search2(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)

	a := V2Api{
		baseRecipeUrl: s.URL + "/r",
		baseSearchUrl: s.URL + "/s",
		defaultClient: http.Client{
			Timeout: 50 * time.Millisecond,
		},
	}

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
	})

	r, err := a.Search(api.Search{Query: "q"})

	if err == nil {
		t.Error("expected error")
	}
	if r != nil {
		t.Error("no result expected")
	}

	s.Close()
}
