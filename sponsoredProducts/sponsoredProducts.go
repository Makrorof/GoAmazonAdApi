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

//region keywords

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

//endregion keywords

//region campaigns

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) CreateCampaigns(ctx context.Context, request ICreateCampaignsRequest, returnRepresentation bool) (ICreateCampaignsResponse, error) {
	response := request.getNewResponse()

	header := map[string][]string{
		"Content-Type": {fmt.Sprintf("application/vnd.spCampaign.v%d+json", request.getVersion())},
	}

	if returnRepresentation {
		header["Prefer"] = []string{"return=representation"}
	}

	if err := p.requestClient.POST(ctx, "/sp/campaigns", header, request, response); err != nil {
		return response, err
	}

	return response, nil
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) UpdateCampaigns(ctx context.Context, request IUpdateCampaignsRequest, returnRepresentation bool) (IUpdateCampaignsResponse, error) {
	response := request.getNewResponse()

	header := map[string][]string{
		"Content-Type": {fmt.Sprintf("application/vnd.spCampaign.v%d+json", request.getVersion())},
	}

	if returnRepresentation {
		header["Prefer"] = []string{"return=representation"}
	}

	if err := p.requestClient.PUT(ctx, "/sp/campaigns", header, request, response); err != nil {
		return response, err
	}

	return response, nil
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) DeleteCampaigns(ctx context.Context, request IDeleteCampaignsRequest) (IDeleteCampaignsResponse, error) {
	response := request.getNewResponse()

	if err := p.requestClient.POST(ctx, "/sp/campaigns/delete", map[string][]string{"Content-Type": {fmt.Sprintf("application/vnd.spCampaign.v%d+json", request.getVersion())}}, request, response); err != nil {
		return response, err
	}

	return response, nil
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) ListCampaigns(ctx context.Context, request IListCampaignsRequest) (IListCampaignsResponse, error) {
	response := request.getNewResponse()

	if err := p.requestClient.POST(ctx, "/sp/campaigns/list", map[string][]string{"Content-Type": {fmt.Sprintf("application/vnd.spCampaign.v%d+json", request.getVersion())}}, request, response); err != nil {
		return response, err
	}

	return response, nil
}

//endregion campaigns

//region adGroups

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) CreateAdGroups(ctx context.Context, request ICreateAdGroupsRequest, returnRepresentation bool) (ICreateAdGroupsResponse, error) {
	response := request.getNewResponse()

	header := map[string][]string{
		"Content-Type": {fmt.Sprintf("application/vnd.spAdGroup.v%d+json", request.getVersion())},
	}

	if returnRepresentation {
		header["Prefer"] = []string{"return=representation"}
	}

	if err := p.requestClient.POST(ctx, "/sp/adGroups", header, request, response); err != nil {
		return response, err
	}

	return response, nil
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) UpdateAdGroups(ctx context.Context, request IUpdateAdGroupsRequest, returnRepresentation bool) (IUpdateAdGroupsResponse, error) {
	response := request.getNewResponse()

	header := map[string][]string{
		"Content-Type": {fmt.Sprintf("application/vnd.spAdGroup.v%d+json", request.getVersion())},
	}

	if returnRepresentation {
		header["Prefer"] = []string{"return=representation"}
	}

	if err := p.requestClient.PUT(ctx, "/sp/adGroups", header, request, response); err != nil {
		return response, err
	}

	return response, nil
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) DeleteAdGroups(ctx context.Context, request IDeleteAdGroupsRequest) (IDeleteAdGroupsResponse, error) {
	response := request.getNewResponse()

	if err := p.requestClient.POST(ctx, "/sp/adGroups/delete", map[string][]string{"Content-Type": {fmt.Sprintf("application/vnd.spAdGroup.v%d+json", request.getVersion())}}, request, response); err != nil {
		return response, err
	}

	return response, nil
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) ListAdGroups(ctx context.Context, request IListAdGroupsRequest) (IListAdGroupsResponse, error) {
	response := request.getNewResponse()

	if err := p.requestClient.POST(ctx, "/sp/adGroups/list", map[string][]string{"Content-Type": {fmt.Sprintf("application/vnd.spAdGroup.v%d+json", request.getVersion())}}, request, response); err != nil {
		return response, err
	}

	return response, nil
}

//endregion adGroups

//region targetingClauses

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) CreateTargetingClauses(ctx context.Context, request ICreateTargetingClausesRequest, returnRepresentation bool) (ICreateTargetingClausesResponse, error) {
	response := request.getNewResponse()

	header := map[string][]string{
		"Content-Type": {fmt.Sprintf("application/vnd.spTargetingClause.v%d+json", request.getVersion())},
	}

	if returnRepresentation {
		header["Prefer"] = []string{"return=representation"}
	}

	if err := p.requestClient.POST(ctx, "/sp/targets", header, request, response); err != nil {
		return response, err
	}

	return response, nil
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) UpdateTargetingClauses(ctx context.Context, request IUpdateTargetingClausesRequest, returnRepresentation bool) (IUpdateTargetingClausesResponse, error) {
	response := request.getNewResponse()

	header := map[string][]string{
		"Content-Type": {fmt.Sprintf("application/vnd.spTargetingClause.v%d+json", request.getVersion())},
	}

	if returnRepresentation {
		header["Prefer"] = []string{"return=representation"}
	}

	if err := p.requestClient.PUT(ctx, "/sp/targets", header, request, response); err != nil {
		return response, err
	}

	return response, nil
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) DeleteTargetingClauses(ctx context.Context, request IDeleteTargetingClausesRequest) (IDeleteTargetingClausesResponse, error) {
	response := request.getNewResponse()

	if err := p.requestClient.POST(ctx, "/sp/targets/delete", map[string][]string{"Content-Type": {fmt.Sprintf("application/vnd.spTargetingClause.v%d+json", request.getVersion())}}, request, response); err != nil {
		return response, err
	}

	return response, nil
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (p *SponsoredProducts) ListTargetingClauses(ctx context.Context, request IListTargetingClausesRequest) (IListTargetingClausesResponse, error) {
	response := request.getNewResponse()

	if err := p.requestClient.POST(ctx, "/sp/targets/list", map[string][]string{"Content-Type": {fmt.Sprintf("application/vnd.spTargetingClause.v%d+json", request.getVersion())}}, request, response); err != nil {
		return response, err
	}

	return response, nil
}

//endregion targetingClauses
