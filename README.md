# Pebble (Not For Use)

GRPC based cache with a customizable backend. Default is badgerDB.

###Run 
```bash
./cache --db-dir="path/to/db" --port=":5555"
```

###Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/oriiolabs/pebble/api"
)

func main() {
	ctx := context.Background()

	c, err := api.NewClient(":4200", nil)
	if err != nil {
		panic(err)
	}

	if err := c.Set(ctx, "hello", []byte("world")); err != nil {
		panic(err)
	}

	value, err := c.Get(ctx, "hello")
	if err != nil {
		panic(err)
	}

	fmt.Printf("value: %s\n", string(value))
}
```
