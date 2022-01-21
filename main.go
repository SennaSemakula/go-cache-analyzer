package main

import (
	"context"
	"fmt"
	"github.com/SennaSemakula/redis-cache/pkg/cache"
	"log"
)

const (
	redisPort     = 6380
	memcachedPort = 11211
)

func main() {
	connectMemcached()
}

func connectCache(ctx context.Context, c cache.Cacher, addr string) {
	c.NewClient()
	if err := c.Healthy(ctx); err != nil {
		log.Fatal(err)
	}
}

func redisOperations(c *cache.Redis) {
	redis := cache.Redis{Addr: "127.0.0.1:6380"}
	redis = *redis.NewClient()
	ctx := context.Background()

	// Ping redis
	if err := c.Healthy(ctx); err != nil {
		log.Fatalf("unable to initiate connection with redis: %v. is redis connection open?", err)
	}

	val := c.GetItem(&ctx, "fsd")
	fmt.Println(val)
	c.SetItem(ctx, "fruit", "peaches")
	fmt.Println(c.GetItem(&ctx, "fruit"))
}

func connectMemcached() {
	memcached := cache.Memcached{Addr: fmt.Sprintf("127.0.0.1:%v", memcachedPort)}
	memcached = *memcached.NewClient()

	// ctx := context.Background()
	c := memcached.NewClient()

	// Ping memcached server
	if err := c.Healthy(); err != nil {
		log.Fatalf("unable to initiate connection with memached: %v", err)
	}
	c.SetItem("bob", "jones")
	fmt.Println(c.GetItem("bob"))

}
