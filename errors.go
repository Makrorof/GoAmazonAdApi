package GoAmazonAdApi

import "fmt"

type AmazonError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
}

func (e *AmazonError) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s, Details: %s", e.Code, e.Message, e.Details)
}
