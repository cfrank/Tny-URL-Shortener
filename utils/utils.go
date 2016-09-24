package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"time"
)

type Uid struct {
	UserId string `json:"uid"`
	Key    string `json:"key"`
}

// Todo: Make sure this is the best way
func ComputeHmac(message []byte) string {
	hmac := hmac.New(sha256.New, HMAC_SECRET)
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

// Generate a random string of provided length
// This is used to create the link and user ids
func CreateId(length int) *Uid {
	const keyspace = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	id := make([]byte, length)
	for i := range id {
		id[i] = keyspace[rand.Int63()%int64(len(keyspace))]
	}

	return &Uid{
		UserId: string(id),
		Key:    ComputeHmac(id),
	}
}
