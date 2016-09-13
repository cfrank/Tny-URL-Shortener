/*
 * This is the entry point for shortened links
 * When a shortened link is clicked it gets routed through this package
 * and is handled with the UnShorten() function
 *
 * Any other helper functions are included under this package and
 * housed in this directory
 */

package handlelink

import (
	"fmt"
	"net/http"
)

func UnShorten(w http.ResponseWriter, req *http.Request, params map[string]string) {
	linkId := params["linkId"]
	fmt.Fprintf(w, "The link with id %s was requested!\r\n", linkId)
}

func Hello(input string) {
	fmt.Printf("%s", input)
}
