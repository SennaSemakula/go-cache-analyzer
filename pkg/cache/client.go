package cache

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	Cacher
}

type Cacher interface {
	NewClient() Cacher
	GetName() string
	NewLogger() *log.Logger
	Healthy(ctx context.Context) error
	GetItem(ctx *context.Context, key string) string
	SetItem(ctx context.Context, key string, val interface{})
}

type Redis struct {
	*redis.Client
	log  *log.Logger
	Addr string
	pass string
}

type Memcached struct {
	*memcache.Client
	log  *log.Logger
	Addr string
}

// func NewClient(c Cacher) *Client {
// 	return &Client{
// 		c.NewClient(),
// 		c.NewLogger(),
// 	}
// }

func (r Redis) NewClient() Cacher {
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.pass, // no password set
		DB:       0,      // use default DB
	})

	return Redis{
		rdb,
		r.NewLogger(),
		r.Addr,
		r.pass,
	}
}

func (r Redis) GetName() string {
	return "redis"
}

func (r Redis) NewLogger() *log.Logger {
	return log.New(os.Stderr, "redis: ", 0)
}

func (r Redis) Healthy(ctx context.Context) error {
	if _, err := r.Ping(ctx).Result(); err != nil {
		return err
	}
	return nil
}

func (r Redis) GetItem(ctx *context.Context, key string) string {
	val, err := r.getItem(ctx, key)
	if err != nil {
		r.log.Println(err)
	}
	return val
}

func (r Redis) SetItem(ctx context.Context, key string, val interface{}) {
	r.setItem(ctx, key, val)
}

func (r Redis) getItem(ctx *context.Context, key string) (string, error) {
	// r.log.Printf("get key: %s\n", key)
	if len(key) == 0 {
		return "", fmt.Errorf("get: key is empty")
	}
	val, err := r.Get(*ctx, key).Result()

	// TODO: refactor this
	if err != nil {
		switch err {
		case redis.Nil:
			return "", fmt.Errorf("%s key does not exist", key)
		default:
			return "", fmt.Errorf("get failed: %v", err)
		}
	}

	return val, nil
}

// writeItem sets a key in redis. By default the TTL is 30 mins
func (r *Redis) setItem(ctx context.Context, key string, value interface{}) error {
	// r.log.Printf("writing key %s\n", key)
	if _, err := r.Set(ctx, key, value, time.Minute*10).Result(); err != nil {
		return fmt.Errorf("writing key %s: %v", key, err)
	}
	return nil
}

// memcached
func (m Memcached) NewClient() Cacher {
	return &Memcached{
		memcache.New(m.Addr),
		m.NewLogger(),
		m.Addr,
	}
}

func (m Memcached) GetName() string {
	return "memcached"
}

func (m Memcached) NewLogger() *log.Logger {
	return log.New(os.Stderr, "memcached: ", 0)
}

func (m Memcached) Healthy(ctx context.Context) error {
	return m.Ping()
}

func (m Memcached) GetItem(ctx *context.Context, key string) string {
	// m.log.Printf("get key %s", key)
	item, err := m.getItem(key)
	if err != nil {
		m.log.Println(err)
		return ""
	}
	return item
}

func (c Memcached) SetItem(ctx context.Context, key string, value interface{}) {
	// c.log.Printf("writing to key %q", key)
	c.setItem(key, value)
}

func (m *Memcached) getItem(key string) (string, error) {
	item, err := m.Get(key)
	if err != nil {
		return "", fmt.Errorf("getting key %q %v", key, err)
	}

	return string(item.Value), nil
}

func (m *Memcached) setItem(key string, val interface{}) error {
	if err := m.Set(&memcache.Item{Key: key, Value: []byte(val.(string))}); err != nil {
		return fmt.Errorf("writing to key %q %v", key, err)
	}

	return nil
}
