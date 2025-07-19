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
	CampaignIDFilter struct {
		Include []string `json:"include"`
	} `json:"campaignIdFilter"`
	StateFilter struct {
		Include []string `json:"include"`
	} `json:"stateFilter"`
	ExpressionTypeFilter struct {
		Include []string `json:"include"`
	} `json:"expressionTypeFilter"`
	MaxResults     int    `json:"maxResults"`
	NextToken      string `json:"nextToken"`
	TargetIDFilter struct {
		Include []string `json:"include"`
	} `json:"targetIdFilter"`
	AsinFilter struct {
		QueryTermMatchType string   `json:"queryTermMatchType"`
		Include            []string `json:"include"`
	} `json:"asinFilter"`
	AdGroupIDFilter struct {
		Include []string `json:"include"`
	} `json:"adGroupIdFilter"`
	IncludeExtendedDataFields bool `json:"includeExtendedDataFields"`
}
type ListTargetingClausesData struct {
	Expression []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"expression"`
	TargetID           string `json:"targetId"`
	ResolvedExpression []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"resolvedExpression"`
	CampaignID     string  `json:"campaignId"`
	ExpressionType string  `json:"expressionType"`
	State          string  `json:"state"`
	Bid            float64 `json:"bid"`
	AdGroupID      string  `json:"adGroupId"`
	ExtendedData   struct {
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
type ListTargetingClausesResponseV3 struct {
	TotalResults     int                        `json:"totalResults"`
	NextToken        string                     `json:"nextToken"`
	TargetingClauses []ListTargetingClausesData `json:"targetingClauses"`
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
	Expression []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"expression"`
	CampaignID     string  `json:"campaignId"`
	ExpressionType string  `json:"expressionType"`
	State          string  `json:"state"`
	Bid            float64 `json:"bid"`
	AdGroupID      string  `json:"adGroupId"`
}
type CreateTargetingClausesRequestV3 struct {
	TargetingClauses []CreateTargetingClausesData `json:"targetingClauses"`
}
type CreateTargetingClausesSuccess struct {
	TargetingClause struct {
		Expression []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"expression"`
		TargetID           string `json:"targetId"`
		ResolvedExpression []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"resolvedExpression"`
		CampaignID     string  `json:"campaignId"`
		ExpressionType string  `json:"expressionType"`
		State          string  `json:"state"`
		Bid            float64 `json:"bid"`
		AdGroupID      string  `json:"adGroupId"`
		ExtendedData   struct {
			LastUpdateDateTime   time.Time `json:"lastUpdateDateTime"`
			ServingStatus        string    `json:"servingStatus"`
			ServingStatusDetails []struct {
				Name    string `json:"name"`
				HelpURL string `json:"helpUrl"`
				Message string `json:"message"`
			} `json:"servingStatusDetails"`
			CreationDateTime time.Time `json:"creationDateTime"`
		} `json:"extendedData"`
	} `json:"targetingClause"`
	TargetID string `json:"targetId"`
	Index    int    `json:"index"`
}
type CreateTargetingClausesResponseV3 struct {
	TargetingClauses struct {
		Success []CreateTargetingClausesSuccess `json:"success"`
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
					ExpressionTypeError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"expressionTypeError"`
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
					TargetingClauseSetupError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"targetingClauseSetupError"`
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
	} `json:"targetingClauses"`
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
	TargetIDFilter struct {
		Include []string `json:"include"`
	} `json:"targetIdFilter"`
}
type DeleteTargetingClausesSuccess struct {
	TargetingClause struct {
		Expression []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"expression"`
		TargetID           string `json:"targetId"`
		ResolvedExpression []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"resolvedExpression"`
		CampaignID     string  `json:"campaignId"`
		ExpressionType string  `json:"expressionType"`
		State          string  `json:"state"`
		Bid            float64 `json:"bid"`
		AdGroupID      string  `json:"adGroupId"`
		ExtendedData   struct {
			LastUpdateDateTime   time.Time `json:"lastUpdateDateTime"`
			ServingStatus        string    `json:"servingStatus"`
			ServingStatusDetails []struct {
				Name    string `json:"name"`
				HelpURL string `json:"helpUrl"`
				Message string `json:"message"`
			} `json:"servingStatusDetails"`
			CreationDateTime time.Time `json:"creationDateTime"`
		} `json:"extendedData"`
	} `json:"targetingClause"`
	TargetID string `json:"targetId"`
	Index    int    `json:"index"`
}
type DeleteTargetingClausesResponseV3 struct {
	TargetingClauses struct {
		Success []DeleteTargetingClausesSuccess `json:"success"`
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
					ExpressionTypeError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"expressionTypeError"`
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
					TargetingClauseSetupError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"targetingClauseSetupError"`
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
	} `json:"targetingClauses"`
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

// region DeleteV3
type UpdateTargetingClausesData struct {
	Expression []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"expression"`
	TargetID       string  `json:"targetId"`
	ExpressionType string  `json:"expressionType"`
	State          string  `json:"state"`
	Bid            float64 `json:"bid"`
}
type UpdateTargetingClausesRequestV3 struct {
	TargetingClauses []UpdateTargetingClausesData `json:"targetingClauses"`
}
type UpdateTargetingClausesSuccess struct {
	TargetingClause struct {
		Expression []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"expression"`
		TargetID           string `json:"targetId"`
		ResolvedExpression []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"resolvedExpression"`
		CampaignID     string  `json:"campaignId"`
		ExpressionType string  `json:"expressionType"`
		State          string  `json:"state"`
		Bid            float64 `json:"bid"`
		AdGroupID      string  `json:"adGroupId"`
		ExtendedData   struct {
			LastUpdateDateTime   time.Time `json:"lastUpdateDateTime"`
			ServingStatus        string    `json:"servingStatus"`
			ServingStatusDetails []struct {
				Name    string `json:"name"`
				HelpURL string `json:"helpUrl"`
				Message string `json:"message"`
			} `json:"servingStatusDetails"`
			CreationDateTime time.Time `json:"creationDateTime"`
		} `json:"extendedData"`
	} `json:"targetingClause"`
	TargetID string `json:"targetId"`
	Index    int    `json:"index"`
}
type UpdateTargetingClausesResponseV3 struct {
	TargetingClauses struct {
		Success []UpdateTargetingClausesSuccess `json:"success"`
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
					ExpressionTypeError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"expressionTypeError"`
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
					TargetingClauseSetupError *struct {
						Reason      string `json:"reason"`
						Marketplace string `json:"marketplace"`
						Cause       struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"targetingClauseSetupError"`
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
	} `json:"targetingClauses"`
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

//endregion DeleteV3
