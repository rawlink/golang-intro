package main

import (
    "fmt"
)

func main() {
    // We are only scratching the surface here. Refer to Effective Go (on the golang.org site) for more in depth coverage.
    methods()
    interfaces()
    typeAssertions()
}


func methods() {
    fmt.Println("Methods")
    fmt.Println("=======")

    swedishChef := Person{
        Phrase: "Bork bork bork!",
    }

    fmt.Printf("swedishChef says '%s'\n", swedishChef.Speak())

    myFloat := MyFloat(16.0)
    myFloat.Square()

    fmt.Printf("myFloat -> value = %f, squared = %f\n", myFloat, myFloat.Square())

    fmt.Println()
}

func interfaces() {
    fmt.Println("Interfaces")
    fmt.Println("==========")

    person := Person{
        Phrase: "Hello",
        Thought: "Cogito ergo sum",
    }
    // You don't have to specify the field names, but I prefer to specify them
    /*
    Person{
        "Hello",
        "Cogito ergo sum",
    }
    */
    dog := Dog{}

    var sentient Sentient
    // Hold on... I don't remember Person implementing the sentient interface.
    sentient = person
    fmt.Printf("sentient says '%s' and thinks '%s'\n", sentient.Speak(), sentient.Reason())
    // So can we just assign dog to sentient? Nope.
    //sentient = dog

    var speaker Speaker
    // Can we assign compatible interfaces? Why yes, yes we can.
    speaker = sentient
    fmt.Printf("sentient speaker says '%s'\n", speaker.Speak())
    speaker = dog
    fmt.Printf("non-sentient speaker says '%s'\n", speaker.Speak())
    speaker.Speak()

    fmt.Println()
}

func typeAssertions() {
    fmt.Println("Type assertions")
    fmt.Println("===============")

    var i interface{} = "All objects match the empty interface"

    // If the type assertion fails, the value will be the zero value for the type
    f, ok := i.(float64)
    fmt.Printf("Type assertion results for float64 -> f = %f, ok = %t\n", f, ok)
    s, ok := i.(string)
    fmt.Printf("Type assertion results for string -> s = %s, ok = %t\n", s, ok)
    // There is a single result version of a type assertion, but it will panic if it fails.
    //fmt.Println("bool assertion result is", i.(bool))

    // We can pair a type assertion with a switch to provide instanceof-like functionality
    switch i.(type) {
    case int:
        fmt.Println("It's an int")
    case string:
        fmt.Println("It's a string")
    default:
        fmt.Println("Hell if I know")
    }
    fmt.Println()
}
