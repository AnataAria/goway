package redis

import "time"

type RedisConfig struct {
	Addr     string
	Password string
	DB       int

	PoolSize     int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
