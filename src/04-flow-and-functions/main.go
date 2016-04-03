package main

// This is a factored import statement
import (
    "fmt"
    "os"
    m "math" // Renamed the math import to m. You should really only do this with external dependencies when there is a naming conflict.
    "errors"
    "math/rand" // A sub-import from math. The import will be available as the last part of the path "rand".
    "time"
)

func main() {
    ifStatements()
    loops()
    switchStatements()
    firstClassFunctions()
    deferredFunctions()
    closures()
    variadicFunctions()

    fmt.Println("Exiting main")
    fmt.Println("=====================")

    // We renamed math to m in the import.
    fmt.Println("Using my horribly renamed import, the square root of 2 is", m.Sqrt(2.0))

    // Note the exit status in the run console.
    os.Exit(1)
}

func multiReturn() (int,error) { // If you have multiple return values the return signature should be surrounded by parentheses.
    // Here we have an example of a multi-return. A return signature of (<type>, error) is very common in Go.
    val := 42
    err := errors.New("Bad stuff happened")
    //err = nil
    return val, err
}

func ifStatements() {
    fmt.Println("If statements")
    fmt.Println("=============")

    // The braces *are* required Bryce. Really, it's not just me. Go thinks so too.
    if true {
        fmt.Println("Always truthful.")
    }
    // This won't compile.
    //if true
    //    fmt.Println("Always truthful.")

    // The else if and else are as expected.
    num := 3
    if num < 0 {
        fmt.Println("Negative, sir")
    } else if num == 0 {
        fmt.Println("A whole lot of nothing")
    } else if num == 1 {
        fmt.Println("One")
    } else if num == 2 {
        fmt.Println("Couple")
    } else if num == 3 {
        fmt.Println("Few")
    } else {
        fmt.Println("Several or more")
    }

    // A better if. Why?...
    if val, err := multiReturn() ; err == nil {
        fmt.Println("Success! The value is", val)
    } else {
        fmt.Println("Error!", err)
    }
    // The values in the assignment block of the if are out of scope here. No more c/c++/c#/java/etc antipattern here.
    //fmt.Println(val)

    fmt.Println()
}

func loops() {
    fmt.Println("Loops")
    fmt.Println("=========")

    // Plain old for. The semicolons are back.
    fmt.Print("i is")
    for i := 0 ; i < 4 ; i++ {
        fmt.Print(" ", i)
    }
    fmt.Println()

    // The init and post are optional.
    j := 0
    fmt.Print("j is")
    for ; j < 2 ; {
        fmt.Print(" ", j)
        j++
    }
    fmt.Println()

    // But if you drop the semicolons, you have a go while loop.
    k := 0
    fmt.Print("k is")
    for k < 4 {
        fmt.Print(" ", k)
        k++
    }
    fmt.Println()

    // For... ever.
    rand.Seed(time.Now().UnixNano())
    fmt.Print("You might be stuck here forever...")
    for {
        number := rand.Intn(1000)
        fmt.Print(" ", number)
        if number == 666 {
            break
        }
    }
    fmt.Println("\n... or not.")

    // You can use the range keyword in conjunction with for loops. This is a for each loop.
    // There are several libraries for more advanced handling of application arguments.
    fmt.Printf("OS arguments type = %T\n", os.Args)
    // Example of a for loop using a range. Note the multi-return from the range.
    for index,val := range os.Args {
        fmt.Println("Argument", index, "value =", val)
    }
    // Lets ignore the index. You can use _ to swallow unneeded variables.
    //for _,val := range os.Args {
    //    fmt.Println("Argument value =", val)
    //}
    // Lets ignore the value.
    //fmt.Print("There are")
    //for index := range os.Args {
    //    fmt.Print(" ", index)
    //}
    //fmt.Println(" values")

    // Range works with maps too.
    squares := map[int]int{1:1, 2:4, 3:9, 4:16}
    for key,val := range squares {
        fmt.Println(key, "squared is", val)
    }

    fmt.Println()
}

