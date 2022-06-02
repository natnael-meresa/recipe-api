package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/natnael-meresa/recipe/api/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	router := mux.NewRouter()

	router.Path("/metrics").Handler(promhttp.Handler())
	router.HandleFunc("/recipes", GetRecipes).Methods(http.MethodGet)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	metricsMiddleware := middleware.NewMetricsMiddleware()

	router.Use(metricsMiddleware.Metrics)
	log.Printf("serve on part 9000")
	err := http.ListenAndServe(":9000", router)

	log.Fatal(err)

}

func GetRecipes(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	w.Write([]byte("Lemon"))

}
