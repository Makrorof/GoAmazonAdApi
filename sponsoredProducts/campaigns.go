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
	SiteRestrictions  []string `json:"siteRestrictions"`
	PortfolioID       string   `json:"portfolioId"`
	OffAmazonSettings struct {
		OffAmazonBudgetControlStrategy string `json:"offAmazonBudgetControlStrategy"`
	} `json:"offAmazonSettings"`
	EndDate        string `json:"endDate"`
	Name           string `json:"name"`
	TargetingType  string `json:"targetingType"`
	State          string `json:"state"`
	DynamicBidding struct {
		ShopperCohortBidding []struct {
			ShopperCohortType string `json:"shopperCohortType"`
			Percentage        int    `json:"percentage"`
			AudienceSegments  []struct {
				AudienceID          string `json:"audienceId"`
				AudienceSegmentType string `json:"audienceSegmentType"`
			} `json:"audienceSegments"`
		} `json:"shopperCohortBidding"`
		PlacementBidding []struct {
			Percentage int    `json:"percentage"`
			Placement  string `json:"placement"`
		} `json:"placementBidding"`
		Strategy string `json:"strategy"`
	} `json:"dynamicBidding"`
	StartDate string `json:"startDate"`
	Budget    struct {
		BudgetType string  `json:"budgetType"`
		Budget     float64 `json:"budget"`
	} `json:"budget"`
	Tags map[string]string `json:"tags"`
}
type CreateCampaignsV3Success struct {
	CampaignID string `json:"campaignId"`
	Index      int    `json:"index"`
	Campaign   struct {
		SiteRestrictions  []string `json:"siteRestrictions"`
		OffAmazonSettings struct {
			OffAmazonBudgetControlStrategy string `json:"offAmazonBudgetControlStrategy"`
		} `json:"offAmazonSettings"`
		EndDate        string `json:"endDate"`
		CampaignID     string `json:"campaignId"`
		DynamicBidding struct {
			ShopperCohortBidding []struct {
				ShopperCohortType string `json:"shopperCohortType"`
				Percentage        int    `json:"percentage"`
				AudienceSegments  []struct {
				} `json:"audienceSegments"`
			} `json:"shopperCohortBidding"`
			PlacementBidding []struct {
				Percentage int    `json:"percentage"`
				Placement  string `json:"placement"`
			} `json:"placementBidding"`
			Strategy string `json:"strategy"`
		} `json:"dynamicBidding"`
		Tags          map[string]string `json:"tags"`
		PortfolioID   string            `json:"portfolioId"`
		Name          string            `json:"name"`
		TargetingType string            `json:"targetingType"`
		State         string            `json:"state"`
		StartDate     string            `json:"startDate"`
		Budget        struct {
			BudgetType      string  `json:"budgetType"`
			Budget          float64 `json:"budget"`
			EffectiveBudget float64 `json:"effectiveBudget"`
		} `json:"budget"`
		ExtendedData struct {
			LastUpdateDateTime   time.Time `json:"lastUpdateDateTime"`
			ServingStatus        string    `json:"servingStatus"`
			ServingStatusDetails []struct {
				Name    string `json:"name"`
				HelpURL string `json:"helpUrl"`
				Message string `json:"message"`
			} `json:"servingStatusDetails"`
			CreationDateTime time.Time `json:"creationDateTime"`
		} `json:"extendedData"`
	} `json:"campaign"`
}
type CreateCampaignsRequestV3 struct {
	Campaigns []CreateCampaignsV3CampaignData `json:"campaigns"`
}

