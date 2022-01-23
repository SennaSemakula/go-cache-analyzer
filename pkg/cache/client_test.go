package cache

import (
	"context"
	"testing"
)

func BenchmarkRedisSet(b *testing.B) {
	redis := &Redis{Addr: "127.0.0.1:6380"}
	c := redis.NewClient()
	ctx := context.TODO()
	err := c.Healthy(ctx)
	if err != nil {
		b.Errorf("unable to connect to redis: %v", err)
	}

	for n := 0; n < b.N; n++ {
		c.SetItem(ctx, "fruit", "banana")
	}
}
func BenchmarkRedisGet(b *testing.B) {
	redis := &Redis{Addr: "127.0.0.1:6380"}
	c := redis.NewClient()
	ctx := context.TODO()
	err := c.Healthy(ctx)
	if err != nil {
		b.Errorf("unable to connect to redis: %v", err)
	}

	for n := 0; n < b.N; n++ {
		c.GetItem(&ctx, "fruits")
	}
}

func BenchmarkMemcacheSet(b *testing.B) {
	memcache := Memcached{Addr: "127.0.0.1:11211"}
	c := memcache.NewClient()
	ctx := context.TODO()
	err := c.Healthy(ctx)
	if err != nil {
		b.Errorf("unable to connect to memcached: %v", err)
	}

	for n := 0; n < b.N; n++ {
		c.SetItem(ctx, "fruit", "banana")
	}
}
func BenchmarkMemcacheGet(b *testing.B) {
	memcache := Memcached{Addr: "127.0.0.1:11211"}
	c := memcache.NewClient()
	ctx := context.TODO()
	err := c.Healthy(ctx)
	if err != nil {
		b.Errorf("unable to connect to memcached: %v", err)
	}

	for n := 0; n < b.N; n++ {
		c.GetItem(&ctx, "fruit")
	}
}
