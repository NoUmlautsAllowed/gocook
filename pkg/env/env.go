package env

import (
	"log"
	"net/url"
	"os"
)

type Env struct {
	bindAddress string
	cdnBaseURL  string
	apiBaseURL  string
	userAgent   string
}

const (
	defaultBindAddress string = ":8080"
	defaultCdnBaseURL  string = "https://img.chefkoch-cdn.de/"
	defaultAPIBaseURL  string = "https://api.chefkoch.de/"
	defaultUserAgent   string = "Mozilla/5.0 (X11; Linux x86_64; rv:108.0) Gecko/20100101 Firefox/108.0"
)

const (
	bindAddress string = "BIND_ADDRESS"
	cdnBaseURL  string = "CDN_BASE_URL"
	apiBaseURL  string = "API_BASE_URL"
	userAgent   string = "USER_AGENT"
)

func NewEnv() *Env {
	env := &Env{
		bindAddress: os.Getenv(bindAddress),
		cdnBaseURL:  os.Getenv(cdnBaseURL),
		apiBaseURL:  os.Getenv(apiBaseURL),
		userAgent:   os.Getenv(userAgent),
	}

	if len(env.bindAddress) == 0 {
		log.Println("Using", defaultBindAddress, "as", bindAddress)
		env.bindAddress = defaultBindAddress
	}

	if _, err := url.Parse(env.cdnBaseURL); err != nil || len(env.cdnBaseURL) == 0 {
		log.Println("Using", defaultCdnBaseURL, "as", cdnBaseURL)
		env.cdnBaseURL = defaultCdnBaseURL
	}

	if _, err := url.Parse(env.apiBaseURL); err != nil || len(env.apiBaseURL) == 0 {
		log.Println("Using", defaultAPIBaseURL, "as", apiBaseURL)
		env.apiBaseURL = defaultAPIBaseURL
	}

	if len(env.userAgent) == 0 {
		env.userAgent = defaultUserAgent
	}

	return env
}

func (e Env) String() string {
	return bindAddress + "=" + e.bindAddress + "; " +
		cdnBaseURL + "=" + e.cdnBaseURL + "; " +
		apiBaseURL + "=" + e.apiBaseURL + "; " +
		userAgent + "=" + e.userAgent
}

func (e Env) BindAddress() string {
	return e.bindAddress
}

func (e Env) CdnBaseURL() string {
	return e.cdnBaseURL
}

func (e Env) APIBaseURL() string {
	return e.apiBaseURL
}

func (e Env) UserAgent() string {
	return e.userAgent
}
