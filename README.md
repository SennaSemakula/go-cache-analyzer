# go-cache-analyzer
Retrieve `GET/SET` benchmarking for redis and memcached. Both caches write the same amount of bytes and data type. By default the benchmark runs 10^6 iterations. 

Due to running both instances of redis and memcached on my laptop, these results may be inconsistent. It's advised to run cache instance on isolated environment for realistic results.

## Dependencies
```memcached 1.6.13```
<br>
```redis-server 6.2.6```
<br>
```go version go1.17 darwin/arm64```

Memcached only supports strings hence why I've only tested using both of those. 

## Running
Once you have both redis and memcached instances running, issue the following command:
```make benchmark```

## Results 
![Results](assets/analyse.png)

As you can see writes to Memcache seem to be faster than redis when it comes to comes to writing a string. Redis trumps Memcached with GETs.
