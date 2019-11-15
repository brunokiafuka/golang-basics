## Section 4 - Maps

#### Maps

> In go a `map` is a collection of `key` `value` pair, both the key and value are staticly typed. Similar to an `Object` in Javascript.

```go
package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
	}

	fmt.Println(colors)
}
```

##### Declaring Maps

```go

var colors map[string]string
colors := make(map[string]string)
colors := map[string]string{
	"red":   "#ff0000",
	"white": "#ffffff",
}
```

> To delete an element from a map we need to state the following `delete(<name of our map>, <key>)`

##### Iterating Over Maps

```go
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
```

#### Maps vs Struct

**Maps:**

- All keys must be of the same type.
- All values must be of the same type.
- Keys are indexed and can be iterated over.
- Used to represent a collection of related properties.
- Don't know all the keys at compile time.
- **Reference type**.

**Struct:**

- Values can be of any type.
- Keys don't support indexing.
- You need to know all the different fields at compile time.
- Used to represent a _thing_ with a lot of different properties.
- **Value type**.
