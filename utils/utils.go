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

func CheckHmac(hmac string, uid []byte) bool {
	var valid string = ComputeHmac(uid)
	if hmac == valid {
		return true
	} else {
		return false
	}
}

// Generate a random string of provided length for use as
// a UserId
func generateUserId(length int) ([]byte, bool) {
	const keyspace = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// Try MAX_TRIES times to find a user id
	// Throw an error if it's not possible
	for i := 0; i < constants.MAX_TRIES; i++ {
		id := make([]byte, length)
		for i := range id {
			id[i] = keyspace[rand.Int63()%int64(len(keyspace))]
		}

		if database.UniqueIdCheck(string(id)) {
			return id, true
		}
	}

	return nil, false
}

// Generate a Uid and with a secure key to provide assurance that it
// is correct
func CreateId(length int) (uid *Uid, err error) {
	// Try 5 times to find a available userid that has not been
	// taken already, if one can't be found return an error
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
