/*
 * This is the entry point for the api
 * Calls to any api endpoint are routed through here and handled with
 * the Handle() function
 *
 * Any other helper functions are included under this package and
 * housed in this directory
 */

package api

import (
	"fmt"
)

func Handle() {
	fmt.Printf("%s", "Hello World12121")
}
