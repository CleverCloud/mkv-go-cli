package connect

import (
	"crypto/tls"
	"os"

	"github.com/redis/go-redis/v9"
)

func GetClient() *redis.Client {
	host := "materiakv.eu-fr-1.services.clever-cloud.com"
	port := "6379"
	password := os.Getenv("KV_TOKEN")

	return redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	})
}
