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

	"github.com/cfrank/tny.al/link"
	"github.com/cfrank/tny.al/utils"
)

type apiSuccess struct {
	Message    string `json:"message"`
	Httpstatus int    `json:"code"`
}

type linkEndPoint struct {
	Linkid     string `json:"linkid"`
	Created    int64  `json:"created"`
	Httpstatus int    `json:"code"`
}

type exposedLink struct {
	Source     string `json:"source"`
	Created    int64  `json:"created"`
	Abuseflags uint16 `json:"flags"`
	Httpstatus int    `json:"code"`
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
	savedLink := SaveLink(shortLink)
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
		Linkid:     shortLink.Linkid,
		Created:    shortLink.Created,
		Httpstatus: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(endpoint)
}

/*
 * Expose a shortened link
 *
 * This function will expose a already shortened link and provide some
 * information about it to the user. Including sourceurl, creation date
 * and abuse report count. Everything will be returned as json for the
 * frontend to deal with
 */
func ExposeLink(w http.ResponseWriter, req *http.Request, params map[string]string) {
	var linkid string
	var linkInfo exposedLink = exposedLink{}

	if req.Body == nil {
		err := &APIError{
			Msg:        "Nothing sent in the body!",
			Httpstatus: http.StatusBadRequest,
		}
		err.NewApiError(&w)
		return
	}

	jsonError := json.NewDecoder(req.Body).Decode(&linkid)

	if jsonError != nil {
		err := &APIError{
			Msg:        "Malformed JSON recieved",
			Httpstatus: http.StatusBadRequest,
		}
		err.NewApiError(&w)
		return
	}

	linkInfoError := GetLinkData(linkid, &linkInfo)

	if linkInfoError != nil {
		err := &APIError{
			Msg:        linkInfoError.Error(),
			Httpstatus: http.StatusInternalServerError,
		}
		err.NewApiError(&w)
		return
	}

	// There were no errors
	linkInfo.Httpstatus = http.StatusOK
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(linkInfo)
}

/*
 * Report a link
 *
 * This function will take a linkid and send it to the database so that
 * it's 'abuseflags' variable can be increased. This will allow the admin
 * to see links which have been flagged, and for action to be taken
 */
func ReportLink(w http.ResponseWriter, req *http.Request, params map[string]string) {
	var linkid string

	if req.Body == nil {
		err := &APIError{
			Msg:        "Nothing sent in the body!",
			Httpstatus: http.StatusBadRequest,
		}
		err.NewApiError(&w)
		return
	}

	jsonError := json.NewDecoder(req.Body).Decode(&linkid)

	if jsonError != nil {
		err := &APIError{
			Msg:        "Malformed JSON recieved",
			Httpstatus: http.StatusBadRequest,
		}
		err.NewApiError(&w)
		return
	}

	reportError, internalError := FlagLink(linkid)

	if reportError != nil {
		err := &APIError{
			Msg: reportError.Error(),
		}
		if internalError == true {
			err.Httpstatus = http.StatusInternalServerError
		} else {
			err.Httpstatus = http.StatusBadRequest
		}
		err.NewApiError(&w)
		return
	}

	success := &apiSuccess{
		Message:    "Report was successful",
		Httpstatus: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(success)
}
