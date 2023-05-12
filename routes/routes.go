package router

import (
	controller "github.com/Hasan-Iqtedar/crud-app-go/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controller.InsertOneMovie).Methods("POST")
	router.HandleFunc("/api/movies", controller.DeleteAllMovies).Methods("DELETE")
	router.HandleFunc("/api/movies/{movieId}", controller.UpdateOneMovie).Methods("PUT")
	router.HandleFunc("/api/movies/{movieId}", controller.DeleteOneMovie).Methods("DELETE")

	return router
}
