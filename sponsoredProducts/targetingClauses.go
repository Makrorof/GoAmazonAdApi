package sponsoredProducts

import "time"

type IListTargetingClausesRequest interface {
	getNewResponse() IListTargetingClausesResponse
	getVersion() int
}
type IListTargetingClausesResponse interface {
	getVersion() int
}

type ICreateTargetingClausesRequest interface {
	getNewResponse() ICreateTargetingClausesResponse
	getVersion() int
}
type ICreateTargetingClausesResponse interface {
	getVersion() int
}

type IDeleteTargetingClausesRequest interface {
	getNewResponse() IDeleteTargetingClausesResponse
	getVersion() int
}
type IDeleteTargetingClausesResponse interface {
	getVersion() int
}

type IUpdateTargetingClausesRequest interface {
	getNewResponse() IUpdateTargetingClausesResponse
	getVersion() int
}
type IUpdateTargetingClausesResponse interface {
	getVersion() int
}

// region Listv3
type ListTargetingClausesRequestV3 struct {
	CampaignIDFilter          *IDFilter   `json:"campaignIdFilter,omitempty"`
	StateFilter               *IDFilter   `json:"stateFilter,omitempty"`
	ExpressionTypeFilter      *IDFilter   `json:"expressionTypeFilter,omitempty"`
	MaxResults                int         `json:"maxResults,omitempty"`
	NextToken                 string      `json:"nextToken,omitempty"`
	TargetIDFilter            *IDFilter   `json:"targetIdFilter,omitempty"`
	AsinFilter                *NameFilter `json:"asinFilter,omitempty"`
	AdGroupIDFilter           *IDFilter   `json:"adGroupIdFilter,omitempty"`
	IncludeExtendedDataFields bool        `json:"includeExtendedDataFields,omitempty"`
}
type ListTargetingClausesData struct {
	Expression []struct {
		Type  string `json:"type,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"expression,omitempty"`
	TargetID           string `json:"targetId,omitempty"`
	ResolvedExpression []struct {
		Type  string `json:"type,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"resolvedExpression,omitempty"`
	CampaignID     string  `json:"campaignId,omitempty"`
	ExpressionType string  `json:"expressionType,omitempty"`
	State          string  `json:"state,omitempty"`
	Bid            float64 `json:"bid,omitempty"`
	AdGroupID      string  `json:"adGroupId,omitempty"`
	ExtendedData   struct {
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
type ListTargetingClausesResponseV3 struct {
	TotalResults     int                        `json:"totalResults,omitempty"`
	NextToken        string                     `json:"nextToken,omitempty"`
	TargetingClauses []ListTargetingClausesData `json:"targetingClauses,omitempty"`
}

func (r *ListTargetingClausesRequestV3) getNewResponse() IListTargetingClausesResponse {
	return new(ListTargetingClausesResponseV3)
}
func (r *ListTargetingClausesRequestV3) getVersion() int {
	return 3
}

func (r *ListTargetingClausesResponseV3) getVersion() int {
	return 3
}

//endregion Listv3

// region CreateV3
type CreateTargetingClausesData struct {
	Expression     []TypeValueFilter `json:"expression,omitempty"`
	CampaignID     string            `json:"campaignId,omitempty"`
	ExpressionType string            `json:"expressionType,omitempty"`
	State          string            `json:"state,omitempty"`
	Bid            float64           `json:"bid,omitempty"`
	AdGroupID      string            `json:"adGroupId,omitempty"`
}
type CreateTargetingClausesRequestV3 struct {
	TargetingClauses []CreateTargetingClausesData `json:"targetingClauses,omitempty"`
}
type CreateTargetingClausesSuccess struct {
	TargetingClause struct {
		Expression []struct {
			Type  string `json:"type,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"expression,omitempty"`
		TargetID           string `json:"targetId,omitempty"`
		ResolvedExpression []struct {
			Type  string `json:"type,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"resolvedExpression,omitempty"`
		CampaignID     string  `json:"campaignId,omitempty"`
		ExpressionType string  `json:"expressionType,omitempty"`
		State          string  `json:"state,omitempty"`
		Bid            float64 `json:"bid,omitempty"`
		AdGroupID      string  `json:"adGroupId,omitempty"`
		ExtendedData   struct {
			LastUpdateDateTime   time.Time `json:"lastUpdateDateTime,omitempty"`
			ServingStatus        string    `json:"servingStatus,omitempty"`
			ServingStatusDetails []struct {
				Name    string `json:"name,omitempty"`
				HelpURL string `json:"helpUrl,omitempty"`
				Message string `json:"message,omitempty"`
			} `json:"servingStatusDetails,omitempty"`
			CreationDateTime time.Time `json:"creationDateTime,omitempty"`
		} `json:"extendedData,omitempty"`
	} `json:"targetingClause,omitempty"`
	TargetID string `json:"targetId,omitempty"`
	Index    int    `json:"index,omitempty"`
}
type CreateTargetingClausesResponseV3 struct {
	TargetingClauses struct {
		Success []CreateTargetingClausesSuccess `json:"success,omitempty"`
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
					ExpressionTypeError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"expressionTypeError,omitempty"`
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
					TargetingClauseSetupError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"targetingClauseSetupError,omitempty"`
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
	} `json:"targetingClauses,omitempty"`
}

func (r *CreateTargetingClausesRequestV3) getNewResponse() ICreateTargetingClausesResponse {
	return new(CreateTargetingClausesResponseV3)
}
func (r *CreateTargetingClausesRequestV3) getVersion() int {
	return 3
}

func (r *CreateTargetingClausesResponseV3) getVersion() int {
	return 3
}

//endregion CreateV3

// region DeleteV3
type DeleteTargetingClausesRequestV3 struct {
	TargetIDFilter *IDFilter `json:"targetIdFilter,omitempty"`
}
type DeleteTargetingClausesSuccess struct {
	TargetingClause struct {
		Expression []struct {
			Type  string `json:"type,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"expression,omitempty"`
		TargetID           string `json:"targetId,omitempty"`
		ResolvedExpression []struct {
			Type  string `json:"type,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"resolvedExpression,omitempty"`
		CampaignID     string  `json:"campaignId,omitempty"`
		ExpressionType string  `json:"expressionType,omitempty"`
		State          string  `json:"state,omitempty"`
		Bid            float64 `json:"bid,omitempty"`
		AdGroupID      string  `json:"adGroupId,omitempty"`
		ExtendedData   struct {
			LastUpdateDateTime   time.Time `json:"lastUpdateDateTime,omitempty"`
			ServingStatus        string    `json:"servingStatus,omitempty"`
			ServingStatusDetails []struct {
				Name    string `json:"name,omitempty"`
				HelpURL string `json:"helpUrl,omitempty"`
				Message string `json:"message,omitempty"`
			} `json:"servingStatusDetails,omitempty"`
			CreationDateTime time.Time `json:"creationDateTime,omitempty"`
		} `json:"extendedData,omitempty"`
	} `json:"targetingClause,omitempty"`
	TargetID string `json:"targetId,omitempty"`
	Index    int    `json:"index,omitempty"`
}
type DeleteTargetingClausesResponseV3 struct {
	TargetingClauses struct {
		Success []DeleteTargetingClausesSuccess `json:"success,omitempty"`
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
					ExpressionTypeError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"expressionTypeError,omitempty"`
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
					TargetingClauseSetupError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"targetingClauseSetupError,omitempty"`
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
	} `json:"targetingClauses,omitempty"`
}

func (r *DeleteTargetingClausesRequestV3) getNewResponse() IDeleteTargetingClausesResponse {
	return new(DeleteTargetingClausesResponseV3)
}
func (r *DeleteTargetingClausesRequestV3) getVersion() int {
	return 3
}

func (r *DeleteTargetingClausesResponseV3) getVersion() int {
	return 3
}

//endregion DeleteV3

// region UpdateV3
type UpdateTargetingClausesData struct {
	Expression     []TypeValueFilter `json:"expression,omitempty"`
	TargetID       string            `json:"targetId,omitempty"`
	ExpressionType string            `json:"expressionType,omitempty"`
	State          string            `json:"state,omitempty"`
	Bid            float64           `json:"bid,omitempty"`
}
type UpdateTargetingClausesRequestV3 struct {
	TargetingClauses []UpdateTargetingClausesData `json:"targetingClauses,omitempty"`
}
type UpdateTargetingClausesSuccess struct {
	TargetingClause struct {
		Expression []struct {
			Type  string `json:"type,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"expression,omitempty"`
		TargetID           string `json:"targetId,omitempty"`
		ResolvedExpression []struct {
			Type  string `json:"type,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"resolvedExpression,omitempty"`
		CampaignID     string  `json:"campaignId,omitempty"`
		ExpressionType string  `json:"expressionType,omitempty"`
		State          string  `json:"state,omitempty"`
		Bid            float64 `json:"bid,omitempty"`
		AdGroupID      string  `json:"adGroupId,omitempty"`
		ExtendedData   struct {
			LastUpdateDateTime   time.Time `json:"lastUpdateDateTime,omitempty"`
			ServingStatus        string    `json:"servingStatus,omitempty"`
			ServingStatusDetails []struct {
				Name    string `json:"name,omitempty"`
				HelpURL string `json:"helpUrl,omitempty"`
				Message string `json:"message,omitempty"`
			} `json:"servingStatusDetails,omitempty"`
			CreationDateTime time.Time `json:"creationDateTime,omitempty"`
		} `json:"extendedData,omitempty"`
	} `json:"targetingClause,omitempty"`
	TargetID string `json:"targetId,omitempty"`
	Index    int    `json:"index,omitempty"`
}
type UpdateTargetingClausesResponseV3 struct {
	TargetingClauses struct {
		Success []UpdateTargetingClausesSuccess `json:"success,omitempty"`
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
					ExpressionTypeError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"expressionTypeError,omitempty"`
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
					TargetingClauseSetupError *struct {
						Reason      string `json:"reason,omitempty"`
						Marketplace string `json:"marketplace,omitempty"`
						Cause       struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"targetingClauseSetupError,omitempty"`
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
	} `json:"targetingClauses,omitempty"`
}

func (r *UpdateTargetingClausesRequestV3) getNewResponse() IUpdateTargetingClausesResponse {
	return new(UpdateTargetingClausesResponseV3)
}
func (r *UpdateTargetingClausesRequestV3) getVersion() int {
	return 3
}

func (r *UpdateTargetingClausesResponseV3) getVersion() int {
	return 3
}

//endregion UpdateV3
