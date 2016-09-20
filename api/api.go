/*
 * This is the entry point for the api
 * Calls to any api endpoint are routed through here and handled with
 * their respective functions
 *
 * Any other helper functions are included under this package and
 * housed in this directory
 */

package api

import (
	"encoding/json"
	"fmt"
	"github.com/cfrank/tny.al/utils"
	"net/http"
	"time"
)

type Link struct {
	linkid     string    `json:"linkid"`
	source     string    `json:"source"`
	created    time.Time `json:"created"`
	userid     string    `json:"userid"`
	abuseflags uint16    `json:"abuseflags"`
	clicks     uint      `json:"clicks"`
	expires    time.Time `json:"expires"`
}

type Test struct {
	Url string
}

func Add(w http.ResponseWriter, req *http.Request, params map[string]string) {
	var newTest Test
	if req.Body == nil {
		err := &APIError{
			Msg:        "Nothing was sent in the body!",
			Httpstatus: http.StatusBadRequest,
		}
		err.NewApiError(&w)
		return
	}

	jsonErr := json.NewDecoder(req.Body).Decode(&newTest)

	if jsonErr != nil {
		err := &APIError{
			Msg:        "Malformed JSON recieved",
			Httpstatus: http.StatusBadRequest,
		}
		err.NewApiError(&w)
		return
	}
}