type CreateCampaignsResponseV3 struct {
	Campaigns struct {
		Success []CreateCampaignsV3Success `json:"success"`
		Error   []struct {
			Index  int `json:"index"`
			Errors []struct {
				ErrorType  string `json:"errorType"`
				ErrorValue struct {
					EntityStateError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						EntityType  string `json:"entityType"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"entityStateError"`
					MissingValueError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"missingValueError"`
					DateError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"dateError"`
					BiddingError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						UpperLimit string `json:"upperLimit"`
						LowerLimit string `json:"lowerLimit"`
						Message    string `json:"message"`
					} `json:"biddingError"`
					DuplicateValueError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"duplicateValueError"`
					RangeError *struct {
						Reason      string   `json:"reason"`
						Marketplace string   `json:"marketplace"`
						Allowed     []string `json:"allowed"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						UpperLimit string `json:"upperLimit"`
						LowerLimit string `json:"lowerLimit"`
						Message    string `json:"message"`
					} `json:"rangeError"`
					ParentEntityError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"parentEntityError"`
					OtherError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"otherError"`
					ThrottledError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"throttledError"`
					EntityNotFoundError *struct {
						Reason     string `json:"reason"`
						EntityType string `json:"entityType"`
						Cause      struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						EntityID string `json:"entityId"`
						Message  string `json:"message"`
					} `json:"entityNotFoundError"`
					MalformedValueError *struct {
						Reason      string `json:"reason"`
						Fragment    string `json:"fragment"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"malformedValueError"`
					BudgetError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						UpperLimit string `json:"upperLimit"`
						LowerLimit string `json:"lowerLimit"`
						Message    string `json:"message"`
					} `json:"budgetError"`
					CurrencyError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"currencyError"`
					BillingError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"billingError"`
					EntityQuotaError *struct {
						Reason     string `json:"reason"`
						QuotaScope string `json:"quotaScope"`
						EntityType string `json:"entityType"`
						Quota      string `json:"quota"`
						Cause      struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"entityQuotaError"`
					InternalServerError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"internalServerError"`
				} `json:"errorValue"`
			} `json:"errors"`
		} `json:"error"`
	} `json:"campaigns"`
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
	SiteRestrictions  []string `json:"siteRestrictions"`
	PortfolioID       string   `json:"portfolioId"`
	OffAmazonSettings struct {
		OffAmazonBudgetControlStrategy string `json:"offAmazonBudgetControlStrategy"`
	} `json:"offAmazonSettings"`
	EndDate        string `json:"endDate"`
	CampaignID     string `json:"campaignId"`
	Name           string `json:"name"`
	TargetingType  string `json:"targetingType"`
	State          string `json:"state"`
	DynamicBidding struct {
		ShopperCohortBidding []struct {
			ShopperCohortType string `json:"shopperCohortType"`
			Percentage        int    `json:"percentage"`
			AudienceSegments  []struct {
				AudienceID          string `json:"audienceId"`
				AudienceSegmentType string `json:"audienceSegmentType"`
			} `json:"audienceSegments"`
		} `json:"shopperCohortBidding"`
		PlacementBidding []struct {
			Percentage int    `json:"percentage"`
			Placement  string `json:"placement"`
		} `json:"placementBidding"`
		Strategy string `json:"strategy"`
	} `json:"dynamicBidding"`
	StartDate string `json:"startDate"`
	Budget    struct {
		BudgetType string  `json:"budgetType"`
		Budget     float64 `json:"budget"`
	} `json:"budget"`
	Tags map[string]string `json:"tags"`
}
type UpdateCampaignsRequestV3 struct {
	Campaigns []UpdateCampaignsRequestV3Data `json:"campaigns"`
}
type UpdateCampaignsResponseV3Success struct {
	CampaignID string `json:"campaignId"`
	Index      int    `json:"index"`
	Campaign   struct {
		SiteRestrictions  []string `json:"siteRestrictions"`
		OffAmazonSettings struct {
			OffAmazonBudgetControlStrategy string `json:"offAmazonBudgetControlStrategy"`
		} `json:"offAmazonSettings"`
		EndDate        string `json:"endDate"`
		CampaignID     string `json:"campaignId"`
		DynamicBidding struct {
			ShopperCohortBidding []struct {
				ShopperCohortType string `json:"shopperCohortType"`
				Percentage        int    `json:"percentage"`
				AudienceSegments  []struct {
				} `json:"audienceSegments"`
			} `json:"shopperCohortBidding"`
			PlacementBidding []struct {
				Percentage int    `json:"percentage"`
				Placement  string `json:"placement"`
			} `json:"placementBidding"`
			Strategy string `json:"strategy"`
		} `json:"dynamicBidding"`
		Tags          map[string]string `json:"tags"`
		PortfolioID   string            `json:"portfolioId"`
		Name          string            `json:"name"`
		TargetingType string            `json:"targetingType"`
		State         string            `json:"state"`
		StartDate     string            `json:"startDate"`
		Budget        struct {
			BudgetType      string  `json:"budgetType"`
			Budget          float64 `json:"budget"`
			EffectiveBudget float64 `json:"effectiveBudget"`
		} `json:"budget"`
		ExtendedData struct {
			LastUpdateDateTime   time.Time `json:"lastUpdateDateTime"`
			ServingStatus        string    `json:"servingStatus"`
			ServingStatusDetails []struct {
				Name    string `json:"name"`
				HelpURL string `json:"helpUrl"`
				Message string `json:"message"`
			} `json:"servingStatusDetails"`
			CreationDateTime time.Time `json:"creationDateTime"`
		} `json:"extendedData"`
	} `json:"campaign"`
}
type UpdateCampaignsResponseV3 struct {
	Campaigns struct {
		Success []UpdateCampaignsResponseV3Success `json:"success"`
		Error   []struct {
			Index  int `json:"index"`
			Errors []struct {
				ErrorType  string `json:"errorType"`
				ErrorValue struct {
					EntityStateError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						EntityType  string `json:"entityType"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"entityStateError"`
					MissingValueError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"missingValueError"`
					DateError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"dateError"`
					BiddingError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						UpperLimit string `json:"upperLimit"`
						LowerLimit string `json:"lowerLimit"`
						Message    string `json:"message"`
					} `json:"biddingError"`
					DuplicateValueError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"duplicateValueError"`
					RangeError *struct {
						Reason      string   `json:"reason"`
						Marketplace string   `json:"marketplace"`
						Allowed     []string `json:"allowed"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						UpperLimit string `json:"upperLimit"`
						LowerLimit string `json:"lowerLimit"`
						Message    string `json:"message"`
					} `json:"rangeError"`
					ParentEntityError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"parentEntityError"`
					OtherError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"otherError"`
					ThrottledError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"throttledError"`
					EntityNotFoundError *struct {
						Reason     string `json:"reason"`
						EntityType string `json:"entityType"`
						Cause      struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						EntityID string `json:"entityId"`
						Message  string `json:"message"`
					} `json:"entityNotFoundError"`
					MalformedValueError *struct {
						Reason      string `json:"reason"`
						Fragment    string `json:"fragment"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"malformedValueError"`
					BudgetError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						UpperLimit string `json:"upperLimit"`
						LowerLimit string `json:"lowerLimit"`
						Message    string `json:"message"`
					} `json:"budgetError"`
					CurrencyError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"currencyError"`
					BillingError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"billingError"`
					EntityQuotaError *struct {
						Reason     string `json:"reason"`
						QuotaScope string `json:"quotaScope"`
						EntityType string `json:"entityType"`
						Quota      string `json:"quota"`
						Cause      struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"entityQuotaError"`
					InternalServerError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"internalServerError"`
				} `json:"errorValue"`
			} `json:"errors"`
		} `json:"error"`
	} `json:"campaigns"`
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
	CampaignIDFilter struct {
		Include []string `json:"include"`
	} `json:"campaignIdFilter"`
}
type DeleteCampaignsResponseV3Success struct {
	CampaignID string `json:"campaignId"`
	Index      int    `json:"index"`
	Campaign   struct {
		SiteRestrictions  []string `json:"siteRestrictions"`
		OffAmazonSettings struct {
			OffAmazonBudgetControlStrategy string `json:"offAmazonBudgetControlStrategy"`
		} `json:"offAmazonSettings"`
		EndDate        string `json:"endDate"`
		CampaignID     string `json:"campaignId"`
		DynamicBidding struct {
			ShopperCohortBidding []struct {
				ShopperCohortType string `json:"shopperCohortType"`
				Percentage        int    `json:"percentage"`
				AudienceSegments  []struct {
				} `json:"audienceSegments"`
			} `json:"shopperCohortBidding"`
			PlacementBidding []struct {
				Percentage int    `json:"percentage"`
				Placement  string `json:"placement"`
			} `json:"placementBidding"`
			Strategy string `json:"strategy"`
		} `json:"dynamicBidding"`
		Tags          map[string]string `json:"tags"`
		PortfolioID   string            `json:"portfolioId"`
		Name          string            `json:"name"`
		TargetingType string            `json:"targetingType"`
		State         string            `json:"state"`
		StartDate     string            `json:"startDate"`
		Budget        struct {
			BudgetType      string  `json:"budgetType"`
			Budget          float64 `json:"budget"`
			EffectiveBudget float64 `json:"effectiveBudget"`
		} `json:"budget"`
		ExtendedData struct {
			LastUpdateDateTime   time.Time `json:"lastUpdateDateTime"`
			ServingStatus        string    `json:"servingStatus"`
			ServingStatusDetails []struct {
				Name    string `json:"name"`
				HelpURL string `json:"helpUrl"`
				Message string `json:"message"`
			} `json:"servingStatusDetails"`
			CreationDateTime time.Time `json:"creationDateTime"`
		} `json:"extendedData"`
	} `json:"campaign"`
}
type DeleteCampaignsResponseV3 struct {
	Campaigns struct {
		Success []DeleteCampaignsResponseV3Success `json:"success"`
		Error   []struct {
			Index  int `json:"index"`
			Errors []struct {
				ErrorType  string `json:"errorType"`
				ErrorValue struct {
					EntityStateError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						EntityType  string `json:"entityType"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"entityStateError"`
					MissingValueError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"missingValueError"`
					DateError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"dateError"`
					BiddingError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						UpperLimit string `json:"upperLimit"`
						LowerLimit string `json:"lowerLimit"`
						Message    string `json:"message"`
					} `json:"biddingError"`
					DuplicateValueError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"duplicateValueError"`
					RangeError *struct {
						Reason      string   `json:"reason"`
						Marketplace string   `json:"marketplace"`
						Allowed     []string `json:"allowed"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						UpperLimit string `json:"upperLimit"`
						LowerLimit string `json:"lowerLimit"`
						Message    string `json:"message"`
					} `json:"rangeError"`
					ParentEntityError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"parentEntityError"`
					OtherError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"otherError"`
					ThrottledError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"throttledError"`
					EntityNotFoundError *struct {
						Reason     string `json:"reason"`
						EntityType string `json:"entityType"`
						Cause      struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						EntityID string `json:"entityId"`
						Message  string `json:"message"`
					} `json:"entityNotFoundError"`
					MalformedValueError *struct {
						Reason      string `json:"reason"`
						Fragment    string `json:"fragment"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"malformedValueError"`
					BudgetError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						UpperLimit string `json:"upperLimit"`
						LowerLimit string `json:"lowerLimit"`
						Message    string `json:"message"`
					} `json:"budgetError"`
					CurrencyError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"currencyError"`
					BillingError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"billingError"`
					EntityQuotaError *struct {
						Reason     string `json:"reason"`
						QuotaScope string `json:"quotaScope"`
						EntityType string `json:"entityType"`
						Quota      string `json:"quota"`
						Cause      struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"entityQuotaError"`
					InternalServerError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"internalServerError"`
				} `json:"errorValue"`
			} `json:"errors"`
		} `json:"error"`
	} `json:"campaigns"`
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
	CampaignIDFilter struct {
		Include []string `json:"include"`
	} `json:"campaignIdFilter"`
	PortfolioIDFilter struct {
		Include []string `json:"include"`
	} `json:"portfolioIdFilter"`
	StateFilter struct {
		Include []string `json:"include"`
	} `json:"stateFilter"`
	MaxResults                int    `json:"maxResults"`
	NextToken                 string `json:"nextToken"`
	IncludeExtendedDataFields bool   `json:"includeExtendedDataFields"`
	NameFilter                struct {
		QueryTermMatchType string   `json:"queryTermMatchType"`
		Include            []string `json:"include"`
	} `json:"nameFilter"`
}
type ListCampaignsResponseV3Campaign struct {
	SiteRestrictions  []string `json:"siteRestrictions"`
	OffAmazonSettings struct {
		OffAmazonBudgetControlStrategy string `json:"offAmazonBudgetControlStrategy"`
	} `json:"offAmazonSettings"`
	EndDate        string `json:"endDate"`
	CampaignID     string `json:"campaignId"`
	DynamicBidding struct {
		ShopperCohortBidding []struct {
			ShopperCohortType string `json:"shopperCohortType"`
			Percentage        int    `json:"percentage"`
			AudienceSegments  []struct {
				AudienceID          string `json:"audienceId"`
				AudienceSegmentType string `json:"audienceSegmentType"`
			} `json:"audienceSegments"`
		} `json:"shopperCohortBidding"`
		PlacementBidding []struct {
			Percentage int    `json:"percentage"`
			Placement  string `json:"placement"`
		} `json:"placementBidding"`
		Strategy string `json:"strategy"`
	} `json:"dynamicBidding"`
	Tags          map[string]string `json:"tags"`
	PortfolioID   string            `json:"portfolioId"`
	Name          string            `json:"name"`
	TargetingType string            `json:"targetingType"`
	State         string            `json:"state"`
	StartDate     string            `json:"startDate"`
	Budget        struct {
		BudgetType      string  `json:"budgetType"`
		Budget          float64 `json:"budget"`
		EffectiveBudget float64 `json:"effectiveBudget"`
	} `json:"budget"`
	ExtendedData struct {
		LastUpdateDateTime   time.Time `json:"lastUpdateDateTime"`
		ServingStatus        string    `json:"servingStatus"`
		ServingStatusDetails []struct {
			Name    string `json:"name"`
			HelpURL string `json:"helpUrl"`
			Message string `json:"message"`
		} `json:"servingStatusDetails"`
		CreationDateTime time.Time `json:"creationDateTime"`
	} `json:"extendedData"`
}
type ListCampaignsResponseV3 struct {
	TotalResults int                               `json:"totalResults"`
	Campaigns    []ListCampaignsResponseV3Campaign `json:"campaigns"`
	NextToken    string                            `json:"nextToken"`
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
