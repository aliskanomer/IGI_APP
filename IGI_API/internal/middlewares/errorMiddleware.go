// Description: Error middleware is a global middleware to handle unexpected errors and panic within the application. Please check documentation for architrecture of  error handling.

package middlewares

import (
	// packages
	"IGI_API/internal/models"
	"IGI_API/internal/utils"

	// modules
	"log"
	"net/http"
)

func ErrorMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("ERR_OPS: %s_STAT:500", r)
				utils.SendErrorResponse(resp, "ERR: Internal Server Error!", http.StatusInternalServerError, models.ErrorInfo{
					Code:    "INTERNAL_SERVER_ERROR",
					Details: "An unexpected error occurred. Please try again later.",
				})
			}
		}()
		next(resp, req)
	}
}

// KISS EXPLANATION:
// If things blow up in application this middleware will catch it and send a 500 response to the client by utils.
// If there is an error on service layer. Related service will return an error to its invoker
// If there is an error on controller layer (e.g. controller recieved an error from service or some other thing went downhill...) Controller will return an error to the client by utils.

// SO -> errors are send by utils.SendErrorResponse but who is calling this function? -> controllers and error middleware. Services just returns the error to controller.
