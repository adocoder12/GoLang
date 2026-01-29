package main

import (
	"Hive/strManipulation"
	"fmt"
)


func main() {
	fmt.Println("Hello, World!")
	// arrays.GenerateRange(-5, 5)
	// arrays.RemoveElementsInRange([]float64{10., .8, -.4, 20., 7.7, 3.}, 4, 1)
	// arrays.BalanceOut([]bool{true, false, false, false})
	// arrays.SortIntegerTable([]int{2, 0, 5, 4, 1, 3})
	strManipulation.GetFirstRune("Hello")
	strManipulation.GetLastRune("Hello")
	strManipulation.NRune("Coding is cool", 4)
	strManipulation.StrLength("Hello")
	strManipulation.StrReverse("Hello, World!")
	strManipulation.IsLower("hello")
	strManipulation.ToUpperCase("zorro")
	strManipulation.ToCapitalCase("Hello! Great to see you! How-are-you-doing-2day?")
	strManipulation.StrConcatWith([]string{"Three", " Two", " One", " Go!"}, " - ")
	strManipulation.SplitWhitespaces("Hello! How have you been?")
	strManipulation.StrSplitBy("HowYOUhaveYOUyouYOUbeen?", "YOU")


}