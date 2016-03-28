package main

import "fmt"

type myObj struct {
    s string
    i int
}

const PI = 3.14

// You can't use the short assignment operator outside of functions because everything here has to begin with a keyword.
//bad := true

// Information about primitive types can be found at https://tour.golang.org/basics/11. Note that complex numbers are a built in type. Cool!
func main() {
    zeroValues()
    typeInference()
    constants()
    pointersAndValues()
    typeConversion()
    extraCredit()
}

func zeroValues() {
    fmt.Println("Zero values")
    fmt.Println("===========")
    // types are specified after the variable (similar to scala)
    var i int // Numeric types default to zero
    var b bool // Boolean defaults to false
    var s string // String defaults to empty string
    var o myObj // Objects default to uninitialized values
    var p *myObj // Pointers default to nil

    fmt.Println("i is", i)
    fmt.Println("b is", b)
    fmt.Printf("s is '%s'\n", s)

    // We'll get back to structs later
    fmt.Printf("o is %+v\n", o)

    fmt.Println("p is", p)

    fmt.Println()
}

// There is an easier way
func typeInference() {
    fmt.Println("Type inference")
    fmt.Println("==============")

    s := "gabbagabbahey"
    b := true

    // Create a new variable
    o := myObj{}
    // Change it
    o = myObj{"answer", 42}
    // We can't reinitialize it
    //o := myObj{"even prime", 2}
    // And we can't change it's type
    //o = 1234

    fmt.Printf("s -> value = %s, type = %T\n", s, s)
    fmt.Printf("b -> value = %t, type = %T\n", b, b)
    fmt.Printf("o -> value = %+v, type = %T\n", o, o)

    fmt.Println()
}

func someFunc() float64 {
    return 1.71828
}

// Information about constants can be found at https://tour.golang.org/basics/15. In particular, note the value restrictions.
func constants() {
    fmt.Println("Constants")
    fmt.Println("=========")

    // constants can't use :=
    const s = "fred"

    // constants can't be structs
    //const o = myObj{}

    // They can be constant expressions
    const i = 10 >> 1

    // Functions are not constant expressions.
    //const f = someFunc()
    // And neither are variables
    //f := someFunc()
    //const constF = f

    fmt.Println("PI is", PI)
    fmt.Println("s is", s)
    fmt.Println("i is", i)

    fmt.Println()
}

func byVal(name string, i int) {
    i = i + i
    fmt.Println("in byVal:", name, "is", i)
}

func byRef(name string, i *int) {
    *i = *i + *i

    // Pointer math is not allowed, unlike c/c++. This means you can't change a pointer and look at other areas of memory.
    // This is a great safety feature
    //i = i + 1

    fmt.Println("in byRef:", name, "is", *i, "with address", i)
}

func pointersAndValues() {
    fmt.Println("Pointers and values")
    fmt.Println("===================")

    val := 5
    ref := 6

    fmt.Println("Before byVal, val is", val)
    byVal("val", val)
    fmt.Println("After byVal, val is", val)

    fmt.Println()

    fmt.Println("Before byRef, ref is", ref)
    byRef("ref", &ref) // not the address of operator
    fmt.Println("After byRef, ref is", ref)

    fmt.Println()

    // Go is friendlier with object pointers. Unlike c, where an object field can be dereferenced with the . operator, but
    // an object pointer field must be dereferenced with the -> operator.
    o := new(myObj)
    o.s = "Isn't that nice"
    o.i = 1
    fmt.Printf("o -> value = %+v, type = %T\n", o, o)

    fmt.Println()
}

func typeConversion() {
    fmt.Println("Type conversion")
    fmt.Println("===============")

    // Go requires explicit type conversion
    var i int = 8675309
    var f float64 = float64(i)
    var i2 int = int(f)
    // The following will not work.
    //var i3 uint = i2
    //fmt.Println("i3 is", i3) // just here so i3 is used if you uncomment

    fmt.Println("i is", i)
    fmt.Println("f is", f)
    fmt.Println("i2 is", i2)


    fmt.Println()
}

func extraCredit() {
    fmt.Println("Extra credit")
    fmt.Println("============")

    // arrays
    var arr1 [5]int
    arr1[2] = 3
    arr1[4] = 9
    arr2 := [3]int{1,2,3}
    arr3 := [3]int{2,4,6}
    // An array's size is part of its type
    // So you can do this
    arr2 = arr3
    // But not this
    //arr2 = arr1
    fmt.Printf("arr2 -> value = %v, type = %T - Notice the size in the type\n", arr2, arr2)

    // slices - backed by an array, but much more flexible
    // There is so much more to slices than we can cover here. Just know that they are similar to slices in Python
    // and you should use them instead of arrays in most situations.
    slice1 := []int{1,2,3,5,8,13}
    slice2 := []int{1,4,9}
    // Unlike arrays, a slice's size is not part of its type
    // So you can totally do this
    slice1 = slice2
    fmt.Printf("slice1 -> value = %v, type = %T - Notice the lack of size in the type\n", slice1, slice1)

    // slices and arrays are one-dimensional, but like c/c++ can be composed into multi-dimensional versions

    // maps - similar to maps in other languages
    // Like slices though, there is much more to maps than we are going to cover here
    var famousPairs map[string]string = make(map[string]string)
    famousPairs["Bonnie"] = "Clyde"
    famousPairs["Penn"] = "Teller"
    mapLiteral := map[string]string{"fred":"wilma", "barney":"betty"}
    fmt.Printf("famousPairs -> value = %v, type = %T\n", famousPairs, famousPairs)
    fmt.Printf("mapLiteral -> value = %v, type = %T\n", mapLiteral, mapLiteral)
    // Unavailable elements are returned as the zero value for the type
    fmt.Printf("famousPairs[\"Thelma\"] -> value = '%s'\n", famousPairs["Thelma"])
    // if you need to know presence of an element, you can use a multi-assignment. More on that when we talk about flow and functions.
    v, present := famousPairs["Ratchet"]
    fmt.Printf("famousPairs[\"Ratchet\"] -> value ='%s', present = %t\n", v, present)

    // Complex numbers
    c1 := 1 + 1i
    c2 := 2 + 3i
    c3 := c1 + c2
    c4 := c1 * c2
    fmt.Println("c3 is", c3)
    fmt.Println("c4 is", c4)

    fmt.Println()
}

// Compilation speed - ludicrous
