Simple utility for creating maps where values have a expiration time

```go
// records should expire after 30 minutes
duration := time.Minute * 30

// check for expired records every minute
tckEvery := time.Minute * 1

ttlMap := New[string](duration, tckEvery)

ttlMap.Put("key", "some value")
ttlMap.Get("key") // some value

time.Sleep(time.Minute * 31)
ttlMap.Get("key") // ""
```