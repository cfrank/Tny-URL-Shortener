/*
 * This is the main entry point of the application all non-static
 * routes are sent here from nginx and routed to their respective
 * handlers, namely:
 * - API
 * - shortlink
 * - Admin
 */

package main

import (
	"fmt"
	"github.com/idawes/httptreemux"
	"log"
	"net/http"

	"github.com/cfrank/tny.al/admin"
	"github.com/cfrank/tny.al/api"
	"github.com/cfrank/tny.al/shortlink"
)

func main() {
	fmt.Print("Starting web server...")

	// Handles all link requests
	router := httptreemux.New()
	router.GET("/:linkId", shortlink.UnShorten)

	// Handles all API routes
	apiRoute := router.NewGroup("/api")
	apiRoute.GET("/hello", api.Handle)

	// Handles all Admin routes
	adminRoute := router.NewGroup("/admin")
	adminRoute.GET("/:task", admin.Handle)

	log.Fatal(http.ListenAndServe(":8081", router))
}
