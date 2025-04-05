There are some functions for eg for starters lets say `print` fn which comes from a go package(a built-in package) called fmt or format.
    We have to import any packages from whic we're using the functionality. for eg- for `print` we import the "fmt" package like this:
        ` import "fmt" ` and then calling the function followed by the name of the package- `fmt.Print`

## Go packages
- Go programs are organized into packages.
- Go's standard library provides diffrent core packages for us to use.
- "fmt"(mentioned above) is one of these that can be used by importing.

A `Package` is a collection of source files.
How do we know that in which package you have that functionaity available?
-> We actually have to go to Golang docs or google search it and as we do and learn we'll automatically start recognizing or remembering the most common packagesand their functionalities at the very least

## Executing a Go application
`go run <name of the file to execute>`