// Description: Centralized logging of the application.
package utils

import "log"

// Logger logs messages with a consistent format based on type (INFO, ERR, SCC)
func Logger(logType string, operation string, statusCode int, message interface{}) {
	prefix := ""

	switch logType {
	case "info":
		prefix = "INF"
	case "error":
		prefix = "ERR"
	case "success":
		prefix = "SCC"
	default:
		prefix = "LOG"
	}

	log.Printf("%s_STAT:%d_OPS:%s %v\n", prefix, statusCode, operation, message)
}
