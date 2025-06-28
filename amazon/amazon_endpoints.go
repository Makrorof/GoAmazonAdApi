package amazon

import (
	"net/url"
	"path"
)

type AMAZON_ENDPOINT string

const (
	//United States (US), Canada (CA), Mexico (MX), Brazil (BR)
	AMAZON_ENDPOINTS_NA AMAZON_ENDPOINT = "https://advertising-api.amazon.com"
	//United Kingdom (UK), France (FR), Italy (IT), Spain (ES), Germany (DE), Netherlands (NL), United Arab Emirates (AE), Poland (PL), Turkey (TR), Egypt (EG), Saudi Arabia (SA), Sweden (SE), Belgium (BE), India (IN), South Africa (ZA)
	AMAZON_ENDPOINTS_EU AMAZON_ENDPOINT = "https://advertising-api-eu.amazon.com"
	//Japan (JP), Australia (AU), Singapore (SG)
	AMAZON_ENDPOINTS_FE AMAZON_ENDPOINT = "https://advertising-api-fe.amazon.com"
)

func (endpoint AMAZON_ENDPOINT) Join(apiPath string) string {
	u, err := url.Parse(string(endpoint))
	if err != nil {
		return path.Join(string(endpoint), apiPath)
	}
	u.Path = path.Join(u.Path, apiPath)
	return u.String()
}
