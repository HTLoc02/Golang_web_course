/*
Variables

Outline:
● Samples
● BlankIdentifiers
● ShortVariableDeclarations 
● Assignment
● ZeroValues
● Scope
● Constants

Problems:
1. Create a program which prints “Hello <NAME>” with <NAME> replaced
     with your name to the terminal using a variable.
- Done
2. Use fmt.Scanf to read a user’s name and print “Hello <NAME>” with
     <NAME> replaced with the user’s name to the terminal.
- Done
*/

package main

import "fmt"

func main(){
	var name string
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello %s", name)
}


