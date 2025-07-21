package sponsoredProducts

type IListCampaignsRequest interface {
	getNewResponse() IListCampaignsResponse
	getVersion() int
}
type IListCampaignsResponse interface {
	getVersion() int
}

type ICreateCampaignsRequest interface {
	getNewResponse() ICreateCampaignsResponse
	getVersion() int
}
type ICreateCampaignsResponse interface {
	getVersion() int
}

type IDeleteCampaignsRequest interface {
	getNewResponse() IDeleteCampaignsResponse
	getVersion() int
}
type IDeleteCampaignsResponse interface {
	getVersion() int
}

type IUpdateCampaignsRequest interface {
	getNewResponse() IUpdateCampaignsResponse
	getVersion() int
}
type IUpdateCampaignsResponse interface {
	getVersion() int
}

type Budget struct {
	BudgetType      string  `json:"budgetType,omitempty"`
	Budget          float64 `json:"budget,omitempty"`
	EffectiveBudget float64 `json:"effectiveBudget,omitempty"`
}

type CampaignV3AudienceSegments struct {
	AudienceID          string `json:"audienceId,omitempty"`
	AudienceSegmentType string `json:"audienceSegmentType,omitempty"`
}
type CampaignV3ShopperCohortBidding struct {
	ShopperCohortType string                       `json:"shopperCohortType,omitempty"`
	Percentage        int                          `json:"percentage,omitempty"`
	AudienceSegments  []CampaignV3AudienceSegments `json:"audienceSegments,omitempty"`
}
type CampaignV3PlacementBidding struct {
	Percentage int    `json:"percentage,omitempty"`
	Placement  string `json:"placement,omitempty"`
}
type CampaignV3DynamicBidding struct {
	ShopperCohortBidding []CampaignV3ShopperCohortBidding `json:"shopperCohortBidding,omitempty"`
	PlacementBidding     []CampaignV3PlacementBidding     `json:"placementBidding,omitempty"`
	Strategy             string                           `json:"strategy,omitempty"`
}

type CampaignV3 struct {
	SiteRestrictions  []string `json:"siteRestrictions,omitempty"`
	OffAmazonSettings struct {
		OffAmazonBudgetControlStrategy string `json:"offAmazonBudgetControlStrategy,omitempty"`
	} `json:"offAmazonSettings,omitempty"`
	EndDate        string                    `json:"endDate,omitempty"`
	CampaignID     string                    `json:"campaignId,omitempty"`
	DynamicBidding *CampaignV3DynamicBidding `json:"dynamicBidding,omitempty"`
	Tags           map[string]string         `json:"tags,omitempty"`
	PortfolioID    string                    `json:"portfolioId,omitempty"`
	Name           string                    `json:"name,omitempty"`
	TargetingType  string                    `json:"targetingType,omitempty"`
	State          string                    `json:"state,omitempty"`
	StartDate      string                    `json:"startDate,omitempty"`
	Budget         *Budget                   `json:"budget,omitempty"`
	ExtendedData   *ExtendedData             `json:"extendedData,omitempty"`
}

// region CreateV3
type CreateCampaignsV3CampaignData struct {
	SiteRestrictions  []string                `json:"siteRestrictions,omitempty"`
	PortfolioID       string                  `json:"portfolioId,omitempty"`
	OffAmazonSettings *OffAmazonSettingFilter `json:"offAmazonSettings,omitempty"`
	EndDate           string                  `json:"endDate,omitempty"`
	Name              string                  `json:"name,omitempty"`
	TargetingType     string                  `json:"targetingType,omitempty"`
	State             string                  `json:"state,omitempty"`
	DynamicBidding    *DynamicBiddingFilter   `json:"dynamicBidding,omitempty"`
	StartDate         string                  `json:"startDate,omitempty"`
	Budget            *BudgetFilter           `json:"budget,omitempty"`
	Tags              map[string]string       `json:"tags,omitempty"`
}

type CreateCampaignsV3Success struct {
	CampaignID string      `json:"campaignId,omitempty"`
	Index      int         `json:"index,omitempty"`
	Campaign   *CampaignV3 `json:"campaign,omitempty"`
}
type CreateCampaignsRequestV3 struct {
	Campaigns []CreateCampaignsV3CampaignData `json:"campaigns,omitempty"`
}
type CreateCampaignsResponseV3 struct {
	Campaigns struct {
		Success []CreateCampaignsV3Success `json:"success,omitempty"`
		Error   []Error                    `json:"error,omitempty"`
	} `json:"campaigns,omitempty"`
}

