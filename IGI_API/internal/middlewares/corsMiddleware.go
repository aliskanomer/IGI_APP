// Description: CORS middleware sets the response headers to allow cross-origin requests.

package middlewares

import (
	// packages
	"IGI_API/internal/models"
	"IGI_API/internal/utils"

	// modules
	"net/http"
)

func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		// CORS Headers
		resp.Header().Set("Access-Control-Allow-Origin", "*")                            // Allow all origins
		resp.Header().Set("Access-Control-Allow-Methods", "GET")                         // Allow only GET requests
		resp.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allow only Content-Type and Authorization headers

		// Preflight OPTIONS request
		if req.Method == http.MethodOptions {
			resp.WriteHeader(http.StatusOK)
			return
		}

		// Method Validation (Only GET requests are allowed)
		if req.Method != http.MethodGet {
			utils.SendErrorResponse(resp, "ERR: Only GET requests are allowed", http.StatusMethodNotAllowed, models.ErrorInfo{
				Code:    "INVALID_METHOD",
				Details: "This API supports only GET requests.",
			})
			return
		}

		// nextTick()
		next(resp, req)
	}
}
