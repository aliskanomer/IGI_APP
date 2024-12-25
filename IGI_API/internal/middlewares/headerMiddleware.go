// Description: Header middleware sets the response headers like content-type for the API response.
package middlewares

import "net/http"

func HeaderMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {

		resp.Header().Set("Content-Type", "application/json") // set content-type as application/json
		// nextTick()
		next(resp, req)
	})
}