func (r *CreateCampaignsRequestV3) getNewResponse() ICreateCampaignsResponse {
	return new(CreateCampaignsResponseV3)
}

func (r *CreateCampaignsRequestV3) getVersion() int {
	return 3
}

func (r *CreateCampaignsResponseV3) getVersion() int {
	return 3
}

//endregion CreateV3

// region UpdateV3
type UpdateCampaignsRequestV3Data struct {
	SiteRestrictions  []string                `json:"siteRestrictions,omitempty"`
	PortfolioID       string                  `json:"portfolioId,omitempty"`
	OffAmazonSettings *OffAmazonSettingFilter `json:"offAmazonSettings,omitempty"`
	EndDate           string                  `json:"endDate,omitempty"`
	CampaignID        string                  `json:"campaignId,omitempty"`
	Name              string                  `json:"name,omitempty"`
	TargetingType     string                  `json:"targetingType,omitempty"`
	State             string                  `json:"state,omitempty"`
	DynamicBidding    *DynamicBiddingFilter   `json:"dynamicBidding,omitempty"`
	StartDate         string                  `json:"startDate,omitempty"`
	Budget            *BudgetFilter           `json:"budget,omitempty"`
	Tags              map[string]string       `json:"tags,omitempty"`
}
type UpdateCampaignsRequestV3 struct {
	Campaigns []UpdateCampaignsRequestV3Data `json:"campaigns,omitempty"`
}
type UpdateCampaignsResponseV3Success struct {
	CampaignID string      `json:"campaignId,omitempty"`
	Index      int         `json:"index,omitempty"`
	Campaign   *CampaignV3 `json:"campaign,omitempty"`
}
type UpdateCampaignsResponseV3 struct {
	Campaigns struct {
		Success []UpdateCampaignsResponseV3Success `json:"success,omitempty"`
		Error   []Error                            `json:"error,omitempty"`
	} `json:"campaigns,omitempty"`
}

func (r *UpdateCampaignsRequestV3) getNewResponse() IUpdateCampaignsResponse {
	return new(UpdateCampaignsResponseV3)
}

func (r *UpdateCampaignsRequestV3) getVersion() int {
	return 3
}

func (r *UpdateCampaignsResponseV3) getVersion() int {
	return 3
}

//endregion UpdateV3

// region DeleteV3
type DeleteCampaignsRequestV3 struct {
	CampaignIDFilter *IDFilter `json:"campaignIdFilter,omitempty"`
}
type DeleteCampaignsResponseV3Success struct {
	CampaignID string      `json:"campaignId,omitempty"`
	Index      int         `json:"index,omitempty"`
	Campaign   *CampaignV3 `json:"campaign,omitempty"`
}
type DeleteCampaignsResponseV3 struct {
	Campaigns struct {
		Success []DeleteCampaignsResponseV3Success `json:"success,omitempty"`
		Error   []Error                            `json:"error,omitempty"`
	} `json:"campaigns,omitempty"`
}

func (r *DeleteCampaignsRequestV3) getNewResponse() IDeleteCampaignsResponse {
	return new(DeleteCampaignsResponseV3)
}

func (r *DeleteCampaignsRequestV3) getVersion() int {
	return 3
}

func (r *DeleteCampaignsResponseV3) getVersion() int {
	return 3
}

//endregion DeleteV3

// region ListV3
type ListCampaignsRequestV3 struct {
	CampaignIDFilter          *IDFilter   `json:"campaignIdFilter,omitempty"`
	PortfolioIDFilter         *IDFilter   `json:"portfolioIdFilter,omitempty"`
	StateFilter               *IDFilter   `json:"stateFilter,omitempty"`
	MaxResults                int         `json:"maxResults,omitempty"`
	NextToken                 string      `json:"nextToken,omitempty"`
	IncludeExtendedDataFields bool        `json:"includeExtendedDataFields,omitempty"`
	NameFilter                *NameFilter `json:"nameFilter,omitempty"`
}

type ListCampaignsResponseV3 struct {
	TotalResults int          `json:"totalResults,omitempty"`
	Campaigns    []CampaignV3 `json:"campaigns,omitempty"`
	NextToken    string       `json:"nextToken,omitempty"`
}

func (r *ListCampaignsRequestV3) getNewResponse() IListCampaignsResponse {
	return new(ListCampaignsResponseV3)
}

func (r *ListCampaignsRequestV3) getVersion() int {
	return 3
}

func (r *ListCampaignsResponseV3) getVersion() int {
	return 3
}

//endregion ListV3
