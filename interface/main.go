// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an interface named speaker with a method named speak. Declare a struct
// named english that represents a person who speaks english and declare a struct named
// chinese for someone who speaks chinese. Implement the speaker interface for each
// struct using a value receiver and these literal strings "Hello World" and "你好世界".
// Declare a variable of type speaker and assign the address of a value of type english
// and call the method. Do it again for a value of type chinese.
//
// Add a new function named sayHello that accepts a value of type speaker.
// Implement that function to call the speak method on the interface value. Then create
// new values of each type and use the function.
package main

// Add imports.
import "fmt"

// Declare the speaker interface with a single method called speak.
type speaker interface {
	speak()
}

// Declare an empty struct type named english.
type english struct{}

// Declare a method named speak for the english type
// using a value receiver. "Hello World"
func (english) speak() {
	fmt.Println("Hello World")
}

// Declare an empty struct type named chinese.
type chinese struct{}

// Declare a method named speak for the chinese type
// using a value receiver. "你好世界"
func (chinese) speak() {
	fmt.Println("你好世界")
}

// sayHello accepts values of the speaker type.
func sayHello(s speaker) {

	// Call the speak method from the speaker parameter.
	s.speak()
}

func main() {

	// Declare a variable of the speaker type set to its zero value.
	var s speaker
	// Declare a variable of type english and assign it to
	// the speaker variable.
	var e english
	// Call the speak method from the speaker variable.
	s = e
	s.speak()
	// Declare a variable of type chinese and assign it to
	// the speaker variable.
	var c chinese
	// Call the speak method from the speaker variable.
	s = c
	s.speak()
	// Call the sayHello function with new values of each concrete type.
	sayHello(e)
	sayHello(c)
}
