# Rate Limiter

Redis is a good service to implement rate limiter, it's fast and east to support distributed system.

## 

[GET](https://redis.io/commands/get)
Get the value of key.

[INCR](https://redis.io/commands/incr)
Increments the number stored at key by one.

[EXPIRE](https://redis.io/commands/expire)
Set a timeout on key.

[MULTI](https://redis.io/commands/multi/)
Marks the start of a transaction block.

[EXEC](https://redis.io/commands/exec/)
Executes all previously queued commands in a transaction and restores the connection state to normal.

[ZCOUNT](https://redis.io/commands/zcount/)
Returns the number of elements in the sorted set at key with a score between min and max.

[RPUSH](https://redis.io/commands/rpush/)
Insert all the specified values at the tail of the list stored at key. If key does not exist, it is created as empty list before performing the push operation. When key holds a value that is not a list, an error is returned.

| Redis key <br>(APIKEY, IPs, UserId ...etc) | 191.168.2.10 | 192.168.2.15 |
|--------------------------------------------|--------------|--------------|
| Value                                      | 8            | 16           |
| Expires at                                 | 12:01| 12:05|


 ## Basic flow

1. GET: [Redis key]:[current number]
2. if current number exceed RATE-LIMITER, return error (drop request)
3. MULTI:
   1. INCR [Redis key]:[current number]
   2. EXPIRE [Redis key]:[current number] 59 (1 min)
4. EXEC

2 key points
1. INCR on non-exist key will always be 1.
2. EXPIRE is inside a MULTI along with the INCR (means this is a atomic operation)


## Algorithm

### Fixed Window
#### Pros
1. easy to implement
#### Cons
1. A burst of requests at the end of the window causes server handling more requests than the limit
### Sliding Logs
#### Pros
1. Overcome fixed window's cons
#### Cons
1. It is not memory-efficient
2. It is very expensive because we count the user’s last window requests in each request.
### Leaking Bucket
#### Pros
1. Overcomes the cons of the fixed window by not imposing a fixed window limit and thus unaffected by a burst of requests at the end of the window.
2. Overcomes the cons of sliding logs by not storing all the requests(only the requests limited to queue size) and thus memory efficient.
#### Cons
1. Bursts of requests can fill up the queue with old requests and most recent requests are slowed from being processed and thus gives no guarantee that requests are processed in a fixed amount of time.
2. This algorithm causes traffic shaping(handling requests at a constant rate, which prevents server overload, a plus point), which slows user’s requests and thus affecting your application.
### Sliding Window
#### Pros
1. Overcomes the cons of the fixed window by not imposing a fixed window limit and thus unaffected by a burst of requests at the end of the window.
2. Overcomes the cons of sliding logs by not storing all the requests and avoiding counting for every request and thus memory and performance efficient.
3. Overcomes the cons of leaky bucket starvation problem by not slowing requests, not traffic shaping.
#### Cons

### Token Bucket
#### Pros
1. Overcomes all the above algorithms cons, no fixed window limit, memory and performance efficient, no traffic shaping.
2. No need for background code to check and delete expired keys.
#### Cons