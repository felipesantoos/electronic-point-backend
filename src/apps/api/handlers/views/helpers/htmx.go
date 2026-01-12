package helpers

import (
	"eletronic_point/src/apps/api/handlers"
	"encoding/json"
)

// HTMXError sets a HX-Trigger header to show a toast on the frontend and returns the error message with the given status code.
func HTMXError(ctx handlers.RichContext, status int, message string) error {
	// Prepare the trigger data
	triggerData := map[string]interface{}{
		"show-toast": map[string]string{
			"type":    "error",
			"message": message,
		},
	}

	// Marshall to JSON to ensure safe encoding (handles quotes, etc.)
	triggerJSON, _ := json.Marshal(triggerData)

	// Adds HX-Trigger header to trigger the 'show-toast' event in app.js
	ctx.Response().Header().Set("HX-Trigger", string(triggerJSON))

	// Return the message as plain text with the error status
	return ctx.String(status, message)
}
