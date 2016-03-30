package main

/*
    Person
 */
type Person struct {
    Phrase string
    Thought string
}

// The new (p Person) syntax before the method name is called a receiver. This is what makes this a method instead
// of a function.
func (p Person) Speak() string {
    return p.Phrase
}

func (p Person) Reason() string {
    return p.Thought
}

func (p Person) Sing() string {
    return "Lalala"
}
/*
    Dog
 */
type Dog struct {
}

func (d Dog) Speak() string {
    return "Bark!"
}

/*
    MyFloat
 */
type MyFloat float64

func (f MyFloat) Square() MyFloat {
    return f * f // Cool, it still acts like a float
}

/*
    Interfaces
 */
type Speaker interface {
    Speak() string
}

type Sentient interface {
    Speak() string
    Reason() string
}
