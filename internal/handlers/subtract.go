package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/groozin/sailor/internal/models"
)

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	var req models.SubtractRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	difference := req.A - req.B

	resp := models.Response{Result: difference, Message: "success"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
