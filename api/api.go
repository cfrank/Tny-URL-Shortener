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
	"net/http"

	"github.com/cfrank/tny.al/database"
	"github.com/cfrank/tny.al/link"
	"github.com/cfrank/tny.al/utils"
)

type linkEndPoint struct {
	Linkid  string `json:"linkid"`
	Created int64  `json:"created"`
	Success bool   `json:"success"`
}

/*
 * Generates a UID
 *
 * Creates a UID hashed with a secret key to make sure it isn't
 * tampered with.
 */
func GetUid(w http.ResponseWriter, req *http.Request, params map[string]string) {
	uid, err := utils.CreateId(10)
	// Most likely the server couldn't find a UserId
	if err != nil {
		uidError := &APIError{
			Msg:        err.Error(),
			Httpstatus: http.StatusInternalServerError,
		}
		uidError.NewApiError(&w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uid)
}

/*
 * Shortens a Link
 *
 * This is the function which will take the request to shorten a Link
 * from the user. And Return a shortened link on success and an error
 * when something goes wrong
 */
func Add(w http.ResponseWriter, req *http.Request, params map[string]string) {
	var shortLink *link.Link = link.NewLink()

	if req.Body == nil {
		err := &APIError{
			Msg:        "Nothing was sent in the body!",
			Httpstatus: http.StatusBadRequest,
		}
		err.NewApiError(&w)
		return
	}

	jsonErr := json.NewDecoder(req.Body).Decode(&shortLink)

	if jsonErr != nil {
		err := &APIError{
			Msg:        "Malformed JSON recieved",
			Httpstatus: http.StatusBadRequest,
		}
		err.NewApiError(&w)
		return
	}

	// Make sure that the userid and userkey are correct
	validUserId := utils.CheckHmac(shortLink.Userkey, []byte(shortLink.Userid))

	if validUserId != true {
		err := &APIError{
			Msg:        "Invalid userid key pair recieved!",
			Httpstatus: http.StatusUnauthorized,
		}
		err.NewApiError(&w)
		return
	}

	// Generate a link id for the new link
	linkId, linkidError := utils.GenerateLinkId(6)
	if linkidError != true {
		err := &APIError{
			Msg:        "Unable to generate linkid",
			Httpstatus: http.StatusInternalServerError,
		}
		err.NewApiError(&w)
		return
	}

	// Add the linkid to the Link struct
	shortLink.Linkid = string(linkId)

	// Save the Link to the database
	savedLink := database.SaveLink(shortLink)
	if savedLink != true {
		err := &APIError{
			Msg:        "Unable to save link to database",
			Httpstatus: http.StatusInternalServerError,
		}
		err.NewApiError(&w)
		return
	}

	// Return the linkid to the user
	var endpoint linkEndPoint = linkEndPoint{
		Linkid:  shortLink.Linkid,
		Created: shortLink.Created,
		Success: true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(endpoint)
}
