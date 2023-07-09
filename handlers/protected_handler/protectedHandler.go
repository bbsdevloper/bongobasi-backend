package protected_handler

import (
	"encoding/json"
	"net/http"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	// Access the user ID from the token claims
	phone := r.Context().Value("phone").(string)

	// Perform actions specific to the protected route

	// Return a response
	response := map[string]interface{}{
		"message": "Protected route accessed",
		"phone":   phone,
	}
	json.NewEncoder(w).Encode(response)
}
