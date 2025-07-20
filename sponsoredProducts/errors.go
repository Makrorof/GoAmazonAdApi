package sponsoredProducts

type Error struct {
	Index  int          `json:"index,omitempty"`
	Errors []ErrorValue `json:"errors,omitempty"`
}

type ErrorValue struct {
	EntityStateError            *EntityStateError            `json:"entityStateError,omitempty"`
	MissingValueError           *MissingValueError           `json:"missingValueError,omitempty"`
	BiddingError                *BiddingError                `json:"biddingError,omitempty"`
	DuplicateValueError         *DuplicateValueError         `json:"duplicateValueError,omitempty"`
	RangeError                  *RangeError                  `json:"rangeError,omitempty"`
	OtherError                  *OtherError                  `json:"otherError,omitempty"`
	ParentEntityError           *ParentEntityError           `json:"parentEntityError,omitempty"`
	ThrottledError              *ThrottledError              `json:"throttledError,omitempty"`
	EntityNotFoundError         *EntityNotFoundError         `json:"entityNotFoundError,omitempty"`
	ApplicableMarketplacesError *ApplicableMarketplacesError `json:"applicableMarketplacesError,omitempty"`
	MalformedValueError         *MalformedValueError         `json:"malformedValueError,omitempty"`
	BillingError                *BillingError                `json:"billingError,omitempty"`
	EntityQuotaError            *EntityQuotaError            `json:"entityQuotaError,omitempty"`
	InternalServerError         *InternalServerError         `json:"internalServerError,omitempty"`
	DateError                   *DateError                   `json:"dateError,omitempty"`
	BudgetError                 *BudgetError                 `json:"budgetError,omitempty"`
	CurrencyError               *CurrencyError               `json:"currencyError,omitempty"`
	TargetingClauseSetupError   *TargetingClauseSetupError   `json:"targetingClauseSetupError,omitempty"`
	LocaleError                 *LocaleError                 `json:"localeError,omitempty"`
	ExpressionTypeError         *ExpressionTypeError         `json:"expressionTypeError,omitempty"`
}

type ErrorCause struct {
	Location string `json:"location,omitempty"`
	Trigger  string `json:"trigger,omitempty"`
}
type EntityStateError struct {
	Reason      string      `json:"reason,omitempty"`
	Marketplace string      `json:"marketplace,omitempty"`
	EntityType  string      `json:"entityType,omitempty"`
	Cause       *ErrorCause `json:"cause,omitempty"`
	Message     string      `json:"message,omitempty"`
}
type MissingValueError struct {
	Reason      string      `json:"reason,omitempty"`
	Marketplace string      `json:"marketplace,omitempty"`
	Cause       *ErrorCause `json:"cause,omitempty"`
	Message     string      `json:"message,omitempty"`
}
type BiddingError struct {
	Reason      string      `json:"reason,omitempty"`
	Marketplace string      `json:"marketplace,omitempty"`
	Cause       *ErrorCause `json:"cause,omitempty"`
	UpperLimit  string      `json:"upperLimit,omitempty"`
	LowerLimit  string      `json:"lowerLimit,omitempty"`
	Message     string      `json:"message,omitempty"`
}
type DuplicateValueError struct {
	Reason      string      `json:"reason,omitempty"`
	Marketplace string      `json:"marketplace,omitempty"`
	Cause       *ErrorCause `json:"cause,omitempty"`
	Message     string      `json:"message,omitempty"`
}
type RangeError struct {
	Reason      string      `json:"reason,omitempty"`
	Marketplace string      `json:"marketplace,omitempty"`
	Allowed     []string    `json:"allowed,omitempty"`
	Cause       *ErrorCause `json:"cause,omitempty"`
	UpperLimit  string      `json:"upperLimit,omitempty"`
	LowerLimit  string      `json:"lowerLimit,omitempty"`
	Message     string      `json:"message,omitempty"`
}
type OtherError struct {
	Reason      string      `json:"reason,omitempty"`
	Marketplace string      `json:"marketplace,omitempty"`
	Cause       *ErrorCause `json:"cause,omitempty"`
	Message     string      `json:"message,omitempty"`
}
type ParentEntityError struct {
	Reason  string      `json:"reason,omitempty"`
	Cause   *ErrorCause `json:"cause,omitempty"`
	Message string      `json:"message,omitempty"`
}
type ThrottledError struct {
	Reason  string      `json:"reason,omitempty"`
	Cause   *ErrorCause `json:"cause,omitempty"`
	Message string      `json:"message,omitempty"`
}
type EntityNotFoundError struct {
	Reason     string      `json:"reason,omitempty"`
	EntityType string      `json:"entityType,omitempty"`
	Cause      *ErrorCause `json:"cause,omitempty"`
	EntityID   string      `json:"entityId,omitempty"`
	Message    string      `json:"message,omitempty"`
}
type ApplicableMarketplacesError struct {
	Reason  string      `json:"reason,omitempty"`
	Cause   *ErrorCause `json:"cause,omitempty"`
	Message string      `json:"message,omitempty"`
}
type MalformedValueError struct {
	Reason      string      `json:"reason,omitempty"`
	Fragment    string      `json:"fragment,omitempty"`
	Marketplace string      `json:"marketplace,omitempty"`
	Cause       *ErrorCause `json:"cause,omitempty"`
	Message     string      `json:"message,omitempty"`
}
type BillingError struct {
	Reason  string      `json:"reason,omitempty"`
	Cause   *ErrorCause `json:"cause,omitempty"`
	Message string      `json:"message,omitempty"`
}
type EntityQuotaError struct {
	Reason     string      `json:"reason,omitempty"`
	QuotaScope string      `json:"quotaScope,omitempty"`
	EntityType string      `json:"entityType,omitempty"`
	Quota      string      `json:"quota,omitempty"`
	Cause      *ErrorCause `json:"cause,omitempty"`
	Message    string      `json:"message,omitempty"`
}
type InternalServerError struct {
	Reason  string      `json:"reason,omitempty"`
	Cause   *ErrorCause `json:"cause,omitempty"`
	Message string      `json:"message,omitempty"`
}

type DateError struct {
	Reason  string      `json:"reason,omitempty"`
	Cause   *ErrorCause `json:"cause,omitempty"`
	Message string      `json:"message,omitempty"`
}

type BudgetError struct {
	Reason     string      `json:"reason,omitempty"`
	Cause      *ErrorCause `json:"cause,omitempty"`
	UpperLimit string      `json:"upperLimit,omitempty"`
	LowerLimit string      `json:"lowerLimit,omitempty"`
	Message    string      `json:"message,omitempty"`
}

type CurrencyError struct {
	Reason  string      `json:"reason,omitempty"`
	Cause   *ErrorCause `json:"cause,omitempty"`
	Message string      `json:"message,omitempty"`
}

type TargetingClauseSetupError struct {
	Reason      string      `json:"reason,omitempty"`
	Marketplace string      `json:"marketplace,omitempty"`
	Cause       *ErrorCause `json:"cause,omitempty"`
	Message     string      `json:"message,omitempty"`
}
type LocaleError struct {
	Reason  string      `json:"reason,omitempty"`
	Cause   *ErrorCause `json:"cause,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ExpressionTypeError struct {
	Reason  string      `json:"reason,omitempty"`
	Cause   *ErrorCause `json:"cause,omitempty"`
	Message string      `json:"message,omitempty"`
}
