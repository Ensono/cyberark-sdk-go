package models

//ApiResult - Details on the result of the operation.
type ApiResult struct {
	DeveloperMessage *string `json:"developerMessage,omitempty"`
	Message          *string `json:"message,omitempty"`
	ResultId         *string `json:"resultId,omitempty"`
}
