/*
Control Structures

Outline: 
● For
● If
● Switch

Problems:
● Write a program that prints out all the numbers evenly divisible
     by 3 between 1 and 100. (3, 6, 9, etc.)
● Write a program that prints the numbers from 1 to 100. But for
     multiples of three print "Fizz" instead of the number and for the
     multiples of five print "Buzz". For numbers which are multiples
     of both three and five print "FizzBuzz".
*/

package main

import "fmt"

func main(){
	// for i := 3; i < 100; i += 3 {
	// 	fmt.Printf("%d ", i)
	// }
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0{
			fmt.Printf("FizzBuzz ")
		} else if i % 3 == 0 {
			fmt.Printf("Fizz ")
		} else if i % 5 == 0 {
			fmt.Printf("Buzz ")
		}
	}
}

