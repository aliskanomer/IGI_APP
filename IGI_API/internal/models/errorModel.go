package models

// ErrorInfo contains detailed error information.
type ErrorInfo struct {
	Code    string `json:"code"`    // Unique error code
	Details string `json:"details"` // Human-readable error details
}
