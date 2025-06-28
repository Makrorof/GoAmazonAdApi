package amazon

import (
	"context"
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

// TODO: Burasi version degisimlerinden dogacak sorunlar icin interface olarak tanimlanacak AccountsV2 seklinde alt structlar uretilecek.
type Accounts struct {
	requestClient *requestClient
}

func NewAccounts(client *Client, endpoint AMAZON_ENDPOINT) *Accounts {
	return &Accounts{
		requestClient: newRequestClient(client, endpoint, API_ERROR_MAX_RETRY, API_ERROR_RETRY_DELAY),
	}
}

// Returns an error, which can be either a standard error or an AmazonError.
func (accounts *Accounts) Profiles(ctx context.Context) ([]ProfileData, error) {
	var profiles []ProfileData

	if err := accounts.requestClient.GET(ctx, "/v2/profiles", &profiles); err != nil {
		return nil, err
	}

	return profiles, nil
}
