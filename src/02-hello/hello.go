// All files containing a main method must be in the main package.
package main

// fmt contains methods for formatting data in strings and io.
import "fmt"
// Unused imports will not compile.
//import "math"

// Note the lack of a return value or arguments for the main method. More on that later.
func main() {
    // The fmt package provides functions for formatting data with io streams, strings, and the console.
    fmt.Println("Hello,", "gopher!") // Arguments will be converted to strings and concatenated together with spaces in between.
    fmt.Printf("Goodbye, %s.\n", "gopher") // Does what you would think if you are familiar with printf in other languages.

    // Go's Printf() is really friendly.
    fmt.Printf("%f is %s", "fred")
    // Unused variables will not compile
    //var i int = 0
}

// BTW, where are all the semicolons? Semicolons are only used in Go when they are needed -- which is not at the end of a line of code.

// Let's try some things at the command line.
// go build
// go install
