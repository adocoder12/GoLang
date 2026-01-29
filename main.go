package main

import (
	"Hive/helpers"
	"fmt"
)


func main() {
	fmt.Println("Hello, World!")
	// helpers.GenerateRange(-5, 5)
	// helpers.RemoveElementsInRange([]float64{10., .8, -.4, 20., 7.7, 3.}, 4, 1)
	// helpers.BalanceOut([]bool{true, false, false, false})
	// helpers.SortIntegerTable([]int{2, 0, 5, 4, 1, 3})
	helpers.GetFirstRune("Hello")
	helpers.GetLastRune("Hello")
	helpers.NRune("Coding is cool", 4)
	helpers.StrLength("Hello")
	helpers.StrReverse("Hello, World!")
	helpers.IsLower("hello")
	helpers	.ToUpperCase("zorro")
	helpers.ToCapitalCase("Hello! Great to see you! How-are-you-doing-2day?")
	helpers.StrConcatWith([]string{"Three", " Two", " One", " Go!"}, " - ")
	helpers.SplitWhitespaces("Hello! How have you been?")
	helpers.StrSplitBy("HowYOUhaveYOUyouYOUbeen?", "YOU")

}