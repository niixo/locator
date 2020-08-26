# Locator
locator package

## Install

```
go get github.com/niixo/locator
```

## Example

```go
package main

import (
	"fmt"

	"github.com/niixo/locator"
)

func main() {
	ql := locator.NewQuery("testdata")
	q, err := ql.Load("select_posts_by_id")
	if err != nil {
		panic(err)
	}
	fmt.Printf(q) // confirm sql
}
```
