package sponsoredProducts

type TargetingExpressionType string
type IncludeAnalysisType string

const (
	CLOSE_MATCH TargetingExpressionType = "CLOSE_MATCH"
	LOOSE_MATCH TargetingExpressionType = "LOOSE_MATCH"
	SUBSTITUTES TargetingExpressionType = "SUBSTITUTES"
	COMPLEMENTS TargetingExpressionType = "COMPLEMENTS"

	INCLUDE_ANALYSIS_TRUE  IncludeAnalysisType = "true"
	INCLUDE_ANALYSIS_FALSE IncludeAnalysisType = "false"
)

type IGetBidRecommendationsRequest interface {
	getNewResponse() IGetBidRecommendationsResponse
	getVersion() int
}
type IGetBidRecommendationsResponse interface {
	getVersion() int
}

// region GetBidRecommendationsV5
type GetBidRecommendationsExistingAdGroupRequestV5 struct {
	TargetingExpressions []struct {
		Type TargetingExpressionType `json:"type"`
	} `json:"targetingExpressions"`
	CampaignID         string              `json:"campaignId"`
	RecommendationType string              `json:"recommendationType"`
	IncludeAnalysis    IncludeAnalysisType `json:"includeAnalysis"`
	AdGroupID          string              `json:"adGroupId"`
}

type GetBidRecommendationsNewAdGroupRequestV5 struct {
	Asins                []string `json:"asins"`
	TargetingExpressions []struct {
		Type TargetingExpressionType `json:"type"`
	} `json:"targetingExpressions"`
	ProductDetailsList []struct {
		GlobalStoreSetting struct {
			CatalogSourceCountryCode string `json:"catalogSourceCountryCode"`
		} `json:"globalStoreSetting"`
		Asin string `json:"asin"`
	} `json:"productDetailsList"`
	Bidding struct {
		Adjustments []struct {
			Predicate  string `json:"predicate"`
			Percentage string `json:"percentage"`
		} `json:"adjustments"`
		Strategy string `json:"strategy"`
	} `json:"bidding"`
	RecommendationType string              `json:"recommendationType"`
	IncludeAnalysis    IncludeAnalysisType `json:"includeAnalysis"`
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
