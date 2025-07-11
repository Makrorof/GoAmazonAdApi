package sponsoredProducts

type IGetKeywordsRecommendationsRequest interface {
	getNewResponse() IGetKeywordsRecommendationsResponse
	getVersion() int
}
type IGetKeywordsRecommendationsResponse interface {
	getVersion() int
}

// region GetKeywordsRecommendationsV5
type GetKeywordsRecommendationsAdGroupRequestV5 struct {
	CampaignID         string `json:"campaignId"`
	RecommendationType string `json:"recommendationType"`
	BidsEnabled        string `json:"bidsEnabled"`
	AdGroupID          string `json:"adGroupId"`
	Targets            []struct {
		MatchType           string  `json:"matchType"`
		Keyword             string  `json:"keyword"`
		Bid                 float64 `json:"bid"`
		UserSelectedKeyword bool    `json:"userSelectedKeyword"`
	} `json:"targets"`
	MaxRecommendations string `json:"maxRecommendations"`
	SortDimension      string `json:"sortDimension"`
	Locale             string `json:"locale"`
}

type GetKeywordsRecommendationsAsinsRequestV5 struct {
	Asins              []string `json:"asins"`
	ProductDetailsList []struct {
		GlobalStoreSetting struct {
			CatalogSourceCountryCode string `json:"catalogSourceCountryCode"`
		} `json:"globalStoreSetting"`
		Asin string `json:"asin"`
	} `json:"productDetailsList"`
	BiddingStrategy    string `json:"biddingStrategy"`
	RecommendationType string `json:"recommendationType"`
	BidsEnabled        string `json:"bidsEnabled"`
	Targets            []struct {
		MatchType           string  `json:"matchType"`
		Keyword             string  `json:"keyword"`
		Bid                 float64 `json:"bid"`
		UserSelectedKeyword bool    `json:"userSelectedKeyword"`
	} `json:"targets"`
	MaxRecommendations string `json:"maxRecommendations"`
	SortDimension      string `json:"sortDimension"`
	Locale             string `json:"locale"`
}

type GetKeywordsRecommendationsResponseV5 struct {
	KeywordTargetList []struct {
		SearchTermImpressionShare int    `json:"searchTermImpressionShare"`
		Translation               string `json:"translation"`
		BidInfo                   []struct {
			SuggestedBid struct {
				Suggested  float64 `json:"suggested"`
				RangeStart float64 `json:"rangeStart"`
				RangeEnd   float64 `json:"rangeEnd"`
			} `json:"suggestedBid"`
			MatchType string `json:"matchType"`
			Rank      int    `json:"rank"`
			Theme     string `json:"theme"`
			Bid       int    `json:"bid"`
		} `json:"bidInfo"`
		SearchTermImpressionRank int    `json:"searchTermImpressionRank"`
		Keyword                  string `json:"keyword"`
		UserSelectedKeyword      bool   `json:"userSelectedKeyword"`
		RecID                    string `json:"recId"`
	} `json:"keywordTargetList"`
	ImpactMetrics []struct {
		Clicks struct {
			Values []struct {
				Lower int `json:"lower"`
				Upper int `json:"upper"`
			} `json:"values"`
		} `json:"clicks"`
		Orders struct {
			Values []struct {
				Lower int `json:"lower"`
				Upper int `json:"upper"`
			} `json:"values"`
		} `json:"orders"`
	} `json:"impactMetrics"`
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
