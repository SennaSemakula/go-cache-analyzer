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
	// TODO: create seperate contexts
	ctx := context.Background()
	memcached := cache.Memcached{Addr: fmt.Sprintf("127.0.0.1:%v", memcachedPort)}
	for _, v := range []cache.Cacher{redis, memcached} {
		if err := writeCache(ctx, v); err != nil {
			log.Fatal(err)
		}
	}
}

func writeCache(ctx context.Context, c cache.Cacher) error {
	client := c.NewClient()
	if err := client.Healthy(ctx); err != nil {
		return err
	}
	// Ping cache
	if err := client.Healthy(ctx); err != nil {
		return err
	}

	val := client.GetItem(&ctx, "fsd")
	fmt.Println(val)
	client.SetItem(ctx, "fruit", "peaches")
	fmt.Println(client.GetItem(&ctx, "fruit"))

	return nil
}
