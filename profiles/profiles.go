package profiles

import (
	"context"
	"fmt"

	"github.com/Makrorof/GoAmazonAdApi"
	"github.com/Makrorof/GoAmazonAdApi/auth"
)

type ProfileData struct {
	ProfileID    int64  `json:"profileId,omitempty"`
	CountryCode  string `json:"countryCode,omitempty"`
	CurrencyCode string `json:"currencyCode,omitempty"`
	Timezone     string `json:"timezone,omitempty"`
	AccountInfo  struct {
		MarketplaceStringID string `json:"marketplaceStringId,omitempty"`
		ID                  string `json:"id,omitempty"`
		Type                string `json:"type,omitempty"`
		Name                string `json:"name,omitempty"`
		ValidPaymentMethod  bool   `json:"validPaymentMethod,omitempty"`
	} `json:"accountInfo,omitempty"`
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

	if err := p.requestClient.GET(ctx, "/v2/profiles", nil, &profiles); err != nil {
		return nil, err
	}

	return profiles, nil
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *Profiles) GetProfile(ctx context.Context, profileId int) (ProfileData, error) {
	var profile ProfileData

	if err := p.requestClient.GET(ctx, fmt.Sprintf("/v2/profiles/%d", profileId), nil, &profile); err != nil {
		return profile, err
	}

	return profile, nil
}
