package sponsoredProducts

type IGetBidRecommendationsRequest interface {
	getNewResponse() IGetBidRecommendationsResponse
	getVersion() int
}
type IGetBidRecommendationsResponse interface {
	getVersion() int
}

// region GetBidRecommendationsV5
type GetBidRecommendationsRequestV5 struct {
	TargetingExpressions []struct {
		Type string `json:"type"`
	} `json:"targetingExpressions"`
	CampaignID         string `json:"campaignId"`
	RecommendationType string `json:"recommendationType"`
	IncludeAnalysis    string `json:"includeAnalysis"`
	AdGroupID          string `json:"adGroupId"`
}

type GetBidRecommendationsResponseV5 struct {
	BidRecommendations []struct {
		Theme                              string `json:"theme"`
		BidAnalysesForTargetingExpressions []struct {
			BidAnalyses struct {
				ALL []struct {
					Bid           string `json:"bid"`
					Type          string `json:"type"`
					ImpactMetrics struct {
						EstimatedImpressionAvg   string `json:"estimatedImpressionAvg"`
						EstimatedImpressionUpper string `json:"estimatedImpressionUpper"`
						EstimatedImpressionLower string `json:"estimatedImpressionLower"`
					} `json:"impactMetrics"`
				} `json:"ALL"`
				PLACEMENTTOP []struct {
					Bid           string `json:"bid"`
					Type          string `json:"type"`
					ImpactMetrics struct {
						EstimatedImpressionAvg   string `json:"estimatedImpressionAvg"`
						EstimatedImpressionUpper string `json:"estimatedImpressionUpper"`
						EstimatedImpressionLower string `json:"estimatedImpressionLower"`
					} `json:"impactMetrics"`
				} `json:"PLACEMENT_TOP"`
				PLACEMENTRESTOFSEARCH []struct {
					Bid           string `json:"bid"`
					Type          string `json:"type"`
					ImpactMetrics struct {
						EstimatedImpressionAvg   string `json:"estimatedImpressionAvg"`
						EstimatedImpressionUpper string `json:"estimatedImpressionUpper"`
						EstimatedImpressionLower string `json:"estimatedImpressionLower"`
					} `json:"impactMetrics"`
				} `json:"PLACEMENT_REST_OF_SEARCH"`
				PLACEMENTPRODUCTPAGE []struct {
					Bid           string `json:"bid"`
					Type          string `json:"type"`
					ImpactMetrics struct {
						EstimatedImpressionAvg   string `json:"estimatedImpressionAvg"`
						EstimatedImpressionUpper string `json:"estimatedImpressionUpper"`
						EstimatedImpressionLower string `json:"estimatedImpressionLower"`
					} `json:"impactMetrics"`
				} `json:"PLACEMENT_PRODUCT_PAGE"`
			} `json:"bidAnalyses"`
			TargetingExpression struct {
				Type string `json:"type"`
			} `json:"targetingExpression"`
		} `json:"bidAnalysesForTargetingExpressions"`
		BidRecommendationsForTargetingExpressions []struct {
			SuggestedBidImpactMetrics struct {
				EstimatedImpressionUpper string `json:"estimatedImpressionUpper"`
				EstimatedImpressionLower string `json:"estimatedImpressionLower"`
			} `json:"suggestedBidImpactMetrics"`
			BidValues []struct {
				SuggestedBid string `json:"suggestedBid"`
			} `json:"bidValues"`
			TargetingExpression struct {
				Type string `json:"type"`
			} `json:"targetingExpression"`
		} `json:"bidRecommendationsForTargetingExpressions"`
	} `json:"bidRecommendations"`
}

func (r *GetBidRecommendationsRequestV5) getNewResponse() IGetBidRecommendationsResponse {
	return new(GetBidRecommendationsResponseV5)
}

func (r *GetBidRecommendationsRequestV5) getVersion() int {
	return 5
}

func (r *GetBidRecommendationsResponseV5) getVersion() int {
	return 5
}

//endregion GetBidRecommendationsV5
