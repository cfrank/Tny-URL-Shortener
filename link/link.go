/*
 * This package deals with the link
 *
 * Any manipulation of the link will be handled here
 */

package link

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
