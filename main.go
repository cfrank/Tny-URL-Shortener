/*
 * This is the main entry point of the application all non-static
 * routes are sent here from nginx and routed to their respective
 * handlers, namely:
 * - API
 * - link
 * - Admin
 */

package main

import (
	"fmt"
	"github.com/idawes/httptreemux"
	"log"
	"net/http"

	"github.com/cfrank/tny.al/api"
	"github.com/cfrank/tny.al/handlelink"
)

func apiHandle(w http.ResponseWriter, req *http.Request, params map[string]string) {
	fmt.Print("HEllo from api")
	handlelink.Hello("CHris")
	api.Handle()
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
	router.GET("/:linkId", handlelink.UnShorten)

	// Handles all API routes
	api := router.NewGroup("/api")
	api.GET("/hello", apiHandle)

	// Handles all Admin routes
	admin := router.NewGroup("/admin")
	admin.GET("/:task", adminHandle)

	log.Fatal(http.ListenAndServe(":8081", router))
}
