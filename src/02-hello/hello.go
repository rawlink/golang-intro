// All files containing a main method must be in the main package
package main

// fmt contains methods for formatting data in strings and io
// https://golang.org/pkg/fmt/
import "fmt"

// Not that there is no return value or arguments to our main method. We will get to that later.
func main() {
    // A function with a capital letter?!?! What is that about? Yeah, we'll get to that later as well.
    fmt.Println("Hello,", "gopher!") // Arguments will be converted to strings and concatenated together with spaces in between
    fmt.Printf("Goodbye, %s.\n", "gopher") // Does what you would think.
    // BTW, where are all the semicolons?
}

// Let's try some things at the command line
// go build
// go install
