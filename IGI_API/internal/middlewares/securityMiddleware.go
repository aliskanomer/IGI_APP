// Description: Security middleware sets the response security headers to prevent common web vulnerabilities.
package middlewares

import "net/http"

func SecurityMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		// Security Headers
		resp.Header().Set("X-Content-Type-Options", "nosniff")                                // Prevent MIME type sniffing
		resp.Header().Set("X-Frame-Options", "DENY")                                          // Prevent Clickjacking
		resp.Header().Set("X-XSS-Protection", "1; mode=block")                                // Prevent Cross-site scripting (XSS)
		resp.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains") // Enable HSTS

		// nextTick()
		next(resp, req)
	}
}
