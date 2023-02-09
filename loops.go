package main

import (
	"fmt"
)

func main() {
	/*
				for i:=0; i<20;i++{
					fmt.Println(i)
				}



			i := 0
			for {
				fmt.Println(i)
				i++
				if i < 3 {
					if i == 1 {
						continue
					}
					fmt.Println(i)
					break
				}

			}
		// cant use independentlylist:= range
		nums := []int{1, 2, 3, 4, 0}
		for k, v := range nums {
			fmt.Println(k)
			fmt.Println(v)
		}*/
	/*nums := []int{1, 2, 3, 4, 0}
	for _, v := range nums {

		fmt.Println(v)
	}
nums := []int{1, 2, 3, 4, 0}
	for _, value := range nums {
		if value == 3 {
			break
		}
		fmt.Println(value)
	}*/
	//nums := []int{1, 2, 3, 4, 0}
	for _, value := range "rahul" {
		/*if value == "u" {
			break
		}*/
		fmt.Println(value)
	}
}
