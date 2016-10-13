package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"math/rand"
	"time"

	"github.com/cfrank/tny.al/constants"
	"github.com/cfrank/tny.al/database"
)

type Uid struct {
	UserId string `json:"uid"`
	Key    string `json:"key"`
}

// Todo: Make sure this is the best way
func ComputeHmac(message []byte) string {
	hmac := hmac.New(sha256.New, constants.HMAC_SECRET)
	hmac.Write(message)
	return base64.StdEncoding.EncodeToString(hmac.Sum(nil))
}

func CheckHmac(uid []byte, hmac string) bool {
	var valid string = ComputeHmac(uid)
	if hmac == valid {
		return true
	} else {
		return false
	}
}

func generateId(length int) []byte {
	const keyspace = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	id := make([]byte, length)
	for i := range id {
		id[i] = keyspace[rand.Int63()%int64(len(keyspace))]
	}

	return id
}

// Generate a random string of provided length for use as
// a UserId
func generateUserId(length int) ([]byte, bool) {
	// Try MAX_TRIES times to find a user id
	// Return false if one can't be found
	for i := 0; i < constants.MAX_TRIES; i++ {
		id := generateId(length)
		if database.UniqueUserIdCheck(string(id)) {
			return id, true
		}
	}

	return nil, false
}

// Generate a random string of provided length for use as
// a LinkId
func GenerateLinkId(length int) ([]byte, bool) {
	// Try MAX_TRIES times to find a link id
	// Return false if one can't be found
	for i := 0; i < constants.MAX_TRIES; i++ {
		id := generateId(length)
		linkidExists, linkidError := database.LinkidExists(string(id))
		if linkidExists == false && linkidError == nil {
			// The id is available! use it
			return id, true
		} else if linkidError != nil {
			// There was an error with the database
			return nil, false
		}
	}

	return nil, false
}

// Generate a Uid and with a secure key to provide assurance that it
// is correct
func CreateId(length int) (*Uid, error) {
	// Generate a user id
	id, succeeded := generateUserId(length)
	if succeeded != true {
		return nil, errors.New("Failed to find an unused UserId!")
	}

	// Add the userid to the database to prevent re-use
	userIdCreated := database.SaveUserid(string(id))
	if userIdCreated != true {
		return nil, errors.New("Failed to add UserId to database!")
	}

	return &Uid{
		UserId: string(id),
		Key:    ComputeHmac(id),
	}, nil
}
