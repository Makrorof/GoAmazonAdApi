package sponsoredProducts

import "time"

type IListKeywordsRequest interface {
	getNewResponse() IListKeywordsResponse
	getVersion() int
}
type IListKeywordsResponse interface {
	getVersion() int
}

type ICreateKeywordsRequest interface {
	getNewResponse() ICreateKeywordsResponse
	getVersion() int
}
type ICreateKeywordsResponse interface {
	getVersion() int
}

type IDeleteKeywordsRequest interface {
	getNewResponse() IDeleteKeywordsResponse
	getVersion() int
}
type IDeleteKeywordsResponse interface {
	getVersion() int
}

type IUpdateKeywordsRequest interface {
	getNewResponse() IUpdateKeywordsResponse
	getVersion() int
}
type IUpdateKeywordsResponse interface {
	getVersion() int
}

type KeywordsDataV3 struct {
	KeywordID             string  `json:"keywordId"`
	NativeLanguageKeyword string  `json:"nativeLanguageKeyword"`
	NativeLanguageLocale  string  `json:"nativeLanguageLocale"`
	CampaignID            string  `json:"campaignId"`
	MatchType             string  `json:"matchType"`
	State                 string  `json:"state"`
	Bid                   float64 `json:"bid"`
	AdGroupID             string  `json:"adGroupId"`
	KeywordText           string  `json:"keywordText"`
	ExtendedData          struct {
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

type UpdateKeywordDataV3 struct {
	KeywordID string  `json:"keywordId"`
	State     string  `json:"state"`
	Bid       float64 `json:"bid"`
}

// region Listv3
type ListKeywordsRequestV3 struct {
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
	IncludeExtendedDataFields bool   `json:"includeExtendedDataFields"`
	Locale                    string `json:"locale"`
	KeywordTextFilter         struct {
		QueryTermMatchType string   `json:"queryTermMatchType"`
		Include            []string `json:"include"`
	} `json:"keywordTextFilter"`
	KeywordIDFilter struct {
		Include []string `json:"include"`
	} `json:"keywordIdFilter"`
	MatchTypeFilter []string `json:"matchTypeFilter"`
}
type ListKeywordsResponseV3 struct {
	TotalResults int              `json:"totalResults"`
	Keywords     []KeywordsDataV3 `json:"keywords"`
	NextToken    string           `json:"nextToken"`
}

func (r *ListKeywordsRequestV3) getNewResponse() IListKeywordsResponse {
	return new(ListKeywordsResponseV3)
}
func (r *ListKeywordsRequestV3) getVersion() int {
	return 3
}

func (r *ListKeywordsResponseV3) getVersion() int {
	return 3
}

//endregion Listv3

// region CreateV3
type CreateKeywordsV3 struct {
	NativeLanguageKeyword string  `json:"nativeLanguageKeyword"`
	NativeLanguageLocale  string  `json:"nativeLanguageLocale"`
	CampaignID            string  `json:"campaignId"`
	MatchType             string  `json:"matchType"`
	State                 string  `json:"state"`
	Bid                   float64 `json:"bid"`
	AdGroupID             string  `json:"adGroupId"`
	KeywordText           string  `json:"keywordText"`
}
type CreateKeywordsSuccessV3 struct {
	KeywordID string         `json:"keywordId"`
	Index     int            `json:"index"`
	Keyword   KeywordsDataV3 `json:"keyword"`
}
type CreateKeywordsErrorV3 struct {
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
		LocaleError *struct {
			Reason string `json:"reason"`
			Cause  struct {
				Location string `json:"location"`
				Trigger  string `json:"trigger"`
			} `json:"cause"`
			Message string `json:"message"`
		} `json:"localeError"`
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
}
type CreateKeywordsErrorsV3 struct {
	Index  int                     `json:"index"`
	Errors []CreateKeywordsErrorV3 `json:"errors"`
}
type CreateKeywordsRequestV3 struct {
	Keywords []CreateKeywordsV3 `json:"keywords"`
}
type CreateKeywordsResponseV3 struct {
	Keywords struct {
		Success []CreateKeywordsSuccessV3 `json:"success"`
		Error   []CreateKeywordsErrorV3   `json:"error"`
	} `json:"keywords"`
}

func (r *CreateKeywordsRequestV3) getNewResponse() ICreateKeywordsResponse {
	return new(CreateKeywordsResponseV3)
}
func (r *CreateKeywordsRequestV3) getVersion() int {
	return 3
}

func (r *CreateKeywordsResponseV3) getVersion() int {
	return 3
}

//endregion CreateV3

// region DeleteV3
type DeleteKeywordsRequestV3 struct {
	KeywordIDFilter struct {
		Include []string `json:"include"`
	} `json:"keywordIdFilter"`
}
type DeleteKeywordsResponseV3 struct {
	Keywords struct {
		Success []struct {
			KeywordID string         `json:"keywordId"`
			Index     int            `json:"index"`
			Keyword   KeywordsDataV3 `json:"keyword"`
		} `json:"success"`
		Error []struct {
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
					LocaleError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"localeError"`
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
	} `json:"keywords"`
}

func (r *DeleteKeywordsRequestV3) getNewResponse() IDeleteKeywordsResponse {
	return new(DeleteKeywordsResponseV3)
}
func (r *DeleteKeywordsRequestV3) getVersion() int {
	return 3
}

func (r *DeleteKeywordsResponseV3) getVersion() int {
	return 3
}

//endregion DeleteV3

// region DeleteV3
type UpdateKeywordsRequestV3 struct {
	Keywords []UpdateKeywordDataV3 `json:"keywords"`
}
type UpdateKeywordsResponseV3 struct {
	Keywords struct {
		Success []struct {
			KeywordID string         `json:"keywordId"`
			Index     int            `json:"index"`
			Keyword   KeywordsDataV3 `json:"keyword"`
		} `json:"success"`
		Error []struct {
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
					LocaleError *struct {
						Reason string `json:"reason"`
						Cause  struct {
							Location string `json:"location"`
							Trigger  string `json:"trigger"`
						} `json:"cause"`
						Message string `json:"message"`
					} `json:"localeError"`
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
	} `json:"keywords"`
}

func (r *UpdateKeywordsRequestV3) getNewResponse() IUpdateKeywordsResponse {
	return new(UpdateKeywordsResponseV3)
}
func (r *UpdateKeywordsRequestV3) getVersion() int {
	return 3
}

func (r *UpdateKeywordsResponseV3) getVersion() int {
	return 3
}

//endregion DeleteV3
