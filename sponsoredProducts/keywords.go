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
	KeywordID             string  `json:"keywordId,omitempty"`
	NativeLanguageKeyword string  `json:"nativeLanguageKeyword,omitempty"`
	NativeLanguageLocale  string  `json:"nativeLanguageLocale,omitempty"`
	CampaignID            string  `json:"campaignId,omitempty"`
	MatchType             string  `json:"matchType,omitempty"`
	State                 string  `json:"state,omitempty"`
	Bid                   float64 `json:"bid,omitempty"`
	AdGroupID             string  `json:"adGroupId,omitempty"`
	KeywordText           string  `json:"keywordText,omitempty"`
	ExtendedData          struct {
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

type UpdateKeywordDataV3 struct {
	KeywordID string  `json:"keywordId,omitempty"`
	State     string  `json:"state,omitempty"`
	Bid       float64 `json:"bid,omitempty"`
}

// region Listv3
type ListKeywordsRequestV3 struct {
	CampaignIDFilter          *IDFilter   `json:"campaignIdFilter,omitempty"`
	StateFilter               *IDFilter   `json:"stateFilter,omitempty"`
	MaxResults                int         `json:"maxResults,omitempty"`
	NextToken                 string      `json:"nextToken,omitempty"`
	AdGroupIDFilter           *IDFilter   `json:"adGroupIdFilter,omitempty"`
	IncludeExtendedDataFields bool        `json:"includeExtendedDataFields,omitempty"`
	Locale                    string      `json:"locale,omitempty"`
	KeywordTextFilter         *NameFilter `json:"keywordTextFilter,omitempty"`
	KeywordIDFilter           *IDFilter   `json:"keywordIdFilter,omitempty"`
	MatchTypeFilter           []string    `json:"matchTypeFilter,omitempty"`
}
type ListKeywordsResponseV3 struct {
	TotalResults int              `json:"totalResults,omitempty"`
	Keywords     []KeywordsDataV3 `json:"keywords,omitempty"`
	NextToken    string           `json:"nextToken,omitempty"`
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
	NativeLanguageKeyword string  `json:"nativeLanguageKeyword,omitempty"`
	NativeLanguageLocale  string  `json:"nativeLanguageLocale,omitempty"`
	CampaignID            string  `json:"campaignId,omitempty"`
	MatchType             string  `json:"matchType,omitempty"`
	State                 string  `json:"state,omitempty"`
	Bid                   float64 `json:"bid,omitempty"`
	AdGroupID             string  `json:"adGroupId,omitempty"`
	KeywordText           string  `json:"keywordText,omitempty"`
}
type CreateKeywordsSuccessV3 struct {
	KeywordID string         `json:"keywordId,omitempty"`
	Index     int            `json:"index,omitempty"`
	Keyword   KeywordsDataV3 `json:"keyword,omitempty"`
}
type CreateKeywordsErrorV3 struct {
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
		LocaleError *struct {
			Reason string `json:"reason,omitempty"`
			Cause  struct {
				Location string `json:"location,omitempty"`
				Trigger  string `json:"trigger,omitempty"`
			} `json:"cause,omitempty"`
			Message string `json:"message,omitempty"`
		} `json:"localeError,omitempty"`
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
}
type CreateKeywordsErrorsV3 struct {
	Index  int                     `json:"index,omitempty"`
	Errors []CreateKeywordsErrorV3 `json:"errors,omitempty"`
}
type CreateKeywordsRequestV3 struct {
	Keywords []CreateKeywordsV3 `json:"keywords,omitempty"`
}
type CreateKeywordsResponseV3 struct {
	Keywords struct {
		Success []CreateKeywordsSuccessV3 `json:"success,omitempty"`
		Error   []CreateKeywordsErrorV3   `json:"error,omitempty"`
	} `json:"keywords,omitempty"`
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
	KeywordIDFilter *IDFilter `json:"keywordIdFilter,omitempty"`
}
type DeleteKeywordsResponseV3 struct {
	Keywords struct {
		Success []struct {
			KeywordID string         `json:"keywordId,omitempty"`
			Index     int            `json:"index,omitempty"`
			Keyword   KeywordsDataV3 `json:"keyword,omitempty"`
		} `json:"success,omitempty"`
		Error []struct {
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
					LocaleError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"localeError,omitempty"`
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
	} `json:"keywords,omitempty"`
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

// region UpdateV3
type UpdateKeywordsRequestV3 struct {
	Keywords []UpdateKeywordDataV3 `json:"keywords,omitempty"`
}
type UpdateKeywordsResponseV3 struct {
	Keywords struct {
		Success []struct {
			KeywordID string         `json:"keywordId,omitempty"`
			Index     int            `json:"index,omitempty"`
			Keyword   KeywordsDataV3 `json:"keyword,omitempty"`
		} `json:"success,omitempty"`
		Error []struct {
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
					LocaleError *struct {
						Reason string `json:"reason,omitempty"`
						Cause  struct {
							Location string `json:"location,omitempty"`
							Trigger  string `json:"trigger,omitempty"`
						} `json:"cause,omitempty"`
						Message string `json:"message,omitempty"`
					} `json:"localeError,omitempty"`
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
	} `json:"keywords,omitempty"`
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

//endregion UpdateV3
