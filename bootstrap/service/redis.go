package service

import (
	"context"
	"fmt"
	"time"

	"go-gin-project/library/config"
	"go-gin-project/library/resource"

	"github.com/redis/go-redis/v9"
)

// InitRedis initializes the Redis database connection.
//
// This function reads the Redis configuration from the global RedisConfig,
// initializes the Redis client using the configuration, and assigns the
// client to the global RedisClient resource.
func InitRedis(_ context.Context) {
	err := InitRedisClient()
	if err != nil {
		// Panic if the Redis client cannot be initialized.
		panic(err.Error())
	}
}

// InitRedisClient initializes the Redis client using the global RedisConfig.
// It sets up the Redis client with the configuration options specified in the
// RedisConfig and tests the connection by pinging the Redis server.
//
// The Redis client is configured with the following options:
//
// - Addr: The address of the Redis server.
// - Password: The password to use for authentication.
// - DB: The number of the Redis database to use.
// - PoolSize: The maximum number of connections to keep in the connection pool.
// - MinIdleConns: The minimum number of idle connections to keep in the connection pool.
// - MaxRetries: The maximum number of times to retry a failed operation.
// - DialTimeout: The timeout for establishing a connection to the Redis server.
// - ReadTimeout: The timeout for reading from the Redis server.
// - WriteTimeout: The timeout for writing to the Redis server.
//
// Returns an error if the connection to the Redis server cannot be established.
func InitRedisClient() error {
	// Retrieve the Redis configuration
	cfg := config.RedisConfig

	// Initialize the Redis client with the specified options
	RedisClient := redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.Addr,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
		PoolSize:     cfg.Redis.PoolSize,
		MinIdleConns: cfg.Redis.MinIdleConns,
		MaxRetries:   cfg.Redis.MaxRetries,
		DialTimeout:  time.Duration(cfg.Redis.DialTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(cfg.Redis.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(cfg.Redis.WriteTimeout) * time.Millisecond,
	})

	// Test the connection to the Redis server with a ping
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Ping the Redis server
	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		return fmt.Errorf("redis ping failed: %w", err)
	}

	// Store the Redis client instance in the global resource
	resource.RedisClient = RedisClient
	return nil
}
