package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/groozin/sailor/internal/models"
	"github.com/groozin/sailor/internal/utils"
)

func IsPrimeHandler(w http.ResponseWriter, r *http.Request) {
	var req models.PrimeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	isPrime := utils.IsPrime(req.Number)

	resp := models.PrimeResponse{Result: isPrime, Message: "success"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
