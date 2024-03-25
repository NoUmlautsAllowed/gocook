package env

import (
	"testing"
)

func Test_NewEnv_1(t *testing.T) {
	customBindAddress := ":8083"
	customCdnBaseURL := "http://localhost:8081/"
	customAPIBaseURL := "http://localhost:8082/"
	customUserAgent := "foobar"

	t.Setenv(bindAddress, customBindAddress)
	t.Setenv(cdnBaseURL, customCdnBaseURL)
	t.Setenv(apiBaseURL, customAPIBaseURL)
	t.Setenv(userAgent, customUserAgent)

	env := NewEnv()

	if env.BindAddress() != customBindAddress {
		t.Error("env.BindAddress()", customBindAddress)
	}

	if env.CdnBaseURL() != customCdnBaseURL {
		t.Error("env.CdnBaseUrl()", customCdnBaseURL)
	}

	if env.APIBaseURL() != customAPIBaseURL {
		t.Error("env.ApiBaseUrl()", customAPIBaseURL)
	}

	if env.UserAgent() != customUserAgent {
		t.Error("env.UserAgent()", customUserAgent)
	}
}

func Test_NewEnv_2(t *testing.T) {
	customBindAddress := ""
	customCdnBaseURL := ""
	customAPIBaseURL := ""
	customUserAgent := ""

	t.Setenv(bindAddress, customBindAddress)
	t.Setenv(cdnBaseURL, customCdnBaseURL)
	t.Setenv(apiBaseURL, customAPIBaseURL)
	t.Setenv(userAgent, customUserAgent)

	env := NewEnv()

	if env.BindAddress() != defaultBindAddress {
		t.Error("env.BindAddress()", defaultBindAddress)
	}

	if env.CdnBaseURL() != defaultCdnBaseURL {
		t.Error("env.CdnBaseUrl()", defaultCdnBaseURL)
	}

	if env.APIBaseURL() != defaultAPIBaseURL {
		t.Error("env.ApiBaseUrl()", defaultAPIBaseURL)
	}

	if env.UserAgent() != defaultUserAgent {
		t.Error("env.UserAgent()", defaultUserAgent)
	}
}

func TestEnv_String(t *testing.T) {
	customBindAddress := ":8083"
	customCdnBaseURL := "http://localhost:8081/"
	customAPIBaseURL := "http://localhost:8082/"
	customUserAgent := "foobar"

	t.Setenv(bindAddress, customBindAddress)
	t.Setenv(cdnBaseURL, customCdnBaseURL)
	t.Setenv(apiBaseURL, customAPIBaseURL)
	t.Setenv(userAgent, customUserAgent)

	env := NewEnv()

	out := env.String()
	expected := bindAddress + "=" + customBindAddress + "; " + cdnBaseURL + "=" + customCdnBaseURL + "; " + apiBaseURL + "=" + customAPIBaseURL + "; " + userAgent + "=" + customUserAgent

	if out != expected {
		t.Error(out, "!=", expected)
	}
}
