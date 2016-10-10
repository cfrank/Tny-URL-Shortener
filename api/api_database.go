package api

import (
	"database/sql"
	"errors"

	"github.com/cfrank/tny.al/database"
	"github.com/cfrank/tny.al/link"
)

/*
 * Save a link to the database
 *
 * Take a link struct and save it's contents to the database
 */
func SaveLink(shortLink *link.Link) bool {
	stmt, stmtError := database.MyDb.Db.Prepare(`INSERT INTO link(linkid, source,
					created, userid) VALUES (?, ?, ?, ?)`)
	defer stmt.Close()

	if stmtError != nil {
		return false
	}

	_, resultError := stmt.Exec(shortLink.Linkid, shortLink.Source, shortLink.Created, shortLink.Userid)

	if resultError != nil {
		return false
	}

	return true
}

/*
 * Expose a shortened links data to a provided struct
 *
 * Take a linkid and a reference to a struct and return the data to that
 * reference.
 */
func GetLinkData(linkid string, linkData *exposedLink) error {
	query := database.MyDb.Db.QueryRow(`SELECT source, created, abuseflags
                                                FROM link WHERE linkid =?`, linkid)
	queryError := query.Scan(&linkData.Source, &linkData.Created, &linkData.Abuseflags)

	if queryError != nil {
		switch {
		case queryError == sql.ErrNoRows:
			return errors.New("There were no links with the provided linkid")
		default:
			return queryError
		}
	}

	return nil
}
