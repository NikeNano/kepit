package main

// Person holds person information
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var ExamplePerson Person
