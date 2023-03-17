# snowflake id
This is a Golang package for generating snowflake IDs.
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
	snowflake.NewSnowflake(1)
}

func main() {
	fmt.Printf("snowflake.CreateId(): %v\n", snowflake.Generate())
}

```
