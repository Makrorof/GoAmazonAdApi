package sponsoredProducts

import "time"

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
	CampaignID string `json:"campaignId,omitempty"`
	Index      int    `json:"index,omitempty"`
	Campaign   struct {
		SiteRestrictions  []string `json:"siteRestrictions,omitempty"`
		OffAmazonSettings struct {
			OffAmazonBudgetControlStrategy string `json:"offAmazonBudgetControlStrategy,omitempty"`
		} `json:"offAmazonSettings,omitempty"`
		EndDate        string `json:"endDate,omitempty"`
		CampaignID     string `json:"campaignId,omitempty"`
		DynamicBidding struct {
			ShopperCohortBidding []struct {
				ShopperCohortType string `json:"shopperCohortType,omitempty"`
				Percentage        int    `json:"percentage,omitempty"`
				AudienceSegments  []struct {
				} `json:"audienceSegments,omitempty"`
			} `json:"shopperCohortBidding,omitempty"`
			PlacementBidding []struct {
				Percentage int    `json:"percentage,omitempty"`
				Placement  string `json:"placement,omitempty"`
			} `json:"placementBidding,omitempty"`
			Strategy string `json:"strategy,omitempty"`
		} `json:"dynamicBidding,omitempty"`
		Tags          map[string]string `json:"tags,omitempty"`
		PortfolioID   string            `json:"portfolioId,omitempty"`
		Name          string            `json:"name,omitempty"`
		TargetingType string            `json:"targetingType,omitempty"`
		State         string            `json:"state,omitempty"`
		StartDate     string            `json:"startDate,omitempty"`
		Budget        struct {
			BudgetType      string  `json:"budgetType,omitempty"`
			Budget          float64 `json:"budget,omitempty"`
			EffectiveBudget float64 `json:"effectiveBudget,omitempty"`
		} `json:"budget,omitempty"`
		ExtendedData struct {
			LastUpdateDateTime   time.Time `json:"lastUpdateDateTime,omitempty"`
			ServingStatus        string    `json:"servingStatus,omitempty"`
			ServingStatusDetails []struct {
				Name    string `json:"name,omitempty"`
				HelpURL string `json:"helpUrl,omitempty"`
				Message string `json:"message,omitempty"`
			} `json:"servingStatusDetails,omitempty"`
			CreationDateTime time.Time `json:"creationDateTime,omitempty"`
		} `json:"extendedData,omitempty"`
	} `json:"campaign,omitempty"`
}
type CreateCampaignsRequestV3 struct {
	Campaigns []CreateCampaignsV3CampaignData `json:"campaigns,omitempty"`
}
type CreateCampaignsResponseV3 struct {
	Campaigns struct {
		Success []CreateCampaignsV3Success `json:"success,omitempty"`
		Error   []struct {
			Index  int `json:"index,omitempty"`
			Errors []struct {
				ErrorType  string `json:"errorType,omitempty"`
				ErrorValue struct {
					EntityStateError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						EntityType  string `json:"entityType,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"entityStateError,omitempty"`
					MissingValueError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"missingValueError,omitempty"`
					DateError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"dateError,omitempty"`
					BiddingError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						UpperLimit string `json:"upperLimit,omitempty"`
						LowerLimit string `json:"lowerLimit,omitempty"`
						Message    string `json:"message,omitempty"`
					} `json:"biddingError,omitempty"`
					DuplicateValueError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"duplicateValueError,omitempty"`
					RangeError *struct {
						Reason      string   `json:"reason,omitempty"`
						Marketplace string   `json:"marketplace,omitempty"`
						Allowed     []string `json:"allowed,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						UpperLimit string `json:"upperLimit,omitempty"`
						LowerLimit string `json:"lowerLimit,omitempty"`
						Message    string `json:"message,omitempty"`
					} `json:"rangeError,omitempty"`
					ParentEntityError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"parentEntityError,omitempty"`
					OtherError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"otherError,omitempty"`
					ThrottledError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"throttledError,omitempty"`
					EntityNotFoundError *struct {
						Reason     string `json:"reason,omitempty"`
						EntityType string `json:"entityType,omitempty"`
						Cause      struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						EntityID string `json:"entityId,omitempty"`
						Message  string `json:"message,omitempty"`
					} `json:"entityNotFoundError,omitempty"`
					MalformedValueError *struct {
						Reason      string `json:"reason,omitempty"`
						Fragment    string `json:"fragment,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"malformedValueError,omitempty"`
					BudgetError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						UpperLimit string `json:"upperLimit,omitempty"`
						LowerLimit string `json:"lowerLimit,omitempty"`
						Message    string `json:"message,omitempty"`
					} `json:"budgetError,omitempty"`
					CurrencyError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"currencyError,omitempty"`
					BillingError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"billingError,omitempty"`
					EntityQuotaError *struct {
						Reason     string `json:"reason,omitempty"`
						QuotaScope string `json:"quotaScope,omitempty"`
						EntityType string `json:"entityType,omitempty"`
						Quota      string `json:"quota,omitempty"`
						Cause      struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"entityQuotaError,omitempty"`
					InternalServerError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"internalServerError,omitempty"`
				} `json:"errorValue,omitempty"`
			} `json:"errors,omitempty"`
		} `json:"error,omitempty"`
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
	CampaignID string `json:"campaignId,omitempty"`
	Index      int    `json:"index,omitempty"`
	Campaign   struct {
		SiteRestrictions  []string `json:"siteRestrictions,omitempty"`
		OffAmazonSettings struct {
			OffAmazonBudgetControlStrategy string `json:"offAmazonBudgetControlStrategy,omitempty"`
		} `json:"offAmazonSettings,omitempty"`
		EndDate        string `json:"endDate,omitempty"`
		CampaignID     string `json:"campaignId,omitempty"`
		DynamicBidding struct {
			ShopperCohortBidding []struct {
				ShopperCohortType string `json:"shopperCohortType,omitempty"`
				Percentage        int    `json:"percentage,omitempty"`
				AudienceSegments  []struct {
				} `json:"audienceSegments,omitempty"`
			} `json:"shopperCohortBidding,omitempty"`
			PlacementBidding []struct {
				Percentage int    `json:"percentage,omitempty"`
				Placement  string `json:"placement,omitempty"`
			} `json:"placementBidding,omitempty"`
			Strategy string `json:"strategy,omitempty"`
		} `json:"dynamicBidding,omitempty"`
		Tags          map[string]string `json:"tags,omitempty"`
		PortfolioID   string            `json:"portfolioId,omitempty"`
		Name          string            `json:"name,omitempty"`
		TargetingType string            `json:"targetingType,omitempty"`
		State         string            `json:"state,omitempty"`
		StartDate     string            `json:"startDate,omitempty"`
		Budget        struct {
			BudgetType      string  `json:"budgetType,omitempty"`
			Budget          float64 `json:"budget,omitempty"`
			EffectiveBudget float64 `json:"effectiveBudget,omitempty"`
		} `json:"budget,omitempty"`
		ExtendedData struct {
			LastUpdateDateTime   time.Time `json:"lastUpdateDateTime,omitempty"`
			ServingStatus        string    `json:"servingStatus,omitempty"`
			ServingStatusDetails []struct {
				Name    string `json:"name,omitempty"`
				HelpURL string `json:"helpUrl,omitempty"`
				Message string `json:"message,omitempty"`
			} `json:"servingStatusDetails,omitempty"`
			CreationDateTime time.Time `json:"creationDateTime,omitempty"`
		} `json:"extendedData,omitempty"`
	} `json:"campaign,omitempty"`
}
type UpdateCampaignsResponseV3 struct {
	Campaigns struct {
		Success []UpdateCampaignsResponseV3Success `json:"success,omitempty"`
		Error   []struct {
			Index  int `json:"index,omitempty"`
			Errors []struct {
				ErrorType  string `json:"errorType,omitempty"`
				ErrorValue struct {
					EntityStateError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						EntityType  string `json:"entityType,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"entityStateError,omitempty"`
					MissingValueError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"missingValueError,omitempty"`
					DateError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"dateError,omitempty"`
					BiddingError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						UpperLimit string `json:"upperLimit,omitempty"`
						LowerLimit string `json:"lowerLimit,omitempty"`
						Message    string `json:"message,omitempty"`
					} `json:"biddingError,omitempty"`
					DuplicateValueError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"duplicateValueError,omitempty"`
					RangeError *struct {
						Reason      string   `json:"reason,omitempty"`
						Marketplace string   `json:"marketplace,omitempty"`
						Allowed     []string `json:"allowed,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						UpperLimit string `json:"upperLimit,omitempty"`
						LowerLimit string `json:"lowerLimit,omitempty"`
						Message    string `json:"message,omitempty"`
					} `json:"rangeError,omitempty"`
					ParentEntityError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"parentEntityError,omitempty"`
					OtherError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"otherError,omitempty"`
					ThrottledError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"throttledError,omitempty"`
					EntityNotFoundError *struct {
						Reason     string `json:"reason,omitempty"`
						EntityType string `json:"entityType,omitempty"`
						Cause      struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						EntityID string `json:"entityId,omitempty"`
						Message  string `json:"message,omitempty"`
					} `json:"entityNotFoundError,omitempty"`
					MalformedValueError *struct {
						Reason      string `json:"reason,omitempty"`
						Fragment    string `json:"fragment,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"malformedValueError,omitempty"`
					BudgetError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						UpperLimit string `json:"upperLimit,omitempty"`
						LowerLimit string `json:"lowerLimit,omitempty"`
						Message    string `json:"message,omitempty"`
					} `json:"budgetError,omitempty"`
					CurrencyError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"currencyError,omitempty"`
					BillingError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"billingError,omitempty"`
					EntityQuotaError *struct {
						Reason     string `json:"reason,omitempty"`
						QuotaScope string `json:"quotaScope,omitempty"`
						EntityType string `json:"entityType,omitempty"`
						Quota      string `json:"quota,omitempty"`
						Cause      struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"entityQuotaError,omitempty"`
					InternalServerError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"internalServerError,omitempty"`
				} `json:"errorValue,omitempty"`
			} `json:"errors,omitempty"`
		} `json:"error,omitempty"`
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
	CampaignID string `json:"campaignId,omitempty"`
	Index      int    `json:"index,omitempty"`
	Campaign   struct {
		SiteRestrictions  []string `json:"siteRestrictions,omitempty"`
		OffAmazonSettings struct {
			OffAmazonBudgetControlStrategy string `json:"offAmazonBudgetControlStrategy,omitempty"`
		} `json:"offAmazonSettings,omitempty"`
		EndDate        string `json:"endDate,omitempty"`
		CampaignID     string `json:"campaignId,omitempty"`
		DynamicBidding struct {
			ShopperCohortBidding []struct {
				ShopperCohortType string `json:"shopperCohortType,omitempty"`
				Percentage        int    `json:"percentage,omitempty"`
				AudienceSegments  []struct {
				} `json:"audienceSegments,omitempty"`
			} `json:"shopperCohortBidding,omitempty"`
			PlacementBidding []struct {
				Percentage int    `json:"percentage,omitempty"`
				Placement  string `json:"placement,omitempty"`
			} `json:"placementBidding,omitempty"`
			Strategy string `json:"strategy,omitempty"`
		} `json:"dynamicBidding,omitempty"`
		Tags          map[string]string `json:"tags,omitempty"`
		PortfolioID   string            `json:"portfolioId,omitempty"`
		Name          string            `json:"name,omitempty"`
		TargetingType string            `json:"targetingType,omitempty"`
		State         string            `json:"state,omitempty"`
		StartDate     string            `json:"startDate,omitempty"`
		Budget        struct {
			BudgetType      string  `json:"budgetType,omitempty"`
			Budget          float64 `json:"budget,omitempty"`
			EffectiveBudget float64 `json:"effectiveBudget,omitempty"`
		} `json:"budget,omitempty"`
		ExtendedData struct {
			LastUpdateDateTime   time.Time `json:"lastUpdateDateTime,omitempty"`
			ServingStatus        string    `json:"servingStatus,omitempty"`
			ServingStatusDetails []struct {
				Name    string `json:"name,omitempty"`
				HelpURL string `json:"helpUrl,omitempty"`
				Message string `json:"message,omitempty"`
			} `json:"servingStatusDetails,omitempty"`
			CreationDateTime time.Time `json:"creationDateTime,omitempty"`
		} `json:"extendedData,omitempty"`
	} `json:"campaign,omitempty"`
}
type DeleteCampaignsResponseV3 struct {
	Campaigns struct {
		Success []DeleteCampaignsResponseV3Success `json:"success,omitempty"`
		Error   []struct {
			Index  int `json:"index,omitempty"`
			Errors []struct {
				ErrorType  string `json:"errorType,omitempty"`
				ErrorValue struct {
					EntityStateError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						EntityType  string `json:"entityType,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"entityStateError,omitempty"`
					MissingValueError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"missingValueError,omitempty"`
					DateError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"dateError,omitempty"`
					BiddingError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						UpperLimit string `json:"upperLimit,omitempty"`
						LowerLimit string `json:"lowerLimit,omitempty"`
						Message    string `json:"message,omitempty"`
					} `json:"biddingError,omitempty"`
					DuplicateValueError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"duplicateValueError,omitempty"`
					RangeError *struct {
						Reason      string   `json:"reason,omitempty"`
						Marketplace string   `json:"marketplace,omitempty"`
						Allowed     []string `json:"allowed,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						UpperLimit string `json:"upperLimit,omitempty"`
						LowerLimit string `json:"lowerLimit,omitempty"`
						Message    string `json:"message,omitempty"`
					} `json:"rangeError,omitempty"`
					ParentEntityError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"parentEntityError,omitempty"`
					OtherError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"otherError,omitempty"`
					ThrottledError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"throttledError,omitempty"`
					EntityNotFoundError *struct {
						Reason     string `json:"reason,omitempty"`
						EntityType string `json:"entityType,omitempty"`
						Cause      struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						EntityID string `json:"entityId,omitempty"`
						Message  string `json:"message,omitempty"`
					} `json:"entityNotFoundError,omitempty"`
					MalformedValueError *struct {
						Reason      string `json:"reason,omitempty"`
						Fragment    string `json:"fragment,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"malformedValueError,omitempty"`
					BudgetError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						UpperLimit string `json:"upperLimit,omitempty"`
						LowerLimit string `json:"lowerLimit,omitempty"`
						Message    string `json:"message,omitempty"`
					} `json:"budgetError,omitempty"`
					CurrencyError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"currencyError,omitempty"`
					BillingError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"billingError,omitempty"`
					EntityQuotaError *struct {
						Reason     string `json:"reason,omitempty"`
						QuotaScope string `json:"quotaScope,omitempty"`
						EntityType string `json:"entityType,omitempty"`
						Quota      string `json:"quota,omitempty"`
						Cause      struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"entityQuotaError,omitempty"`
					InternalServerError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"internalServerError,omitempty"`
				} `json:"errorValue,omitempty"`
			} `json:"errors,omitempty"`
		} `json:"error,omitempty"`
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
type ListCampaignsResponseV3Campaign struct {
	SiteRestrictions  []string `json:"siteRestrictions,omitempty"`
	OffAmazonSettings struct {
		OffAmazonBudgetControlStrategy string `json:"offAmazonBudgetControlStrategy,omitempty"`
	} `json:"offAmazonSettings,omitempty"`
	EndDate        string `json:"endDate,omitempty"`
	CampaignID     string `json:"campaignId,omitempty"`
	DynamicBidding struct {
		ShopperCohortBidding []struct {
			ShopperCohortType string `json:"shopperCohortType,omitempty"`
			Percentage        int    `json:"percentage,omitempty"`
			AudienceSegments  []struct {
				AudienceID          string `json:"audienceId,omitempty"`
				AudienceSegmentType string `json:"audienceSegmentType,omitempty"`
			} `json:"audienceSegments,omitempty"`
		} `json:"shopperCohortBidding,omitempty"`
		PlacementBidding []struct {
			Percentage int    `json:"percentage,omitempty"`
			Placement  string `json:"placement,omitempty"`
		} `json:"placementBidding,omitempty"`
		Strategy string `json:"strategy,omitempty"`
	} `json:"dynamicBidding,omitempty"`
	Tags          map[string]string `json:"tags,omitempty"`
	PortfolioID   string            `json:"portfolioId,omitempty"`
	Name          string            `json:"name,omitempty"`
	TargetingType string            `json:"targetingType,omitempty"`
	State         string            `json:"state,omitempty"`
	StartDate     string            `json:"startDate,omitempty"`
	Budget        struct {
		BudgetType      string  `json:"budgetType,omitempty"`
		Budget          float64 `json:"budget,omitempty"`
		EffectiveBudget float64 `json:"effectiveBudget,omitempty"`
	} `json:"budget,omitempty"`
	ExtendedData struct {
		LastUpdateDateTime   time.Time `json:"lastUpdateDateTime,omitempty"`
		ServingStatus        string    `json:"servingStatus,omitempty"`
		ServingStatusDetails []struct {
			Name    string `json:"name,omitempty"`
			HelpURL string `json:"helpUrl,omitempty"`
			Message string `json:"message,omitempty"`
		} `json:"servingStatusDetails,omitempty"`
		CreationDateTime time.Time `json:"creationDateTime,omitempty"`
	} `json:"extendedData,omitempty"`
}
type ListCampaignsResponseV3 struct {
	TotalResults int                               `json:"totalResults,omitempty"`
	Campaigns    []ListCampaignsResponseV3Campaign `json:"campaigns,omitempty"`
	NextToken    string                            `json:"nextToken,omitempty"`
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
