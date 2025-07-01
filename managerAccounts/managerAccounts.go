package managerAccounts

import (
	"context"

	"github.com/Makrorof/GoAmazonAdApi"
	"github.com/Makrorof/GoAmazonAdApi/auth"
)

type LinkedAccount struct {
	MarketplaceID   string `json:"marketplaceId"`
	AccountID       string `json:"accountId"`
	AccountName     string `json:"accountName"`
	ProfileID       string `json:"profileId"`
	AccountType     string `json:"accountType"`
	DspAdvertiserID string `json:"dspAdvertiserId"`
}

type ManagerAccount struct {
	ManagerAccountName string          `json:"managerAccountName"`
	ManagerAccountID   string          `json:"managerAccountId"`
	LinkedAccounts     []LinkedAccount `json:"linkedAccounts"`
}

type ManagerAccountsData struct {
	ManagerAccounts []ManagerAccount `json:"managerAccounts"`
}
type ManagerAccounts struct {
	requestClient *auth.RequestClient
}

func New(c *auth.Client, endpoint auth.AMAZON_ENDPOINT) *ManagerAccounts {
	return &ManagerAccounts{
		requestClient: auth.NewRequestClient(c, endpoint, GoAmazonAdApi.API_ERROR_MAX_RETRY, GoAmazonAdApi.API_ERROR_RETRY_DELAY),
	}
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *ManagerAccounts) List(ctx context.Context) (ManagerAccountsData, error) {
	var managerAccounts ManagerAccountsData

	if err := p.requestClient.GET(ctx, "/managerAccounts", map[string][]string{"Content-Type": {"application/vnd.getmanageraccountsresponse.v1+json"}}, &managerAccounts); err != nil {
		return managerAccounts, err
	}

	return managerAccounts, nil
}
