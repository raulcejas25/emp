package handler

import (
	"encoding/json"
	"net/http"
	dto "emp/internal/dto"
	service "emp/internal/service"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories := h.service.GetCategories()

	var response []dto.CategoryResponse
	for _, c := range categories {
		response = append(response, dto.ToCategoryResponse(c))
	}
	w.Header().Set("Contente-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	category, err := h.service.CreateCategory(req.Name, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := dto.ToCategoryResponse(category)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
