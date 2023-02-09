package main

import "fmt"

func main() {
	data := 10
	fmt.Print(data)

	/*	if(condition){
			statements
		}else{
			statements
		}

		if(condition){
			statements
		}else{
			statements
		}

	*/
	fmt.Print("enter a number ")

	var input int
	fmt.Scanln(&input)

	if input%2 == 0 {
		fmt.Print("even")
	} else {
		fmt.Print("odd")
	}
}
