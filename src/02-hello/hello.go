// All files containing a main method must be in the main package
package main

// fmt contains methods for formatting data in strings and io
// https://golang.org/pkg/fmt/
import "fmt"
// unused imports will not compile
//import "math"

// Note the lack of a return value or arguments on the main method. More on that later.
func main() {
    // A function with a capital letter?!?! What is that about? Yeah, we'll get to that later as well.
    fmt.Println("Hello,", "gopher!") // Arguments will be converted to strings and concatenated together with spaces in between
    fmt.Printf("Goodbye, %s.\n", "gopher") // Does what you would think.

    // Go's Printf() is really friendly.
    fmt.Printf("%f is %s", "fred")
    // Unused variables will not compile
    //var i int = 0
}

// BTW, where are all the semicolons? Semicolons are only used in Go when they are needed -- which is not at the end of a line of code

// Let's try some things at the command line
// go build
// go install
