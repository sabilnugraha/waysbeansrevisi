package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	// UserRoutes(r)
	ProductRoutes(r)
	// ToppingRoutes(r)
	AuthRoutes(r)
	CartRoutes(r)
	Transaction(r)
}
