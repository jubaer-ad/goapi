package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jubaer-ad/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	var server string = "localhost:"
	var port string = "8000"

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting Card Count Service (GO API)")
	err := http.ListenAndServe(server+port, r)
	if err != nil {
		log.Error(err)
	}
	fmt.Printf("Server started at port %v\n", port)
}
