package handler

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"performance-service/internal/service"
)

type Handler struct {
	service *service.PerformanceService
}

func NewPerformanceHandler(s *service.PerformanceService) *Handler {
	return &Handler{service: s}
}

func RegisterPerformanceRoutes(r chi.Router, h *Handler) {
	r.Route("/performance/{student_id}", func(r chi.Router) {
		r.Get("/", h.GetPerformance)
		r.Get("/grades", h.GetGrades)
		r.Get("/attendance", h.GetAttendance)
		r.Get("/debts", h.GetDebts)
		r.Get("/rating", h.GetRating)
	})
}

func (h *Handler) GetPerformance(w http.ResponseWriter, r *http.Request) {
	studentID := chi.URLParam(r, "student_id")
	perf, err := h.service.GetPerformance(r.Context(), studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(perf)
}

func (h *Handler) GetGrades(w http.ResponseWriter, r *http.Request) {
	studentID := chi.URLParam(r, "student_id")
	grades, err := h.service.GetGrades(r.Context(), studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(grades)
}

func (h *Handler) GetAttendance(w http.ResponseWriter, r *http.Request) {
	studentID := chi.URLParam(r, "student_id")
	att, err := h.service.GetAttendance(r.Context(), studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(att)
}

func (h *Handler) GetDebts(w http.ResponseWriter, r *http.Request) {
	studentID := chi.URLParam(r, "student_id")
	debts, err := h.service.GetDebts(r.Context(), studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(debts)
}

func (h *Handler) GetRating(w http.ResponseWriter, r *http.Request) {
	studentID := chi.URLParam(r, "student_id")
	rating, err := h.service.GetRating(r.Context(), studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]float64{"rating": rating})
} 