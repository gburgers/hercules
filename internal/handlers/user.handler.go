// Handlers process HTTP requests and call the corresponding services.
// These functions use a web framework like net/http, gorilla/mux, or chi.
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"your-project/internal/services"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserProfile(r.Context(), id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
