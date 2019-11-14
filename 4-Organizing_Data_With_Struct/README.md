## Section 3 - Organizing Data With Structs

#### Struct

> `Struct` is short for structure, it is a data structure in GoLang that defines a collection of properties that are related together serving a common propose. [More](https://www.golangprograms.com/go-language/struct.html)

##### Defining a struct

Whenever we are making a `Struct` we have to first define all the different properties a struct might have. Then we create a new value that will receive the defined struct

```go
type person struct {
    firstName string
    lastName string
    age int
}
```

##### Declaring a struct

```go

type person struct {
	firstName string
    lastName  string
    age int
}

func main() {
    // Bad ❌
    bruno := person{"Bruno", "Kiafuka", 23}

    // Good ✅
    bruno := person{firstName: "Bruno", lastName: "Kiafuka", age: 23}

    // other declaration option
    var bruno person
	fmt.Printf("%+v", bruno)

}
```

##### Updating a struct

> When declaring a struct using `var bruno person` go assigns a zero value to each property of your struct.

```go
func main() {
	var bruno person

	bruno.firstName = "Bruno"
	bruno.lastName = "Kiafuka"
	bruno.age = 23
	fmt.Printf("%+v", bruno)
}
```

##### Embedding structs

```go
type contactInfo struct {
	email string
	zip   string
}

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

func main() {

	bruno := person{
		firstName: "Bruno",
		lastName:  "Kiafuka",
		age:       23,
		contact: contactInfo{
			email: "bruno@email.com",
			zip:   "1021",
		},
	}

	fmt.Printf("%+v", bruno.contact)
}
```

##### Structs with Receiver Functions

```go
func main() {

	bruno := person{
		firstName: "Bruno",
		lastName:  "Kiafuka",
		age:       23,
		contact: contactInfo{
			email: "bruno@email.com",
			zip:   "1021",
		},
	}

    bruno.updateName("Sebas")
	bruno.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
```
