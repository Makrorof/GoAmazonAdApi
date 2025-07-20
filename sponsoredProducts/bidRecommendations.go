package sponsoredProducts

type IGetBidRecommendationsRequest interface {
	getNewResponse() IGetBidRecommendationsResponse
	getVersion() int
}
type IGetBidRecommendationsResponse interface {
	getVersion() int
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

type GetBidRecommendationsResponseV5 struct {
	BidRecommendations []struct {
		Theme                              string `json:"theme,omitempty"`
		BidAnalysesForTargetingExpressions []struct {
			BidAnalyses struct {
				ALL []struct {
					Bid           string `json:"bid,omitempty"`
					Type          string `json:"type,omitempty"`
					ImpactMetrics struct {
						EstimatedImpressionAvg   string `json:"estimatedImpressionAvg,omitempty"`
						EstimatedImpressionUpper string `json:"estimatedImpressionUpper,omitempty"`
						EstimatedImpressionLower string `json:"estimatedImpressionLower,omitempty"`
					} `json:"impactMetrics,omitempty"`
				} `json:"ALL,omitempty"`
				PLACEMENTTOP []struct {
					Bid           string `json:"bid,omitempty"`
					Type          string `json:"type,omitempty"`
					ImpactMetrics struct {
						EstimatedImpressionAvg   string `json:"estimatedImpressionAvg,omitempty"`
						EstimatedImpressionUpper string `json:"estimatedImpressionUpper,omitempty"`
						EstimatedImpressionLower string `json:"estimatedImpressionLower,omitempty"`
					} `json:"impactMetrics,omitempty"`
				} `json:"PLACEMENT_TOP,omitempty"`
				PLACEMENTRESTOFSEARCH []struct {
					Bid           string `json:"bid,omitempty"`
					Type          string `json:"type,omitempty"`
					ImpactMetrics struct {
						EstimatedImpressionAvg   string `json:"estimatedImpressionAvg,omitempty"`
						EstimatedImpressionUpper string `json:"estimatedImpressionUpper,omitempty"`
						EstimatedImpressionLower string `json:"estimatedImpressionLower,omitempty"`
					} `json:"impactMetrics,omitempty"`
				} `json:"PLACEMENT_REST_OF_SEARCH,omitempty"`
				PLACEMENTPRODUCTPAGE []struct {
					Bid           string `json:"bid,omitempty"`
					Type          string `json:"type,omitempty"`
					ImpactMetrics struct {
						EstimatedImpressionAvg   string `json:"estimatedImpressionAvg,omitempty"`
						EstimatedImpressionUpper string `json:"estimatedImpressionUpper,omitempty"`
						EstimatedImpressionLower string `json:"estimatedImpressionLower,omitempty"`
					} `json:"impactMetrics,omitempty"`
				} `json:"PLACEMENT_PRODUCT_PAGE,omitempty"`
			} `json:"bidAnalyses,omitempty"`
			TargetingExpression struct {
				Type string `json:"type,omitempty"`
			} `json:"targetingExpression,omitempty"`
		} `json:"bidAnalysesForTargetingExpressions,omitempty"`
		BidRecommendationsForTargetingExpressions []struct {
			SuggestedBidImpactMetrics struct {
				EstimatedImpressionUpper string `json:"estimatedImpressionUpper,omitempty"`
				EstimatedImpressionLower string `json:"estimatedImpressionLower,omitempty"`
			} `json:"suggestedBidImpactMetrics,omitempty"`
			BidValues []struct {
				SuggestedBid string `json:"suggestedBid,omitempty"`
			} `json:"bidValues,omitempty"`
			TargetingExpression struct {
				Type string `json:"type,omitempty"`
			} `json:"targetingExpression,omitempty"`
		} `json:"bidRecommendationsForTargetingExpressions,omitempty"`
	} `json:"bidRecommendations,omitempty"`
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
