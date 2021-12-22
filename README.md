# functions to go

Handy collection-functions for the `go` language. 

Functions for `map`, `reduce`, `filter`, `associate`, `associateby`, `contains`.



## getting started

    go get github.com/nilsmagnus/functions

```go
package main

import (
	"fmt"
	"github.com/nilsmagnus/functions"
	"log"
)

func main() {
	items := []int{1, 2, 99}

	mappedItems := functions.Map(items, func(a int) string { return fmt.Sprintf("mapped %d", a) })

	log.Printf("Mapped items: %v", mappedItems)
}

```

Run this example on the [playground](https://go.dev/play/p/tpcenABL62q?v=gotip)


Have a look at [functions_test.go](functions_test.go) for usage.


### Map-function
```go

numbers := []int{1,2,3,199}

mapd := Map(numbers, func(a int) string { return fmt.Sprintf("%d", a)}) 

// mapd => []string{"1", "2", "3", "199"}
```

### AssociateBy-function 
```go
numbers := []int{1,2,3,199}

mapd := AssociateBy(numbers, func(t int) string {
    return fmt.Sprintf("%d", t)
}) 

// => map[string]int {"1":1, "2":2, "3":3, "199": 199}
```
