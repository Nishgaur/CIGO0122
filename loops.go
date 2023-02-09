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

		}*/
	// cant use independentlylist:= range
	nums := []int{1, 2, 3, 4, 0}
	for k, v := range nums {
		fmt.Println(k)
		fmt.Println(v)
	}
}
