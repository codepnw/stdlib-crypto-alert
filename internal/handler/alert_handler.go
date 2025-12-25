package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stdlib-crypto-alert/internal/service"
	"github.com/stdlib-crypto-alert/pkg/validate"
)

type alertHandler struct {
	srv service.AlertService
}

func NewAlertHandler(srv service.AlertService) *alertHandler {
	return &alertHandler{srv: srv}
}

func (h *alertHandler) CreateAlertHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allow", http.StatusMethodNotAllowed)
		return
	}
	
	// JSON Decode
	req := new(CreateAlertReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, fmt.Sprintf("invalid json decode: %s", err.Error()), http.StatusBadRequest)
		return
	}
	
	// Validate Request
	if err := validate.Struct(req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request: %s", err.Error()), http.StatusBadRequest)
		return
	}
	
	// Create Alert
	if err := h.srv.CreateAlert(r.Context(), req.Symbol, req.TargetPrice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// JSON Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	resp := JSONResponse{
		Code:    http.StatusCreated,
		Message: "alert created successfully",
		Data:    nil,
	}
	json.NewEncoder(w).Encode(resp)
}
