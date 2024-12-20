package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/groozin/sailor/internal/models"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	var req models.AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	sum := 0
	for _, num := range req.Numbers {
		sum += num
	}

	resp := models.AddResponse{Result: sum, Message: "success"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
