package handlelink

import (
	"fmt"
	"net/http"
)

func LinkHandler(w http.ResponseWriter, req *http.Request, params map[string]string) {
	linkId := params["linkId"]
	fmt.Fprintf(w, "The link with id %s was requested\r\n", linkId)
}
