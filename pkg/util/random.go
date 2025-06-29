package util

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/paper-indonesia/pg-mcp-server/pkg/redisExt"
)

// GenerateRandomString will generate a random string of 10 characters.
func GenerateRandomString() string {
	timestamp := time.Now().UnixNano() // Format the current time as a timestamp string

	return fmt.Sprintf("%d", timestamp)
}

func GenerateRandomAlphanumeric(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)

}

// GetReferenceNumber will generate a reference number for the given group. It will saved at redis.
// If any error from redis, it will call GenerateRandomString()
func GetReferenceNumber(ctx context.Context, cache redisExt.IRedisExt, group string) string {
	key := fmt.Sprintf("%s:reffnum", group)
	referenceNumber, err := cache.Get(ctx, key).Result()
	if err != nil {
		referenceNumber = "1"
		err = cache.Set(ctx, key, referenceNumber, 24*time.Hour).Err()
		if err != nil {
			return GenerateRandomString()
		}
	}

	if err := cache.Incr(ctx, key).Err(); err != nil {
		return GenerateRandomString()
	}

	return fmt.Sprintf("%s%s", GenerateRandomString(), referenceNumber)
}
