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
	CampaignID string  `json:"campaignId"`
	Name       string  `json:"name"`
	State      string  `json:"state"`
	DefaultBid float64 `json:"defaultBid"`
}
type CreateAdGroupsRequestV3 struct {
	AdGroups []CreateAdGroupsRequestV3Data `json:"adGroups"`
}
type CreateAdGroupsResponseV3Success struct {
	AdGroup struct {
		CampaignID   string  `json:"campaignId"`
		Name         string  `json:"name"`
		State        string  `json:"state"`
		AdGroupID    string  `json:"adGroupId"`
		DefaultBid   float64 `json:"defaultBid"`
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
	} `json:"adGroup"`
	Index     int    `json:"index"`
	AdGroupID string `json:"adGroupId"`
}
type CreateAdGroupsResponseV3 struct {
	AdGroups struct {
		Success []CreateAdGroupsResponseV3Success `json:"success"`
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
					OtherError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"otherError"`
					ParentEntityError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"parentEntityError"`
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
					ApplicableMarketplacesError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"applicableMarketplacesError"`
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
	} `json:"adGroups"`
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
	Name       string  `json:"name"`
	State      string  `json:"state"`
	AdGroupID  string  `json:"adGroupId"`
	DefaultBid float64 `json:"defaultBid"`
}
type UpdateAdGroupsRequestV3 struct {
	AdGroups []UpdateAdGroupsRequestV3Data `json:"adGroups"`
}
type UpdateAdGroupsResponseV3Success struct {
	AdGroup struct {
		CampaignID   string  `json:"campaignId"`
		Name         string  `json:"name"`
		State        string  `json:"state"`
		AdGroupID    string  `json:"adGroupId"`
		DefaultBid   float64 `json:"defaultBid"`
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
	} `json:"adGroup"`
	Index     int    `json:"index"`
	AdGroupID string `json:"adGroupId"`
}
type UpdateAdGroupsResponseV3 struct {
	AdGroups struct {
		Success []UpdateAdGroupsResponseV3Success `json:"success"`
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
					OtherError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"otherError"`
					ParentEntityError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"parentEntityError"`
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
					ApplicableMarketplacesError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"applicableMarketplacesError"`
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
	} `json:"adGroups"`
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
	AdGroupIDFilter struct {
		Include []string `json:"include"`
	} `json:"adGroupIdFilter"`
}
type DeleteAdGroupsResponseV3Success struct {
	AdGroup struct {
		CampaignID   string  `json:"campaignId"`
		Name         string  `json:"name"`
		State        string  `json:"state"`
		AdGroupID    string  `json:"adGroupId"`
		DefaultBid   float64 `json:"defaultBid"`
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
	} `json:"adGroup"`
	Index     int    `json:"index"`
	AdGroupID string `json:"adGroupId"`
}
type DeleteAdGroupsResponseV3 struct {
	AdGroups struct {
		Success []DeleteAdGroupsResponseV3Success `json:"success"`
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
					OtherError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"otherError"`
					ParentEntityError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"parentEntityError"`
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
					ApplicableMarketplacesError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"applicableMarketplacesError"`
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
	} `json:"adGroups"`
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
	CampaignIDFilter struct {
		Include []string `json:"include"`
	} `json:"campaignIdFilter"`
	StateFilter struct {
		Include []string `json:"include"`
	} `json:"stateFilter"`
	MaxResults      int    `json:"maxResults"`
	NextToken       string `json:"nextToken"`
	AdGroupIDFilter struct {
		Include []string `json:"include"`
	} `json:"adGroupIdFilter"`
	IncludeExtendedDataFields bool `json:"includeExtendedDataFields"`
	NameFilter                struct {
		QueryTermMatchType string   `json:"queryTermMatchType"`
		Include            []string `json:"include"`
	} `json:"nameFilter"`
	CampaignTargetingTypeFilter string `json:"campaignTargetingTypeFilter"`
}
type ListAdGroupsResponseV3Data struct {
	CampaignID   string  `json:"campaignId"`
	Name         string  `json:"name"`
	State        string  `json:"state"`
	AdGroupID    string  `json:"adGroupId"`
	DefaultBid   float64 `json:"defaultBid"`
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
type ListAdGroupsResponseV3 struct {
	TotalResults int                          `json:"totalResults"`
	AdGroups     []ListAdGroupsResponseV3Data `json:"adGroups"`
	NextToken    string                       `json:"nextToken"`
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
