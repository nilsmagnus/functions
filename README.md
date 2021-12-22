# functions to go

Handy collection-functions for the `go` language. 

Functions for `map`, `reduce`, `filter`, `associate`, `associateby`, `contains`.



## getting started

    go get github.com/nilsmagnus/functions


Have a look at [functions_test.go](functions_test.go) for usage.


### Map-function
```go

numbers := []int{1,2,3,199}

mapd := Map(numbers, func(a int) string { return fmt.Sprintf("%d", a)}) 

// mapd => ["1", "2", "3", "199"] // => ["1", "2", "3", "199"]
```

### AssociateBy-function 
```go
numbers := []int{1,2,3,199}

mapd := AssociateBy(numbers, func(t int) string {
    return fmt.Sprintf("%d", t)
}) 

// => map[string]int {"1":1, "2":2, "3":3, "199": 199}
```
