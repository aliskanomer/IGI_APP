package utils

import (
	// packages
	"IGI_API/internal/models"

	// modules
	"encoding/json"
	"net/http"
)

// Centralized function to send success response
func SendSuccessResponse(resp http.ResponseWriter, message string, data interface{}, statusCode int) {
	resp.WriteHeader(statusCode)    // set status code
	response := models.APIResponse{ // model response
		Message:  message,
		Response: data,
		Status:   statusCode,
	}
	json.NewEncoder(resp).Encode(response) // encode response to JSON
}

// Centralized function to send error response
func SendErrorResponse(resp http.ResponseWriter, message string, statusCode int, errorInfo models.ErrorInfo) {
	resp.WriteHeader(statusCode)    // set status code
	response := models.APIResponse{ // model response
		Message: message,
		Response: map[string]interface{}{
			"error": errorInfo,
		},
		Status: statusCode,
	}
	json.NewEncoder(resp).Encode(response) // encode response to JSON
}
