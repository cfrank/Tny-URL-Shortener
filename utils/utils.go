package utils

import (
	"math/rand"
	"time"
)

// Generate a random string of provided length
// This is used to create the link and user ids
func CreateId(length int) string {
	const keyspace = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	id := make([]byte, length)
	for i := range id {
		id[i] = keyspace[rand.Int63()%int64(len(keyspace))]
	}

	return string(id)
}
