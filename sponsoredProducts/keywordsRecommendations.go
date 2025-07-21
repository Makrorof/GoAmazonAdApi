package sponsoredProducts

type IGetKeywordsRecommendationsRequest interface {
	getNewResponse() IGetKeywordsRecommendationsResponse
	getVersion() int
}
type IGetKeywordsRecommendationsResponse interface {
	getVersion() int
}

type LowUpMetrics struct {
	Lower int `json:"lower,omitempty"`
	Upper int `json:"upper,omitempty"`
}
type LowUpMetricsValues struct {
	Values []LowUpMetrics `json:"values,omitempty"`
}
type ClicksOrdersMetrics struct {
	Clicks *LowUpMetricsValues `json:"clicks,omitempty"`
	Orders *LowUpMetricsValues `json:"orders,omitempty"`
}

// region GetKeywordsRecommendationsV5
type GetKeywordsRecommendationsAdGroupRequestV5 struct {
	CampaignID         string          `json:"campaignId,omitempty"`
	RecommendationType string          `json:"recommendationType,omitempty"`
	BidsEnabled        string          `json:"bidsEnabled,omitempty"`
	AdGroupID          string          `json:"adGroupId,omitempty"`
	Targets            []TargetsFilter `json:"targets,omitempty"`
	MaxRecommendations string          `json:"maxRecommendations,omitempty"`
	SortDimension      string          `json:"sortDimension,omitempty"`
	Locale             string          `json:"locale,omitempty"`
}

type GetKeywordsRecommendationsAsinsRequestV5 struct {
	Asins              []string                   `json:"asins,omitempty"`
	ProductDetailsList []ProductDetailsListFilter `json:"productDetailsList,omitempty"`
	BiddingStrategy    string                     `json:"biddingStrategy,omitempty"`
	RecommendationType string                     `json:"recommendationType,omitempty"`
	BidsEnabled        string                     `json:"bidsEnabled,omitempty"`
	Targets            []TargetsFilter            `json:"targets,omitempty"`
	MaxRecommendations string                     `json:"maxRecommendations,omitempty"`
	SortDimension      string                     `json:"sortDimension,omitempty"`
	Locale             string                     `json:"locale,omitempty"`
}
type GetKeywordV5SuggestedBid struct {
	Suggested  float64 `json:"suggested,omitempty"`
	RangeStart float64 `json:"rangeStart,omitempty"`
	RangeEnd   float64 `json:"rangeEnd,omitempty"`
}
type GetKeywordV5BidInfo struct {
	SuggestedBid *GetKeywordV5SuggestedBid `json:"suggestedBid,omitempty"`
	MatchType    string                    `json:"matchType,omitempty"`
	Rank         int                       `json:"rank,omitempty"`
	Theme        string                    `json:"theme,omitempty"`
	Bid          int                       `json:"bid,omitempty"`
}
type GetKeywordTargetV5 struct {
	SearchTermImpressionShare int                   `json:"searchTermImpressionShare,omitempty"`
	Translation               string                `json:"translation,omitempty"`
	BidInfo                   []GetKeywordV5BidInfo `json:"bidInfo,omitempty"`
	SearchTermImpressionRank  int                   `json:"searchTermImpressionRank,omitempty"`
	Keyword                   string                `json:"keyword,omitempty"`
	UserSelectedKeyword       bool                  `json:"userSelectedKeyword,omitempty"`
	RecID                     string                `json:"recId,omitempty"`
}
type GetKeywordsRecommendationsResponseV5 struct {
	KeywordTargetList []GetKeywordTargetV5  `json:"keywordTargetList,omitempty"`
	ImpactMetrics     []ClicksOrdersMetrics `json:"impactMetrics,omitempty"`
}

func (r *GetKeywordsRecommendationsAdGroupRequestV5) getNewResponse() IGetKeywordsRecommendationsResponse {
	return new(GetKeywordsRecommendationsResponseV5)
}
func (r *GetKeywordsRecommendationsAdGroupRequestV5) getVersion() int {
	return 5
}

func (r *GetKeywordsRecommendationsResponseV5) getVersion() int {
	return 5
}

//endregion GetKeywordsRecommendationsV5
