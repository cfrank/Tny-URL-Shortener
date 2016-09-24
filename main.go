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
	const PORT string = ":8081"

	fmt.Printf("Starting web server on %s...", PORT)

	// Define the mux
	router := httptreemux.New()

	// Handles all link requests
	router.GET("/:linkId", shortlink.UnShorten)

	// Handles all API routes
	apiRoute := router.NewGroup("/api")
	apiRoute.GET("/uid", api.GetUid)
	apiRoute.POST("/add", api.Add)

	// Handles all Admin routes
	adminRoute := router.NewGroup("/admin")
	adminRoute.GET("/:task", admin.Handle)

	// Serve
	log.Fatal(http.ListenAndServe(PORT, router))
}
