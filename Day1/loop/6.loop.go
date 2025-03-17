package main

import "fmt"

func main() {
	fmt.Println("Labeled control statements")
outer:
	for i := 0; i < 10; i++ {

		for j := 10; j > 0; j-- {
			fmt.Println(i, j)
			if i == 4 && j == 5 {
				continue outer
			}
			// if i == 2 && j == 3 {
			// 	break outer
			// }
			// if i == 9 {
			// 	break inner
			// }
		}
	}
}
