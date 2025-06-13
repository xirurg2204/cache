# GO in memory cache
================================

Go in memory cache used to store information in RAM and quickly access it.

See it in action:

## Example

```go
package main

import (
	"fmt"
	"github.com/xirurg2204/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)
}

```
