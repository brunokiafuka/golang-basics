## Section 3 - Depper into go

#### Variables

`var <variable name> <data type> = "Ace of spades"`

Eg: `var cards string = "Ace of spades"`

- _var_ - indicates that we are about to create a new variable;
- _variable name_ - the name of our variable;
- _data type_ - indicates which data type our variable should support.

> **Note:** Go is a static typed language. We can also declare variables in the following manner: `cards := "Ace of spades"`, this way the Go compiler tries to figure out which data type is being used. `:=` is only used only when we are initializing a variable.

##### Go Basic types

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

#### Functions and return types

Defining functions:

```go
func <function name> () <return data type>{
    // ... function body
}

// The return data type indicates which type your return will be.
```

#### Arrays & Slice

> In go an `Array` is list that has fixed length whereas a `Slice` is an array that can grow or shrink. Ever element on an array or slice must be of the same data type.

```go
cards := []string{"A", "B", "C"} // Declaring a new slice
cards = append(cards, "D") // Adding element to a slice

// * The append function doesn't modify the current slice, instead it returns a new slice that then reassigned the housing variable.
```

##### Iterate over a slice (Loops)

```go
for i, card := range cards {
		fmt.Println(i, card)
}

// i - index of current element
// card - current "card" we are iterating over
// range cards - takes the slice of "cards" and loops over it
```

#### Custom Type Declarations

> A custom type is a type that borrows the behavior of the standard types

```go
type <your_custom_type> []string // in a separe file

cards := <your_custom_type>{"A", "B", "C"} // defining a slice using a custom type
```

#### Function reciever

> Function Receiver sets a method on variables that we create.

`func(t type) functionName() {}`

```go
type deck []string

func (d deck) print() {
    // d - is a ref to the var we created, by convection we use only a letter.
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// ... on main
func main() {
	cards := deck{"A", "B", "C"}
	cards = append(cards, "D")
	cards.print()
}
```

#### Type Conversion or Type Casting

```go
var index int8 = 15

var bigIndex int32

bigIndex = int32(index)

fmt.Println(bigIndex)
```

#### Testing with Go

> To make a test in Go we need to create a file ending in `<file_name>_test.go`. To run all tests in a package we run `go test`.
