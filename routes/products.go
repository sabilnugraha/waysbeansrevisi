package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/mysql"
	"waysbeans/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	r.HandleFunc("/products", h.FindProducts).Methods("GET")
	r.HandleFunc("/product/{id}", middleware.Auth(h.GetProduct)).Methods("GET")
	r.HandleFunc("/productimage/{id}", middleware.Auth(h.GetProductImage)).Methods("GET")
	r.HandleFunc("/product", middleware.UploadFile(h.CreateProduct)).Methods("POST")
	r.HandleFunc("/product/{id}", middleware.Auth(middleware.UploadFile(h.UpdateProduct))).Methods("PATCH")
	r.HandleFunc("/product/{id}", middleware.Auth(h.DeleteProduct)).Methods("DELETE")
	r.HandleFunc("/stock/{id}", h.UpdateProductQty).Methods("PATCH")
}
