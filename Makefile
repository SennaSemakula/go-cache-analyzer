.PHONY: redis.start redis.teardown memcache.start start benchmark

PORT:=6380

start:
	@go run .

benchmark:
	@go test -bench . ./... -benchtime=1000000x

redis.start:
	@echo "starting redis server on port ${{PORT}}"
	@redis-server --port ${PORT} &

redis.teardown:
	@echo "tearing down redis server"
	@redis-cli -p ${PORT} shutdown

memcache.start:
	@echo "starting memcached server"
	@memcached start &

memcache.teardown:
	@echo "stopping memcached server"
	@lsof -i TCP | grep memcach | awk '{ print $2 }' | xargs kill -9


