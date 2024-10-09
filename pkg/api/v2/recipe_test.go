package v2

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/api"
	"codeberg.org/NoUmlautsAllowed/gocook/pkg/env"
	"codeberg.org/NoUmlautsAllowed/gocook/pkg/utils"

	"go.uber.org/mock/gomock"
)

func TestV2Api_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)

	// go s.Start()

	a := API{
		baseRecipeURL: s.URL + "/r",
		baseSearchURL: s.URL + "/s",
	}

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		e := json.NewEncoder(w)
		err := e.Encode(api.Recipe{})
		if err != nil {
			t.Error("expected no error")
		}
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

	// go s.Start()

	a := API{
		baseRecipeURL: s.URL + "/r",
		baseSearchURL: s.URL + "/s",
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

	a := API{
		baseRecipeURL: s.URL + "/r",
		baseSearchURL: s.URL + "/s",
		defaultClient: http.Client{
			Timeout: 50 * time.Millisecond,
		},
	}

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(_ http.ResponseWriter, _ *http.Request) {
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

	// go s.Start()

	a := API{
		baseRecipeURL: s.URL + "/r",
		baseSearchURL: s.URL + "/s",
	}

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		e := json.NewEncoder(w)
		err := e.Encode(api.RecipeSearch{
			Results: []api.RecipeSearchResult{
				{},
				{},
				{},
			},
		})
		if err != nil {
			t.Error("expected no error")
		}
		if r.URL.Path != "/s/recipes" {
			t.Error("expected recipe search")
		}
		if r.URL.RawQuery != "limit=1&offset=&query=q&tags=" {
			t.Error("expected query q")
		}
	})

	r, err := a.Search(api.Search{Query: "q", Limit: "1"})
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

	a := API{
		baseRecipeURL: s.URL + "/r",
		baseSearchURL: s.URL + "/s",
		defaultClient: http.Client{
			Timeout: 50 * time.Millisecond,
		},
	}

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(_ http.ResponseWriter, _ *http.Request) {
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

func TestV2Api_UserAgentGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)

	e := env.NewEnv()

	a := API{
		baseRecipeURL: s.URL + "/r",
		baseSearchURL: s.URL + "/s",
		userAgent:     e.UserAgent(),
	}

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		expectedUserAgent := "Mozilla/5.0 (X11; Linux x86_64; rv:108.0) Gecko/20100101 Firefox/108.0"
		if r.UserAgent() != expectedUserAgent {
			t.Error("expected user agent '" + expectedUserAgent + "', got '" + r.UserAgent() + "'")
		}

		w.WriteHeader(200)
		e := json.NewEncoder(w)
		err := e.Encode(api.Recipe{})
		if err != nil {
			t.Error("expected no error")
		}
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

func TestV2Api_UserAgentSearch(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)
	e := env.NewEnv()

	a := API{
		baseRecipeURL: s.URL + "/r",
		baseSearchURL: s.URL + "/s",
		defaultClient: http.Client{
			Timeout: 50 * time.Millisecond,
		},
		userAgent: e.UserAgent(),
	}

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		expectedUserAgent := "Mozilla/5.0 (X11; Linux x86_64; rv:108.0) Gecko/20100101 Firefox/108.0"
		if r.UserAgent() != expectedUserAgent {
			t.Error("expected user agent '" + expectedUserAgent + "', got '" + r.UserAgent() + "'")
		}

		w.WriteHeader(200)
		e := json.NewEncoder(w)
		err := e.Encode(api.RecipeSearch{
			Results: []api.RecipeSearchResult{
				{},
				{},
				{},
			},
		})
		if err != nil {
			t.Error("expected no error")
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
