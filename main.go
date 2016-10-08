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
	"os"
	"os/signal"

	"github.com/cfrank/tny.al/admin"
	"github.com/cfrank/tny.al/api"
	"github.com/cfrank/tny.al/database"
)

func main() {
	const PORT string = ":8081"

	fmt.Printf("Starting web server on %s...", PORT)

	// Define the mux
	router := httptreemux.New()

	// Handles all link requests
	//router.GET("/:linkId", shortlink.UnShorten)

	// Handles all API routes
	apiRoute := router.NewGroup("/api")
	apiRoute.GET("/uid", api.GetUid)
	apiRoute.POST("/add", api.Add)

	// Handles all Admin routes
	adminRoute := router.NewGroup("/admin")
	adminRoute.GET("/:task", admin.Handle)

	// Open a connection to the database
	databaseError := database.OpenDatabase()
	if databaseError != nil && database.MyDb.Alive {
		log.Fatal("Error: Problem establishing database connection!")
	}

	// Close the database when the server drops
	defer database.CloseDatabase()

	// Catch SIGTERM and close the database
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt)
	go func() {
		<-termChan
		database.CloseDatabase()
		os.Exit(1)
	}()

	// Serve
	log.Fatal(http.ListenAndServe(PORT, router))
}
