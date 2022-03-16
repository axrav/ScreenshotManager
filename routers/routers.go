package routers

import (
	"log"
	"net/http"

	manager "ssmanager/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	app := mux.NewRouter()
	app.HandleFunc("/upload", manager.UploadScreenshot).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", app))
	return app
}