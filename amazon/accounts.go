package amazon

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

type Accounts struct {
	client *Client
}

func NewAccounts(client *Client) *Accounts {
	return &Accounts{
		client: client,
	}
}

func (accounts *Accounts) Profiles() ([]ProfileData, error) {
	accounts.client.

	return nil, nil
}
