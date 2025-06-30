package GoAmazonAdApi

import "time"

var (
	API_ERROR_MAX_RETRY   int           = 8
	API_ERROR_RETRY_DELAY time.Duration = time.Millisecond * 500
	DEBUG_PRINT_API_BODY  bool          = false
)
