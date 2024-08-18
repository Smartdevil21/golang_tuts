package main

import "fmt"

var a ="hello"

func main(){
	const name string = "Pratik";
	var surname string = "Vaidya";

	surname = "Smartdevil";

	fmt.Println(name);
	fmt.Println(surname);

	// Constant grouping
	const (
		port=4000
		host="localhost"
	)

	fmt.Println(port);
}