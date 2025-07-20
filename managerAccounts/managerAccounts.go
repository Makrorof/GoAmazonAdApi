package managerAccounts

import (
	"context"

	"github.com/Makrorof/GoAmazonAdApi"
	"github.com/Makrorof/GoAmazonAdApi/auth"
)

type LinkedAccount struct {
	MarketplaceID   string `json:"marketplaceId,omitempty"`
	AccountID       string `json:"accountId,omitempty"`
	AccountName     string `json:"accountName,omitempty"`
	ProfileID       string `json:"profileId,omitempty"`
	AccountType     string `json:"accountType,omitempty"`
	DspAdvertiserID string `json:"dspAdvertiserId,omitempty"`
}

type ManagerAccount struct {
	ManagerAccountName string          `json:"managerAccountName,omitempty"`
	ManagerAccountID   string          `json:"managerAccountId,omitempty"`
	LinkedAccounts     []LinkedAccount `json:"linkedAccounts,omitempty"`
}

type ManagerAccountsData struct {
	ManagerAccounts []ManagerAccount `json:"managerAccounts,omitempty"`
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
