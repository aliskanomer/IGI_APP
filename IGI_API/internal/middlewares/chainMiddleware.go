// Description: ChainMiddleware function chains the middlewares to the handler function. It helps to bind multiple middlewares to the handler function like CORS, Security, and Header middlewares.

package middlewares

import "net/http"

// Declare middleware type
type Middleware func(http.HandlerFunc) http.HandlerFunc

// ChainMiddleware function chains the middlewares to the handler function.
func ChainMiddleware(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	// Loop through the middlewares
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