func getFirstName() string {
    return "wilma"
    //return "bambam"
    //return "rockhead"
}

func switchStatements() {
    fmt.Println("Switch statements")
    fmt.Println("=================")

    // Like an if statement, there is an optional assignment block.
    // And you don't need that stupid "break" that everyone else does. If you need it there is fallthrough.
    switch firstName := getFirstName() ; firstName {
    case "bambam":
        fmt.Println("Adopted")
        fallthrough
    case "betty":
        fallthrough
    case "barney":
        fmt.Println("Rubbles family")
    case "fred","wilma","pebbles","dino":
        // If you only need to handle multiple cases as one, you can forget fallthrough altogether.
        fmt.Println("Flintstones family")
    default:
        fmt.Println("Nobody important")
    }

    // Remember that long if statement... a conditionless switch might be a better way to go. Nice and clean.
    num := 100
    switch {
    case num < 0:
        fmt.Println("Negative, sir")
    case num == 0:
        fmt.Println("A whole lot of nothing")
    case num == 1:
        fmt.Println("One")
    case num == 2 :
        fmt.Println("Couple")
    case num == 3 :
        fmt.Println("Few")
    default:
        fmt.Println("Several or more")
    }

    fmt.Println()
}

func calc(a int, b int, operation func(int, int) int) int {
    return operation(a,b)
}

func firstClassFunctions() {
    fmt.Println("First class functions")
    fmt.Println("=====================")

    add := func (a int, b int) int { return a + b }
    multiply := func (a int, b int) int { return a * b }

    fmt.Println("calc(3, 5, add) is", calc(3, 5, add))
    fmt.Println("calc(4, 6, multiply) is", calc(4, 6, multiply))

    fmt.Println()
}

func deferral(i int) {
    fmt.Println("Deferral ", i)
}

func measureFunc(functionName string, start time.Time) {
    duration := time.Since(start)
    fmt.Println(functionName, "took", duration, "to run")
}

func anotherCoolDeferredExample() {
    defer measureFunc("anotherCoolDeferredExample", time.Now())
    fmt.Println("Doing work in anotherCoolDeferredExample")
}

func deferredFunctions() {
    fmt.Println("Deferred functions")
    fmt.Println("==================")

    // This is taking care of making sure we print an empty line after we leave the function and all other defers have run.
    // Why is it our first defer?
    defer fmt.Println()

    // Deferred function parameters are resolved at time of deferment. Deferred function calls are put onto a stack resulting
    // in them being called in reverse order.
    fmt.Println("Before for loop")
    for i := 0 ; i < 5 ; i++ {
        defer deferral(i)
    }
    fmt.Println("After for loop")

    /* Defer is a great mechanism to do things you would normally do with finally in Java. For example:

        f, err := os.Open("/path/to/file.txt")
        if err != nil {
            // log and return an error here
        }
        defer f.Close()
        f.Read(...)

    */

    anotherCoolDeferredExample()

    fmt.Println("Exiting deferredFunctions")
}

func getFibber() func() int {
    a, b := 0, 1
    return func() int {
        a, b = b, a+b // It's worth noting that these assignments happen "simultaneously"
        // The following would return an altogether different result. Do you know what the result would be?
        //a = b
        //b = a+b
        return a
    }
}

func closures() {
    fmt.Println("Closures")
    fmt.Println("========")

    // What an elegant functional fibonacci iterator.
    fib := getFibber()
    for i := 1 ; i <= 10 ; i++ {
        fmt.Println("Fibonnaci", i, "is", fib())
    }

    fmt.Println()
}

func sumOf(what string, nums ...int) {
    sum := 0
    for _, val := range nums {
        sum += val
    }
    fmt.Println("The sum of", what, "is", sum)
}

func variadicFunctions() {
    fmt.Println("Variadic functions")
    fmt.Println("==================")

    // varargs functions are as expected
    sumOf("all fears", 15, 34, 53, 22, 91)

    fmt.Println()
}
