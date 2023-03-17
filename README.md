# snowflake id
This is a golang package that generates snowflake id
```
go get github.com/404669366/snowflake@latest
```
```golang
package main

import (
	"fmt"
	"github.com/404669366/snowflake"
)

func init() {
	snowflake.Init(1)
}

func main() {
	fmt.Printf("snowflake.CreateId(): %v\n", snowflake.CreateId())
}

```
