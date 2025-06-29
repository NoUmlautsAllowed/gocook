package v2

import (
	"context"
	"encoding/json"
	"errors"
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
	test := []struct {
		name          string
		client        *http.Client
		get           string
		responseCode  int
		expectedError error
		responseTime  time.Duration
	}{
		{
			"get",
			nil,
			"123456",
			http.StatusOK,
			nil,
			0,
		},
		{
			"error",
			nil,
			"123456",
			http.StatusInternalServerError,
			ErrRequestFailed,
			0,
		},
		{
			"timeout",
			&http.Client{
				Timeout: 50 * time.Millisecond,
			},
			"123456",
			http.StatusOK,
			context.DeadlineExceeded,
			100 * time.Millisecond,
		},
		{
			"forbidden",
			nil,
			"123456",
			http.StatusForbidden,
			ErrRequestForbidden,
			0,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			m := utils.NewMockHandler(ctrl)

			s := httptest.NewServer(m)

			a := API{
				baseRecipeURL: s.URL + "/r",
				baseSearchURL: s.URL + "/s",
			}

			if tt.client != nil {
				a.defaultClient = *tt.client
			}

			m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(tt.responseTime)

				if r.URL.Path != "/r/"+tt.get {
					t.Error("expected", "/r/"+tt.get, "got", r.URL.Path)
				}

				w.WriteHeader(tt.responseCode)
				if tt.responseCode != http.StatusOK {
					return
				}

				e := json.NewEncoder(w)
				err := e.Encode(api.Recipe{})
				if err != nil {
					t.Error("expected no error")
				}
			})

			r, err := a.Get(tt.get)
			if !errors.Is(err, tt.expectedError) {
				t.Error("expected error ", tt.expectedError, " got ", err)
			}

			if r == nil && tt.expectedError == nil {
				t.Error("recipe expected")
			}
			if r != nil && tt.expectedError != nil {
				t.Error("no recipe expected")
			}

			s.Close()
		})
	}
}

func TestV2Api_Search(t *testing.T) {
	test := []struct {
		name          string
		client        *http.Client
		query         api.Search
		expectedQuery string
		responseCode  int
		expectedError error
		responseTime  time.Duration
	}{
		{
			"search",
			nil,
			api.Search{Query: "q", Limit: "1"},
			"limit=1&offset=&query=q&tags=",
			http.StatusOK,
			nil,
			0,
		},
		{
			"timeout",
			&http.Client{
				Timeout: 50 * time.Millisecond,
			},
			api.Search{Query: "q"},
			"limit=1&offset=&query=q&tags=",
			http.StatusOK,
			context.DeadlineExceeded,
			100 * time.Millisecond,
		},
		{
			"forbidden",
			nil,
			api.Search{Query: "q"},
			"limit=1&offset=&query=q&tags=",
			http.StatusForbidden,
			ErrRequestForbidden,
			0,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			m := utils.NewMockHandler(ctrl)

			s := httptest.NewServer(m)

			a := API{
				baseRecipeURL: s.URL + "/r",
				baseSearchURL: s.URL + "/s",
			}
			if tt.client != nil {
				a.defaultClient = *tt.client
			}

			m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(tt.responseTime)

				w.WriteHeader(tt.responseCode)

				if r.URL.Path != "/s/recipes" {
					t.Error("expected recipe search")
				}
				if r.URL.RawQuery != tt.expectedQuery {
					t.Error("expected query ", tt.expectedQuery, " got ", r.URL.RawQuery)
				}

				if tt.responseCode != http.StatusOK {
					return
				}

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

			r, err := a.Search(api.Search{Query: "q", Limit: "1"})
			if !errors.Is(err, tt.expectedError) {
				t.Error("expected error ", tt.expectedError, " got ", err)
			}

			if r == nil && tt.expectedError == nil {
				t.Error("result expected")
			}
			if r != nil && tt.expectedError != nil {
				t.Error("no result expected")
			}

			s.Close()
		})
	}
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
