package env

import (
	"log"
	"net/url"
	"os"
)

type Env struct {
	bindAddress string
	cdnBaseUrl  string
	apiBaseUrl  string
	userAgent   string
}

const defaultBindAddress string = ":8080"
const defaultCdnBaseUrl string = "https://img.chefkoch-cdn.de/"
const defaultApiBaseUrl string = "https://api.chefkoch.de/"
const defaultUserAgent string = "Mozilla/5.0 (X11; Linux x86_64; rv:108.0) Gecko/20100101 Firefox/108.0"

const bindAddress string = "BIND_ADDRESS"
const cdnBaseUrl string = "CDN_BASE_URL"
const apiBaseUrl string = "API_BASE_URL"
const userAgent string = "USER_AGENT"

func NewEnv() *Env {

	env := &Env{
		bindAddress: os.Getenv(bindAddress),
		cdnBaseUrl:  os.Getenv(cdnBaseUrl),
		apiBaseUrl:  os.Getenv(apiBaseUrl),
		userAgent:   os.Getenv(userAgent),
	}

	if len(env.bindAddress) == 0 {
		log.Println("Using", defaultBindAddress, "as", bindAddress)
		env.bindAddress = defaultBindAddress
	}

	if _, err := url.Parse(env.cdnBaseUrl); err != nil || len(env.cdnBaseUrl) == 0 {
		log.Println("Using", defaultCdnBaseUrl, "as", cdnBaseUrl)
		env.cdnBaseUrl = defaultCdnBaseUrl
	}

	if _, err := url.Parse(env.apiBaseUrl); err != nil || len(env.apiBaseUrl) == 0 {
		log.Println("Using", defaultApiBaseUrl, "as", apiBaseUrl)
		env.apiBaseUrl = defaultApiBaseUrl
	}

	if len(env.userAgent) == 0 {
		env.userAgent = defaultUserAgent
	}

	return env
}

func (e Env) String() string {
	return bindAddress + "=" + e.bindAddress + "; " +
		cdnBaseUrl + "=" + e.cdnBaseUrl + "; " +
		apiBaseUrl + "=" + e.apiBaseUrl + "; " +
		userAgent + "=" + e.userAgent
}

func (e Env) BindAddress() string {
	return e.bindAddress
}

func (e Env) CdnBaseUrl() string {
	return e.cdnBaseUrl
}

func (e Env) ApiBaseUrl() string {
	return e.apiBaseUrl
}

func (e Env) UserAgent() string {
	return e.userAgent
}
