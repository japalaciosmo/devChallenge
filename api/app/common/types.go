package app

// Confirmation generic request confirmation.
type Confirmation struct {
	Message string `json:"message"` // confirmation message
}

// ErrorResponse standard error response data.
type ErrorResponse struct {
	ID      string                   `json:"id,omitempty"`   // ID is the unique error instance identifier.
	Code    string                   `json:"code,omitempty"` // Code identifies the class of errors.
	Status  int                      `json:"status"`         // Status is the HTTP status code used by responses that cary the error.
	Details string                   `json:"details"`        // Detail describes the specific error occurrence.
	Meta    []map[string]interface{} `json:"meta,omitempty"` // Meta contains additional key/value pairs useful to clients.
}

// SNSPayload used to extract data from sns messages and or
// sqs data sent from sns.
type SNSPayload struct {
	Message interface{} `json:"Message,string"` // the sns message body containing the event data
}
