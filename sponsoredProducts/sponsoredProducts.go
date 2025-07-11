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

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) ListKeywords(ctx context.Context, request IListKeywordsRequest) (IListKeywordsResponse, error) {
	response := request.getNewResponse()

	if err := p.requestClient.POST(ctx, "/sp/keywords/list", map[string][]string{"Content-Type": {fmt.Sprintf("application/vnd.spKeyword.v%d+json", request.getVersion())}}, request, response); err != nil {
		return response, err
	}

	return response, nil
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) CreateKeywords(ctx context.Context, request ICreateKeywordsRequest) (ICreateKeywordsResponse, error) {
	response := request.getNewResponse()

	if err := p.requestClient.POST(ctx, "/sp/keywords", map[string][]string{"Content-Type": {fmt.Sprintf("application/vnd.spKeyword.v%d+json", request.getVersion())}}, request, response); err != nil {
		return response, err
	}

	return response, nil
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) DeleteKeywords(ctx context.Context, request IDeleteKeywordsRequest) (IDeleteKeywordsResponse, error) {
	response := request.getNewResponse()

	if err := p.requestClient.POST(ctx, "/sp/keywords", map[string][]string{"Content-Type": {fmt.Sprintf("application/vnd.spKeyword.v%d+json", request.getVersion())}}, request, response); err != nil {
		return response, err
	}

	return response, nil
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) UpdateKeywords(ctx context.Context, request IUpdateKeywordsRequest) (IUpdateKeywordsResponse, error) {
	response := request.getNewResponse()

	if err := p.requestClient.PUT(ctx, "/sp/keywords", map[string][]string{"Content-Type": {fmt.Sprintf("application/vnd.spKeyword.v%d+json", request.getVersion())}}, request, response); err != nil {
		return response, err
	}

	return response, nil
}
