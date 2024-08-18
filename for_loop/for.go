package main

import "fmt"

// No while or do while loop in golang
func main(){
	// Implementation of while loop using for loop
	i:=1;
	for  i <= 10 {
		// fmt.Println(i)
		i=i+1
	}

	// Classic for loop
	for j:=0; j <3; j++{
		if j == 1{
			continue;
		}
		// fmt.Println(j)
	}

	// Range in golang
	for k:=range 3{
		fmt.Println(k)
	}
}