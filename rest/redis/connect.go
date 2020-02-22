package redisutils

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/betonr/go-utils/base"
	"github.com/gomodule/redigo/redis"
)

var Pool *redis.Pool

func init() {
	redisHost := base.GetBetween([]string{os.Getenv("REDIS_HOST"), ""})
	redisPort := base.GetBetween([]string{os.Getenv("REDIS_PORT"), "6379"})
	redisPassword := base.GetBetween([]string{os.Getenv("REDIS_PASS"), "redis"})
	host := redisHost + ":" + redisPort
	Pool = newPool(host, redisPassword)
	cleanupHook()
}

func newPool(server string, password string) *redis.Pool {
	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}
