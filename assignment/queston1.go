package main

import (
	"fmt"

	"golang.org/x/text/cases"
)

func main() {

	var name string
	fmt.Println("Hello there \n Kindly enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Welcome, Mr.%v ", CapitalizeWords(name))

}

func CapitalizeWords(words string) string {
	cases.Title(words)
	return words

}
