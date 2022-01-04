package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
)

// Source Define some method who need to be impl for each source
type Source interface {
	init()
	refresh() []string
	getRandom() string
}

// Redis Source for redis
type Redis struct {
	rdb *redis.Client
	*redis.Options
}

func (r *Redis) init() error {
	if r.rdb == nil {
		if r.Options == nil {
			panic("Redis Options is null!!")
		}
		r.rdb = redis.NewClient(r.Options)
	}
	ping := r.rdb.Ping(ctx)
	return ping.Err()
}

func (r *Redis) refresh() []string {
	servers, err := r.rdb.LRange(ctx, "improvised:servers", 0, -1).Result()
	checkError(err)
	return servers
}

func (r *Redis) getRandom() *string {
	servers := r.refresh()
	if len(servers) == 0 {
		fmt.Println("There is no server available")
		return nil
	}
	return &servers[rand.Intn(len(servers))]
}
