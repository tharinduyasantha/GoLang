package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
)

// Variable Assignment and Data Types
func variableExamples() {
	fmt.Println("Variable Assignment and Data Types:")

	var a int = 10
	var b float64 = 20.5
	c := "Hello, Go"
	var d bool = true
	var e complex128 = complex(1, 2) // Complex number

	fmt.Printf("a: %d, b: %.2f, c: %s, d: %t, e: %v\n", a, b, c, d, e)
}

// Constants and Iota
func constantsExamples() {
	fmt.Println("\nConstants and Iota:")

	const Pi = 3.14
	fmt.Println("Pi:", Pi)

	const (
		A = iota
		B
		C
	)
	fmt.Println("Iota constants:", A, B, C)
}

// Slices and Maps
func slicesAndMapsExamples() {
	fmt.Println("\nSlices and Maps:")

	// Slices
	s := []int{1, 2, 3, 4, 5}
	s = append(s, 6)
	fmt.Println("Slice:", s)

	// Maps
	m := map[string]int{"one": 1, "two": 2}
	m["three"] = 3
	fmt.Println("Map:", m)
}

// Control Structures
func controlStructures() {
	fmt.Println("\nControl Structures:")

	// If-Else
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	// For Loop
	fmt.Println("For Loop:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For Range
	fmt.Println("For Range:")
	array := []int{1, 2, 3, 4, 5}
	for i, v := range array {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// Switch
	day := "Tuesday"
	fmt.Println("Switch Statement:")
	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	default:
		fmt.Println("Today is neither Monday nor Tuesday")
	}
}

// Functions
func functionExamples() {
	fmt.Println("\nFunctions:")

	// Simple Function
	sum := add(2, 3)
	fmt.Printf("add(2, 3): %d\n", sum)

	// Function with Multiple Return Values
	quotient, err := divide(10.0, 2.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("divide(10.0, 2.0): %.2f\n", quotient)
	}

	// Named Return Values
	area, perimeter := rectangleDimensions(5.0, 3.0)
	fmt.Printf("rectangleDimensions(5.0, 3.0): Area: %.2f, Perimeter: %.2f\n", area, perimeter)

	// Variadic Function
	total := sumVariadic(1, 2, 3, 4, 5)
	fmt.Printf("sumVariadic(1, 2, 3, 4, 5): %d\n", total)

	// Anonymous Function
	message := func(name string) string {
		return "Hello, " + name
	}("Go")
	fmt.Println(message)

	// Closure
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	fmt.Printf("Counter: %d, %d\n", counter(), counter())

	// Method Function
	rect := Rectangle{length: 5.0, width: 3.0}
	fmt.Printf("Rectangle Area: %.2f\n", rect.area())
}

// Simple Function
func add(a int, b int) int {
	return a + b
}

// Function with Multiple Return Values
func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Function with Named Return Values
func rectangleDimensions(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

// Variadic Function
func sumVariadic(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Method Function
type Rectangle struct {
	length, width float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

// Pointers
func pointerExamples() {
	fmt.Println("\nPointers:")

	var a int = 10
	var p *int = &a // Pointer to a

	fmt.Printf("Value of a: %d\n", a)
	fmt.Printf("Address of a: %p\n", &a)
	fmt.Printf("Value of p: %p\n", p)
	fmt.Printf("Value pointed by p: %d\n", *p)
}

// Interfaces
func interfaceExamples() {
	fmt.Println("\nInterfaces:")

	var s Shape
	s = Rectangle{5, 3}
	fmt.Printf("Rectangle Area: %.2f\n", s.area())

	s = Circle{2.5}
	fmt.Printf("Circle Area: %.2f\n", s.area())
}

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Error Handling
func errorHandling() {
	fmt.Println("\nError Handling:")

	result, err := sqrt(-1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("cannot compute square root of a negative number")
	}
	return math.Sqrt(x), nil
}

// Goroutines and Channels
func concurrencyExamples() {
	fmt.Println("\nGoroutines and Channels:")

	messages := make(chan string)

	go func() {
		messages <- "Hello from Goroutine"
	}()

	msg := <-messages
	fmt.Println(msg)
}

// Deferred Function Calls
func deferredFunctionCalls() {
	fmt.Println("\nDeferred Function Calls:")

	defer fmt.Println("Deferred call 1")
	defer fmt.Println("Deferred call 2")
	fmt.Println("Normal call")
}

// Panic and Recover
func panicAndRecover() {
	fmt.Println("\nPanic and Recover:")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("This is a panic")
}

// Type Assertions
func typeAssertions() {
	fmt.Println("\nType Assertions:")

	var i interface{} = "Hello"
	s, ok := i.(string)
	if ok {
		fmt.Println("Type assertion succeeded:", s)
	} else {
		fmt.Println("Type assertion failed")
	}
}

// Embedding Interfaces
func embeddingInterfaces() {
	fmt.Println("\nEmbedding Interfaces:")

	var d Driver
	d = PersonDriver{}
	d.Drive()
	d.Work()
}

type Worker interface {
	Work()
}

type Driver interface {
	Worker
	Drive()
}

type PersonDriver struct{}

func (p PersonDriver) Work() {
	fmt.Println("Person is working")
}

func (p PersonDriver) Drive() {
	fmt.Println("Person is driving")
}

// JSON Handling
func jsonHandling() {
	fmt.Println("\nJSON Handling:")

	person := Person{
		Name: "John",
		Age:  30,
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
	} else {
		fmt.Println("JSON Data:", string(jsonData))
	}

	var person2 Person
	err = json.Unmarshal(jsonData, &person2)
	if err != nil {
		fmt.Println("Error unmarshalling from JSON:", err)
	} else {
		fmt.Printf("Unmarshalled Person: %+v\n", person2)
	}
}

// Struct Embedding
func structEmbeddingExamples() {
	fmt.Println("\nStruct Embedding:")

	person := Person{
		Name: "John",
		Age:  30,
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	fmt.Printf("Person: %+v\n", person)
}

type Address struct {
	City, State string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

// Main function to run all examples
func runExamples() {
	variableExamples()
	constantsExamples()
	slicesAndMapsExamples()
	controlStructures()
	functionExamples()
	pointerExamples()
	interfaceExamples()
	errorHandling()
	concurrencyExamples()
	deferredFunctionCalls()
	panicAndRecover()
	typeAssertions()
	embeddingInterfaces()
	jsonHandling()
	structEmbeddingExamples()
}
