# Go Env

This package to make it easy to work with env

**Example usage**
```go
package main

import (
    "fmt"
    "github.com/frstpw/go-env"
    "os"
)

func main() {
    os.Setenv("ENV_EXAMPLE", "123123")
    defer os.Unsetenv("ENV_EXAMPLE")

    fmt.Println(env.Get("ENV_EXAMPLE").Int(133))
}
```

Output: **123123**

```go
env.Get("ENV_EXAMPLE").Int(133)
env.Get("ENV_EXAMPLE").String("default value")
env.Get("ENV_EXAMPLE").Bool(false)
env.Get("ENV_EXAMPLE").Float32(3.1415927)
env.Get("ENV_EXAMPLE").Float64(3.141592653589793)
```