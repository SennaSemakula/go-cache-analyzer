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
	redis := cache.Redis{Addr: "127.0.0.1:6380"}
	ctx := context.Background()
	writeCache(ctx, redis)

	memcacheCtx := context.Background()
	memcached := cache.Memcached{Addr: fmt.Sprintf("127.0.0.1:%v", memcachedPort)}
	writeCache(memcacheCtx, memcached)
}

func writeCache(ctx context.Context, c cache.Cacher) {
	client := c.NewClient()
	if err := client.Healthy(ctx); err != nil {
		log.Fatal(err)
	}
	cacheName := client.GetName()
	// Ping cache
	if err := client.Healthy(ctx); err != nil {
		log.Fatalf("unable to initiate connection with %s: %v. is %s connection open?", cacheName, err, cacheName)
	}

	val := client.GetItem(&ctx, "fsd")
	fmt.Println(val)
	client.SetItem(ctx, "fruit", "peaches")
	fmt.Println(client.GetItem(&ctx, "fruit"))
}
