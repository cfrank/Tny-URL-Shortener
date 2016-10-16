package api

import (
	"database/sql"
	"errors"

	"github.com/cfrank/tny.al/constants"
	"github.com/cfrank/tny.al/database"
	"github.com/cfrank/tny.al/link"
)

/*
 * Save a link to the database
 *
 * Take a link struct and save it's contents to the database
 */
func SaveLink(shortLink *link.Link) bool {
	stmt, stmtError := database.MyDb.Db.Prepare(`INSERT INTO link(linkid, source, created, userid) VALUES (?, ?, ?, ?)`)
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
 * Get link data
 *
 * Get from the database a history of links shortened by a user
 * based on the provided userid. Save the information to the link array
 * pointer
 */
func getLinkHistory(user *userInfo, linkData *[]linkHisory) error {
	rows, queryError := database.MyDb.Db.Query(`SELECT linkid, source, created, abuseflags, clicks FROM link WHERE userid =?`, user.Userid)

	if queryError != nil {
		return queryError
	}
	defer rows.Close()

	// Keep track of how many entries are returned
	historyEntries := 0
	for rows.Next() {
		// Only return a maximum of MAX_HISTORY entries
		if historyEntries >= constants.MAX_HISTORY {
			return nil
		}
		var link linkHisory = linkHisory{}
		scanError := rows.Scan(&link.Linkid, &link.Source, &link.Created, &link.Abuseflags, &link.Clicks)

		if scanError != nil {
			return scanError
		}

		// Add the link to the slice
		*linkData = append(*linkData, link)
		historyEntries++
	}
	return nil
}

/*
 * Expose a shortened links data to a provided struct
 *
 * Take a linkid and a reference to a struct and return the data to that
 * reference.
 */
func GetLinkData(linkid string, linkData *exposedLink) error {
	query := database.MyDb.Db.QueryRow(`SELECT source, created, abuseflags FROM link WHERE linkid =?`, linkid)
	queryError := query.Scan(&linkData.Source, &linkData.Created, &linkData.Abuseflags)

	if queryError != nil {
		switch {
		case queryError == sql.ErrNoRows:
			return errors.New("There were no links with that linkid")
		default:
			return queryError
		}
	}

	return nil
}

/*
 * Flag a link as abusive in the database
 *
 * Just increment the 'abuseflags' variable for a specified linkid
 *
 * This function returns all custom errors, as well as a bool value
 * telling the caller if the error was a server(db) error, or a invalid
 * request (invalid linkid)
 */
func FlagLink(linkid string) (error, bool) {
	// Check to make sure the linkid exists
	linkidExists, linkidError := database.LinkidExists(linkid)
	if linkidExists == false {
		return errors.New("That linkid does not exist"), false
	} else if linkidError != nil {
		return errors.New("There was a problem with the database"), true
	}

	// Linkid exists, now flag it
	stmt, stmtError := database.MyDb.Db.Prepare(`UPDATE link SET abuseflags = abuseflags + 1 WHERE linkid =?`)
	defer stmt.Close()

	if stmtError != nil {
		return errors.New("There was a problem with the database"), true
	}

	_, resultError := stmt.Exec(linkid)

	if resultError != nil {
		return errors.New("There was a problem with the database"), true
	}

	return nil, false
}
