/*
Hello World

Outline:
● Files & Folders 
● Terminal
● TextEditor
● HelloWorld
● godoc 

Problems:
1. Create a program which prints “Hello World” to the terminal.
- Done
2. What is whitespace?
- Done
3. What is a comment? What are the two ways of writing a comment?
- Comment that is ignored by compiler.
- 2 ways: 
+ 1 line comment
+ block comment
4. Our program began with package main. What would the files in the
     fmt package begin with?
- package fmt
5. We used the Println function defined in the fmt package. If we
     wanted to use the Exit function from the os package what would we
need to do?
- you can code: import ("os", "fmt",)
6. Modify the program we wrote so that instead of printing Hello
     World it prints Hello, my name is followed by your name.
- Done
*/

package main

import "fmt"

// this is a comment

func main(){
	fmt.Println("Hello world")
	fmt.Println("Loc-Huynh-Tan")
}
