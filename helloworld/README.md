## Section 2 - A simple Start

### Five basic questions

#### 🔥 How do we run the project in go?

> We run a go project running the following command `go run <File-name>.go`

#### Go CLI - basic commands

- `go build` - compiles a bunch of go source code files.
- `go run` - compiles and executes one or two files.
- `go fmt` - formats all the code in each in each file in the current directory.
- `go install` - compiles and installs a package.
- `go get` - downloads the raw source code of someone else's package.
- `go test`- runs ant test associated with the current project.

> `go install` and `go get` allow us to install external dependencies.

#### 📦 What does `package main` mean?

> A package indicates a project or workplace.

**Types of packages**

- Executable: generates a file that we can run;
- Reusable: code used as _helpers_. Goof place to put reusable logic.

> The name of your package is what defines if you are either creating an executable package or a reusable.

#### 📥 What does `import "fmt"` mean?

> Import means that we are getting access to files within an external package ('fmt' as example).

The `fmt` packages is a standard package of goLang. You an also import your own reusable packages.

[Go 📥 Standard packages](https://golang.org/pkg/)

#### 🏷 What's that `func` thing?

> `func` is a way to declare a function in go. Refer to previous programming languages

#### 📚 How is the main.go file organize?

```go
// Package Declaratio
package main

// Import packages we'll need
import "fmt"

// Declare functions, tell go to do things
func main() {
	fmt.Println("Hi there")
}

```

[Download ⬇️ Class diagrams](https://github.com/StephenGrider/GoCasts/tree/master/diagrams)
