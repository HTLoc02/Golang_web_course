/*
Types

Outline:
● Integers
● FloatingPointNumbers 
● Strings
● Booleans

Problems:
1. How are integers stored on a computer?
- it's depends
- a lot of types of integers:  uint8, uint16, uint32, uint64, int8, 
int16, int32, int64, int.
2. We know that (in base 10) the largest 1 digit number is 9 and the
largest 2 digit number is 99. Given that in binary the largest 2 digit 
number is 11 (3), the largest 3 digit number is 111 (7) and the largest
4 digit number is 1111 (15) what's the largest 8 digit number? 
(hint: 10^1-1 = 9 and 10^2-1 = 99)
- Ans: 10^8 - 1
3. Although overpowered for the task you can use Go as a calculator. Write 
a program that computes 32132 × 42452 and prints it to the terminal. 
(Use the * operator for multiplication)
- Ans: 1364067664
4. What is a string? How do you find its length?
- Ans: we use len() function
5. What's the value of the expression (true && false) || (false &&
     true) || !(false && false)?
- Ans: false || false || !false >> true
*/

package main

import "fmt"

func main(){
	// fmt.Println(32132 * 42452)
	var str string
	str = "Hello world"
	fmt.Println(len(str))
	// fmt.Println((true && false) || (false && true) || !(false && false))
}