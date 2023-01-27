/*
Arrays, Slices and Maps

Outline:
● Arrays
● Slices 
● Maps

Problems:
1. How do you access the 4th element of an array or slice?
- arr[3]
2. What is the length of a slice created using: make([]int, 3, 9)?
- Ans: 3
3. Given the array:
           x := [6]string{"a","b","c","d","e","f"}
     what would x[2:5] give you?
- Ans: [c d e]
4. Write a program that finds the smallest number in this list: x := []int{
                48,96,86,68,
                57,82,63,70,
                37,34,83,27,
                19,97, 9,17,
}
- Ans: 9
5. Write a program that takes in a state code and returns the
     state’s name. (eg CA -> California)
*/

package main

import "fmt"

func main(){
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	min := 10000
	for i := 0; i < len(x); i++{
		if x[i] < min {
			min = x[i]
		} 
	} 
	fmt.Println(min)
}