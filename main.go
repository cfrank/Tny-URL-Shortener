/*
	Package main

	-- Entry file for tny.al backend

	-- Incoming routes are sent to their respective handler functions
*/
package main

import (
	"fmt"
	"github.com/idawes/httptreemux"
	"log"
	"net/http"
)

func link(w http.ResponseWriter, req *http.Request, params map[string]string) {
	linkId := params["linkId"]
	fmt.Fprintf(w, "The link with id %s was requested\r\n", linkId)
}

func apiHandle(w http.ResponseWriter, req *http.Request, params map[string]string) {
	fmt.Print("HEllo from api")
}

func adminHandle(w http.ResponseWriter, req *http.Request, params map[string]string) {
	task := params["task"]
	fmt.Fprintf(w, "The admin function you requested was: %s", task)
}

func main() {
	fmt.Print("Starting web server...")
	router := httptreemux.New()

	// Handles all link requests.
	// The static index page is handled by nginx
	router.GET("/:linkId", link)

	// Handles all API routes
	api := router.NewGroup("/api")
	api.GET("/hello", apiHandle)

	// Handles all Admin routes
	admin := router.NewGroup("/admin")
	admin.GET("/:task", adminHandle)

	log.Fatal(http.ListenAndServe(":8081", router))
}
