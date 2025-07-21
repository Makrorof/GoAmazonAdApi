package sponsoredProducts

type IGetBidRecommendationsRequest interface {
	getNewResponse() IGetBidRecommendationsResponse
	getVersion() int
}
type IGetBidRecommendationsResponse interface {
	getVersion() int
}

type ImpactMetrics struct {
	EstimatedImpressionAvg   string `json:"estimatedImpressionAvg,omitempty"`
	EstimatedImpressionUpper string `json:"estimatedImpressionUpper,omitempty"`
	EstimatedImpressionLower string `json:"estimatedImpressionLower,omitempty"`
}
type SuggestedBidImpactMetrics struct {
	EstimatedImpressionUpper string `json:"estimatedImpressionUpper,omitempty"`
	EstimatedImpressionLower string `json:"estimatedImpressionLower,omitempty"`
}
type BidValues struct {
	SuggestedBid string `json:"suggestedBid,omitempty"`
}
type BidMetrics struct {
	Bid           string         `json:"bid,omitempty"`
	Type          string         `json:"type,omitempty"`
	ImpactMetrics *ImpactMetrics `json:"impactMetrics,omitempty"`
}

type BidAnalyses struct {
	ALL                   []BidMetrics `json:"ALL,omitempty"`
	PLACEMENTTOP          []BidMetrics `json:"PLACEMENT_TOP,omitempty"`
	PLACEMENTRESTOFSEARCH []BidMetrics `json:"PLACEMENT_REST_OF_SEARCH,omitempty"`
	PLACEMENTPRODUCTPAGE  []BidMetrics `json:"PLACEMENT_PRODUCT_PAGE,omitempty"`
}

type BidAnalysesForTargetingExpression struct {
	BidAnalyses         *BidAnalyses `json:"bidAnalyses,omitempty"`
	TargetingExpression struct {
		Type string `json:"type,omitempty"`
	} `json:"targetingExpression,omitempty"`
}

type BidRecommendationsForTargetingExpression struct {
	SuggestedBidImpactMetrics *SuggestedBidImpactMetrics `json:"suggestedBidImpactMetrics,omitempty"`
	BidValues                 []BidValues                `json:"bidValues,omitempty"`
	TargetingExpression       *TargetingExpression       `json:"targetingExpression,omitempty"`
}

// region GetBidRecommendationsV5
type GetBidRecommendationsExistingAdGroupRequestV5 struct {
	TargetingExpressions []TargetingExpression `json:"targetingExpressions,omitempty"`
	CampaignID           string                `json:"campaignId,omitempty"`
	RecommendationType   string                `json:"recommendationType,omitempty"`
	IncludeAnalysis      IncludeAnalysisType   `json:"includeAnalysis,omitempty"`
	AdGroupID            string                `json:"adGroupId,omitempty"`
}

type GetBidRecommendationsNewAdGroupRequestV5 struct {
	Asins                []string                   `json:"asins,omitempty"`
	TargetingExpressions []TargetingExpression      `json:"targetingExpressions,omitempty"`
	ProductDetailsList   []ProductDetailsListFilter `json:"productDetailsList,omitempty"`
	Bidding              *BiddingFilter             `json:"bidding,omitempty"`
	RecommendationType   string                     `json:"recommendationType,omitempty"`
	IncludeAnalysis      IncludeAnalysisType        `json:"includeAnalysis,omitempty"`
}
type GetBidRecommendationsV5 struct {
	Theme                                     string                                     `json:"theme,omitempty"`
	BidAnalysesForTargetingExpressions        []BidAnalysesForTargetingExpression        `json:"bidAnalysesForTargetingExpressions,omitempty"`
	BidRecommendationsForTargetingExpressions []BidRecommendationsForTargetingExpression `json:"bidRecommendationsForTargetingExpressions,omitempty"`
}
type GetBidRecommendationsResponseV5 struct {
	BidRecommendations []GetBidRecommendationsV5 `json:"bidRecommendations,omitempty"`
}

// BIDS_FOR_NEW_AD_GROUP
func (r *GetBidRecommendationsExistingAdGroupRequestV5) getNewResponse() IGetBidRecommendationsResponse {
	return new(GetBidRecommendationsResponseV5)
}

func (r *GetBidRecommendationsExistingAdGroupRequestV5) getVersion() int {
	return 5
}

// BIDS_FOR_EXISTING_AD_GROUP
func (r *GetBidRecommendationsNewAdGroupRequestV5) getNewResponse() IGetBidRecommendationsResponse {
	return new(GetBidRecommendationsResponseV5)
}

func (r *GetBidRecommendationsNewAdGroupRequestV5) getVersion() int {
	return 5
}

func (r *GetBidRecommendationsResponseV5) getVersion() int {
	return 5
}

//endregion GetBidRecommendationsV5
