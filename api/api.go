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
	"github.com/cfrank/tny.al/utils"
	"net/http"
	"time"
)

type Link struct {
	Linkid     string    `json:"linkid"`
	Source     string    `json:"source"`
	Created    time.Time `json:"created"`
	Userid     string    `json:"userid"`
	Abuseflags uint16    `json:"abuseflags"`
	Clicks     uint      `json:"clicks"`
	Expires    time.Time `json:"expires"`
}

type Test struct {
	Url string `json:"url"`
}

/*
 * Generates a UID
 *
 * Creates a UID hashed with a secret key to make sure it isn't
 * tampered with.
 */
func GetUid(w http.ResponseWriter, req *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	var id *utils.Uid = utils.CreateId(10)
	json.NewEncoder(w).Encode(id)
}

/*
 * Shortens a Link
 *
 * This is the function which will take the request to shorten a Link
 * from the user. And Return a shortened link on success and an error
 * when something goes wrong
 */
func Add(w http.ResponseWriter, req *http.Request, params map[string]string) {
	var newTest Test
	if req.Body != nil {
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
