package sponsoredProducts

import (
	"context"
	"fmt"

	"github.com/Makrorof/GoAmazonAdApi"
	"github.com/Makrorof/GoAmazonAdApi/auth"
)

//https://advertising.amazon.com/API/docs/en-us/sponsored-products/3-0/openapi/prod

type SponsoredProducts struct {
	requestClient *auth.RequestClient
}

func NewV3(c *auth.Client, endpoint auth.AMAZON_ENDPOINT) *SponsoredProducts {
	return &SponsoredProducts{
		requestClient: auth.NewRequestClient(c, endpoint, GoAmazonAdApi.API_ERROR_MAX_RETRY, GoAmazonAdApi.API_ERROR_RETRY_DELAY),
	}
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) GetBidRecommendations(ctx context.Context, request IGetBidRecommendationsRequest) (IGetBidRecommendationsResponse, error) {
	response := request.getNewResponse()

	if err := p.requestClient.POST(ctx, "/sp/targets/bid/recommendations", map[string][]string{"Content-Type": {fmt.Sprintf("application/vnd.spthemebasedbidrecommendation.v%d+json", request.getVersion())}}, request, response); err != nil {
		return response, err
	}

	return response, nil
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) GetKeywordsRecommendations(ctx context.Context, request IGetKeywordsRecommendationsRequest) (IGetKeywordsRecommendationsResponse, error) {
	response := request.getNewResponse()

	if err := p.requestClient.POST(ctx, "/sp/targets/keywords/recommendations", map[string][]string{"Content-Type": {fmt.Sprintf("application/vnd.spkeywordsrecommendation.v%d+json", request.getVersion())}}, request, response); err != nil {
		return response, err
	}

	return response, nil
}
