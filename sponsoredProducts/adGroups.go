package sponsoredProducts

import "time"

type IListAdGroupsRequest interface {
	getNewResponse() IListAdGroupsResponse
	getVersion() int
}
type IListAdGroupsResponse interface {
	getVersion() int
}

type ICreateAdGroupsRequest interface {
	getNewResponse() ICreateAdGroupsResponse
	getVersion() int
}
type ICreateAdGroupsResponse interface {
	getVersion() int
}

type IDeleteAdGroupsRequest interface {
	getNewResponse() IDeleteAdGroupsResponse
	getVersion() int
}
type IDeleteAdGroupsResponse interface {
	getVersion() int
}

type IUpdateAdGroupsRequest interface {
	getNewResponse() IUpdateAdGroupsResponse
	getVersion() int
}
type IUpdateAdGroupsResponse interface {
	getVersion() int
}

// region CreateV3
type CreateAdGroupsRequestV3Data struct {
	CampaignID string  `json:"campaignId,omitempty"`
	Name       string  `json:"name,omitempty"`
	State      string  `json:"state,omitempty"`
	DefaultBid float64 `json:"defaultBid,omitempty"`
}
type CreateAdGroupsRequestV3 struct {
	AdGroups []CreateAdGroupsRequestV3Data `json:"adGroups,omitempty"`
}
type CreateAdGroupsResponseV3Success struct {
	AdGroup struct {
		CampaignID   string  `json:"campaignId,omitempty"`
		Name         string  `json:"name,omitempty"`
		State        string  `json:"state,omitempty"`
		AdGroupID    string  `json:"adGroupId,omitempty"`
		DefaultBid   float64 `json:"defaultBid,omitempty"`
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
	} `json:"adGroup,omitempty"`
	Index     int    `json:"index,omitempty"`
	AdGroupID string `json:"adGroupId,omitempty"`
}
type CreateAdGroupsResponseV3 struct {
	AdGroups struct {
		Success []CreateAdGroupsResponseV3Success `json:"success,omitempty"`
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
					OtherError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"otherError,omitempty"`
					ParentEntityError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"parentEntityError,omitempty"`
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
					ApplicableMarketplacesError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"applicableMarketplacesError,omitempty"`
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
	} `json:"adGroups,omitempty"`
}

func (r *CreateAdGroupsRequestV3) getNewResponse() ICreateAdGroupsResponse {
	return new(CreateAdGroupsResponseV3)
}

func (r *CreateAdGroupsRequestV3) getVersion() int {
	return 3
}

func (r *CreateAdGroupsResponseV3) getVersion() int {
	return 3
}

//endregion CreateV3

// region UpdateV3
type UpdateAdGroupsRequestV3Data struct {
	Name       string  `json:"name,omitempty"`
	State      string  `json:"state,omitempty"`
	AdGroupID  string  `json:"adGroupId,omitempty"`
	DefaultBid float64 `json:"defaultBid,omitempty"`
}
type UpdateAdGroupsRequestV3 struct {
	AdGroups []UpdateAdGroupsRequestV3Data `json:"adGroups,omitempty"`
}
type UpdateAdGroupsResponseV3Success struct {
	AdGroup struct {
		CampaignID   string  `json:"campaignId,omitempty"`
		Name         string  `json:"name,omitempty"`
		State        string  `json:"state,omitempty"`
		AdGroupID    string  `json:"adGroupId,omitempty"`
		DefaultBid   float64 `json:"defaultBid,omitempty"`
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
	} `json:"adGroup,omitempty"`
	Index     int    `json:"index,omitempty"`
	AdGroupID string `json:"adGroupId,omitempty"`
}
type UpdateAdGroupsResponseV3 struct {
	AdGroups struct {
		Success []UpdateAdGroupsResponseV3Success `json:"success,omitempty"`
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
					OtherError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"otherError,omitempty"`
					ParentEntityError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"parentEntityError,omitempty"`
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
					ApplicableMarketplacesError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"applicableMarketplacesError,omitempty"`
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
	} `json:"adGroups,omitempty"`
}

func (r *UpdateAdGroupsRequestV3) getNewResponse() IUpdateAdGroupsResponse {
	return new(UpdateAdGroupsResponseV3)
}

func (r *UpdateAdGroupsRequestV3) getVersion() int {
	return 3
}

func (r *UpdateAdGroupsResponseV3) getVersion() int {
	return 3
}

//endregion UpdateV3

// region DeleteV3
type DeleteAdGroupsRequestV3 struct {
	AdGroupIDFilter *IDFilter `json:"adGroupIdFilter,omitempty"`
}
type DeleteAdGroupsResponseV3Success struct {
	AdGroup struct {
		CampaignID   string  `json:"campaignId,omitempty"`
		Name         string  `json:"name,omitempty"`
		State        string  `json:"state,omitempty"`
		AdGroupID    string  `json:"adGroupId,omitempty"`
		DefaultBid   float64 `json:"defaultBid,omitempty"`
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
	} `json:"adGroup,omitempty"`
	Index     int    `json:"index,omitempty"`
	AdGroupID string `json:"adGroupId,omitempty"`
}
type DeleteAdGroupsResponseV3 struct {
	AdGroups struct {
		Success []DeleteAdGroupsResponseV3Success `json:"success,omitempty"`
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
					OtherError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"otherError,omitempty"`
					ParentEntityError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"parentEntityError,omitempty"`
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
					ApplicableMarketplacesError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"applicableMarketplacesError,omitempty"`
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
	} `json:"adGroups,omitempty"`
}

func (r *DeleteAdGroupsRequestV3) getNewResponse() IDeleteAdGroupsResponse {
	return new(DeleteAdGroupsResponseV3)
}

func (r *DeleteAdGroupsRequestV3) getVersion() int {
	return 3
}

func (r *DeleteAdGroupsResponseV3) getVersion() int {
	return 3
}

//endregion DeleteV3

// region ListV3
type ListAdGroupsRequestV3 struct {
	CampaignIDFilter            *IDFilter   `json:"campaignIdFilter,omitempty"`
	StateFilter                 *IDFilter   `json:"stateFilter,omitempty"`
	MaxResults                  int         `json:"maxResults,omitempty"`
	NextToken                   string      `json:"nextToken,omitempty"`
	AdGroupIDFilter             *IDFilter   `json:"adGroupIdFilter,omitempty"`
	IncludeExtendedDataFields   bool        `json:"includeExtendedDataFields,omitempty"`
	NameFilter                  *NameFilter `json:"nameFilter,omitempty"`
	CampaignTargetingTypeFilter string      `json:"campaignTargetingTypeFilter,omitempty"`
}
type ListAdGroupsResponseV3Data struct {
	CampaignID   string  `json:"campaignId,omitempty"`
	Name         string  `json:"name,omitempty"`
	State        string  `json:"state,omitempty"`
	AdGroupID    string  `json:"adGroupId,omitempty"`
	DefaultBid   float64 `json:"defaultBid,omitempty"`
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
type ListAdGroupsResponseV3 struct {
	TotalResults int                          `json:"totalResults,omitempty"`
	AdGroups     []ListAdGroupsResponseV3Data `json:"adGroups,omitempty"`
	NextToken    string                       `json:"nextToken,omitempty"`
}

func (r *ListAdGroupsRequestV3) getNewResponse() IListAdGroupsResponse {
	return new(ListAdGroupsResponseV3)
}

func (r *ListAdGroupsRequestV3) getVersion() int {
	return 3
}

func (r *ListAdGroupsResponseV3) getVersion() int {
	return 3
}

//endregion ListV3
