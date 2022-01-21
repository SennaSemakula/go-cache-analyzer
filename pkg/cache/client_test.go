package cache

import (
	"context"
	"testing"
)

func BenchmarkRedisSet(b *testing.B) {
	c := &Redis{Addr: "127.0.0.1:6380"}
	c = c.NewClient()
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
	c := &Redis{Addr: "127.0.0.1:6380"}
	c = c.NewClient()
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
	c := Memcached{Addr: "127.0.0.1:11211"}
	c = *c.NewClient()
	err := c.Healthy()
	if err != nil {
		b.Errorf("unable to connect to memcached: %v", err)
	}

	for n := 0; n < b.N; n++ {
		c.SetItem("fruit", "banana")
	}
}
func BenchmarkMemcacheGet(b *testing.B) {
	c := Memcached{Addr: "127.0.0.1:11211"}
	c = *c.NewClient()
	err := c.Healthy()
	if err != nil {
		b.Errorf("unable to connect to memcached: %v", err)
	}

	for n := 0; n < b.N; n++ {
		c.GetItem("fruit")
	}
}
