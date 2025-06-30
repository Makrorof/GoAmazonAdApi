package profiles

import (
	"context"
	"fmt"

	"github.com/Makrorof/GoAmazonAdApi"
	"github.com/Makrorof/GoAmazonAdApi/auth"
)

type ProfileData struct {
	ProfileID    int64  `json:"profileId"`
	CountryCode  string `json:"countryCode"`
	CurrencyCode string `json:"currencyCode"`
	Timezone     string `json:"timezone"`
	AccountInfo  struct {
		MarketplaceStringID string `json:"marketplaceStringId"`
		ID                  string `json:"id"`
		Type                string `json:"type"`
		Name                string `json:"name"`
		ValidPaymentMethod  bool   `json:"validPaymentMethod"`
	} `json:"accountInfo"`
}

type Profiles struct {
	requestClient *auth.RequestClient
}

func New(c *auth.Client, endpoint auth.AMAZON_ENDPOINT) *Profiles {
	return &Profiles{
		requestClient: auth.NewRequestClient(c, endpoint, GoAmazonAdApi.API_ERROR_MAX_RETRY, GoAmazonAdApi.API_ERROR_RETRY_DELAY),
	}
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *Profiles) List(ctx context.Context) ([]ProfileData, error) {
	var profiles []ProfileData

	if err := p.requestClient.GET(ctx, "/v2/profiles", &profiles, nil); err != nil {
		return nil, err
	}

	return profiles, nil
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *Profiles) GetProfile(ctx context.Context, profileId int) (ProfileData, error) {
	var profile ProfileData

	if err := p.requestClient.GET(ctx, fmt.Sprintf("/v2/profiles/%d", profileId), &profile, nil); err != nil {
		return profile, err
	}

	return profile, nil
}
