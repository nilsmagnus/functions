# functions

Handy collection-functions for the `go` language. 

Functions for `map`, `reduce`, `filter`, `associate`, `associateby`, `contains`.


Have a look at [functions_test.go] for usage

## getting started

    go get github.com/nilsmagnus/functions


### Map-function
```language=go

numbers := int[]{1,2,3,199}

mapd := map(numbers, func(a int) string { return fmt.Sprintf("%d", a)} // => ["1", "2", "3", "199"]
```

### AssociateBy-function 
```
numbers := int[]{1,2,3,199}

mapd := AssociateBy(numbers, func(t int) string {
    return fmt.Sprintf("%d", t)
}) 

// => map[string]int {"1":1, "2":2, "3":3, "199": 199}

```
