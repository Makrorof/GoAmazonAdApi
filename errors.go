package GoAmazonAdApi

import "fmt"

type AmazonError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func (e *AmazonError) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s, Details: %s", e.Code, e.Message, e.Details)
}
