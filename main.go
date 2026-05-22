package main

import (
	"fmt"
	"net/http"
	"emp/internal/handler"
	"emp/internal/service"
	"emp/internal/store"
)

func transferHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Hello world")
	}
}

func main() {
	
	store := store.NewCategoryStore()
	service := service.NewCategoryService(store)
	handler := handler.NewCategoryHandler(service)

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetCategories(w, r)
		case http.MethodPost:
			handler.CreateCategory(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}
