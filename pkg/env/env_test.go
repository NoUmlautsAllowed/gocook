package env

import (
	"os"
	"testing"
)

func Test_NewEnv_1(t *testing.T) {
	customBindAddress := ":8083"
	customCdnBaseUrl := "http://localhost:8081/"
	customApiBaseUrl := "http://localhost:8082/"

	_ = os.Setenv(bindAddress, customBindAddress)
	_ = os.Setenv(cdnBaseUrl, customCdnBaseUrl)
	_ = os.Setenv(apiBaseUrl, customApiBaseUrl)

	env := NewEnv()

	if env.BindAddress() != customBindAddress {
		t.Error("env.BindAddress()", customBindAddress)
	}

	if env.CdnBaseUrl() != customCdnBaseUrl {
		t.Error("env.CdnBaseUrl()", customCdnBaseUrl)
	}

	if env.ApiBaseUrl() != customApiBaseUrl {
		t.Error("env.ApiBaseUrl()", customApiBaseUrl)
	}
}

func Test_NewEnv_2(t *testing.T) {
	customBindAddress := ""
	customCdnBaseUrl := ""
	customApiBaseUrl := ""

	_ = os.Setenv(bindAddress, customBindAddress)
	_ = os.Setenv(cdnBaseUrl, customCdnBaseUrl)
	_ = os.Setenv(apiBaseUrl, customApiBaseUrl)

	env := NewEnv()

	if env.BindAddress() != defaultBindAddress {
		t.Error("env.BindAddress()", defaultBindAddress)
	}

	if env.CdnBaseUrl() != defaultCdnBaseUrl {
		t.Error("env.CdnBaseUrl()", defaultCdnBaseUrl)
	}

	if env.ApiBaseUrl() != defaultApiBaseUrl {
		t.Error("env.ApiBaseUrl()", defaultApiBaseUrl)
	}
}

func TestEnv_String(t *testing.T) {
	customBindAddress := ":8083"
	customCdnBaseUrl := "http://localhost:8081/"
	customApiBaseUrl := "http://localhost:8082/"

	_ = os.Setenv(bindAddress, customBindAddress)
	_ = os.Setenv(cdnBaseUrl, customCdnBaseUrl)
	_ = os.Setenv(apiBaseUrl, customApiBaseUrl)

	env := NewEnv()

	out := env.String()
	expected := bindAddress + "=" + customBindAddress + "; " + cdnBaseUrl + "=" + customCdnBaseUrl + "; " + apiBaseUrl + "=" + customApiBaseUrl

	if out != expected {
		t.Error(out, "!=", expected)
	}
}
