/*
 * This is the entry point for the admin interface
 * Any call within the admin interface will be sent through here
 *
 * Any other helper functions are included under this package and
 * housed in this directory
 */

package admin

import (
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, req *http.Request, params map[string]string) {
	fmt.Fprintf(w, "From the admin: %s", params["task"])
}
