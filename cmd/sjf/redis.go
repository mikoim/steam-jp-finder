package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

func newPool(host string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 300 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", host) },
	}
}
