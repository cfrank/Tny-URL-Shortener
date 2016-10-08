/*
 * This package deals with the link
 *
 * Any manipulation of the link will be handled here
 */

package link

import (
	"database/sql"
	"net/http"

	"github.com/cfrank/tny.al/database"
)

type Link struct {
	Linkid     string `json:"linkid"`
	Source     string `json:"source"`
	Created    int64  `json:"created"`
	Userid     string `json:"userid"`
	Userkey    string `json:"userkey"`
	Abuseflags uint16 `json:"abuseflags"`
	Clicks     uint   `json:"clicks"`
	Expires    int64  `json:"expires"`
}

func NewLink() *Link {
	return &Link{}
}

func Unshorten(w http.ResponseWriter, req *http.Request, params map[string]string) {
	linkid := params["linkId"]

	sourceUrl, sourceUrlError := database.GetSourceUrl(linkid)
	if sourceUrlError != nil {
		if sourceUrlError == sql.ErrNoRows {
			// There is no source url with this linkid
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Couldn't find a source url!"))
			return
		}
		// There was some other error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Database problem"))
		return
	}

	// Report the click for analytics
	newClickError := database.ReportNewClick(linkid)
	if newClickError != nil {
		w.Write([]byte(newClickError.Error()))
		return
	}

	w.Header().Set("Location", sourceUrl)
	w.WriteHeader(http.StatusMovedPermanently)
	w.Write([]byte(sourceUrl))
}
