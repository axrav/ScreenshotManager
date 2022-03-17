package routers

import (
	"fmt"
	"log"
	"net/http"

	manager "ssmanager/controllers"

	"github.com/gorilla/mux"
)

// the routor function
func Router() *mux.Router {
	app := mux.NewRouter() // mux router
	port := manager.ThePort()
	portUsed := fmt.Sprintf(":%v", port)

	// managing handlers
	app.HandleFunc("/upload", manager.UploadScreenshot).Methods("POST")
	app.HandleFunc("/screen/{path}", manager.SendScreenshot).Methods("GET")
	
	// starting server
	log.Fatal(http.ListenAndServe(portUsed, app))
	return app
}