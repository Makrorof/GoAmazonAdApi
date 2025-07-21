package sponsoredProducts

import "time"

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

type IDFilter struct {
	Include []string `json:"include,omitempty"`
}

type NameFilter struct {
	QueryTermMatchType string   `json:"queryTermMatchType,omitempty"`
	Include            []string `json:"include,omitempty"`
}

type TypeValue struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type TargetingExpression struct {
	Type TargetingExpressionType `json:"type,omitempty"`
}

type GlobalStoreSettingFilter struct {
	CatalogSourceCountryCode string `json:"catalogSourceCountryCode,omitempty"`
}

type ProductDetailsListFilter struct {
	GlobalStoreSetting *GlobalStoreSettingFilter `json:"globalStoreSetting,omitempty"`
	Asin               string                    `json:"asin,omitempty"`
}

type AdjustmentFilter struct {
	Predicate  string `json:"predicate,omitempty"`
	Percentage string `json:"percentage,omitempty"`
}

type BiddingFilter struct {
	Adjustments []AdjustmentFilter `json:"adjustments,omitempty"`
	Strategy    string             `json:"strategy,omitempty"`
}

type OffAmazonSettingFilter struct {
	OffAmazonBudgetControlStrategy string `json:"offAmazonBudgetControlStrategy,omitempty"`
}

type PlacementBiddingFilter struct {
	Percentage int    `json:"percentage,omitempty"`
	Placement  string `json:"placement,omitempty"`
}
type AudienceSegmentsFilter struct {
	AudienceID          string `json:"audienceId,omitempty"`
	AudienceSegmentType string `json:"audienceSegmentType,omitempty"`
}
type ShopperCohortBiddingFilter struct {
	ShopperCohortType string                   `json:"shopperCohortType,omitempty"`
	Percentage        int                      `json:"percentage,omitempty"`
	AudienceSegments  []AudienceSegmentsFilter `json:"audienceSegments,omitempty"`
}
type DynamicBiddingFilter struct {
	ShopperCohortBidding []ShopperCohortBiddingFilter `json:"shopperCohortBidding,omitempty"`
	PlacementBidding     []PlacementBiddingFilter     `json:"placementBidding,omitempty"`
	Strategy             string                       `json:"strategy,omitempty"`
}

type BudgetFilter struct {
	BudgetType string  `json:"budgetType,omitempty"`
	Budget     float64 `json:"budget,omitempty"`
}

type TargetsFilter struct {
	MatchType           string  `json:"matchType,omitempty"`
	Keyword             string  `json:"keyword,omitempty"`
	Bid                 float64 `json:"bid,omitempty"`
	UserSelectedKeyword bool    `json:"userSelectedKeyword,omitempty"`
}

type ServingStatusDetail struct {
	Name    string `json:"name,omitempty"`
	HelpURL string `json:"helpUrl,omitempty"`
	Message string `json:"message,omitempty"`
}
type ExtendedData struct {
	LastUpdateDateTime   time.Time             `json:"lastUpdateDateTime,omitempty"`
	ServingStatus        string                `json:"servingStatus,omitempty"`
	ServingStatusDetails []ServingStatusDetail `json:"servingStatusDetails,omitempty"`
	CreationDateTime     time.Time             `json:"creationDateTime,omitempty"`
}
