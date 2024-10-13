package drivers

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func ProviderRedisClient(addr string, opts ...func(*redis.Options)) (*redis.Client, error) {
	return provideRedisClient(addr, opts...)
}

func provideRedisClient(addr string, opts ...func(*redis.Options)) (*redis.Client, error) {
	options := &redis.Options{
		Addr: addr, // Redis server address
		DB:   0,    // Default database to use
	}

	// Apply any additional options provided
	for _, opt := range opts {
		opt(options)
	}
	client := redis.NewClient(options)
	// Perform an initial ping to check connectivity
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		log.Printf("\n error: %v\n", err)
		return nil, err // Return nil and the error if the ping fails
	}

	// Return the initialized client and nil as error if successful
	return client, nil
}
