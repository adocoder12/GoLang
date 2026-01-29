package helpers

import "fmt"



func GetFirstRune(s string) rune {

	fmt.Println("first rune:", []rune(s)[0])
	return []rune(s)[0]


}

func GetLastRune(s string) rune {
	runes := []rune(s)
	fmt.Println("last rune:", runes[len(runes)-1])
	return runes[len(runes)-1]
}

func NRune(s string, i int) rune {
	runes := []rune(s)
	
	if s == "" {
		return 0
	}
	

	fmt.Println("rune:", string(runes[i]))

	return runes[i]

}

func StrLength(s string) []int {

	runes := []rune(s)
	strLen := len(s)
	runeLen := len(runes)

	fmt.Println("string length:", strLen)
	fmt.Println("rune length:", runeLen)

	return []int{strLen, runeLen}

}

func StrReverse(s string) string {
	runes := []rune(s)

	var reversed string
	for i := len(runes) - 1; i >= 0; i-- {
		reversed += string(runes[i])
		// fmt.Print(string(runes[i]))
	}
	fmt.Println(reversed)
	return reversed

}

func IsLower(s string) bool {

	for _,v := range s {
		if v < 'a' || v > 'z' {
			fmt.Println(false)
			return false
		}
	}
	fmt.Println(true)
	return true

}

func ToUpperCase(s string) string {
	var uppercased string
	for _,v := range s {
		if v >= 'a' && v <= 'z' {
			uppercased += string(v - ('a' - 'A'))
		} else {
			uppercased += string(v)
		}
	}
	fmt.Println(uppercased)
	return uppercased

}

func ToCapitalCase(s string) string {
	var capitalized string
	var firstLetterFound bool = false

	for _,v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			if !firstLetterFound {
				// Capitalize the first letter
				if v >= 'a' && v <= 'z' {
					capitalized += string(v - ('a' - 'A'))
				} else {
					// this letter is already uppercase
					capitalized += string(v)
				}
				firstLetterFound = true
			} else {
				// Lowercase subsequent letters
				if v >= 'A' && v <= 'Z' {
					capitalized += string(v + ('a' - 'A'))
				} else {
					capitalized += string(v)
				}
			}
		} else {
			// Non-letter characters
			capitalized += string(v)
			firstLetterFound = false
		}
	}
	fmt.Println(capitalized)
	return capitalized

}


func StrConcatWith(strs []string, sep string) string {

	var result string
	result = strs[0]
	for _, v := range strs {
		result += sep + v
	}
	fmt.Println(result)
	return result
}

func SplitWhitespaces(s string) []string {
	var result []string
	var currentWord string

	for _,v := range s {
		if v == ' ' || v == '\n' || v == '\t' {
			fmt.Println("currentWord:", currentWord)
				result = append(result, currentWord)
				currentWord = ""
		} else {
			currentWord += string(v)
		}
	}
	result = append(result, currentWord)
	fmt.Println(result)
	return result


}

func StrSplitBy(s, sep string) []string {

	var currentWord string
	var result []string

	sepRunes := []rune(sep)
	sepLen := len(sepRunes)
	sRunes := []rune(s)

	i := 0
	for i < len(sRunes) {
		if i <= len(sRunes)-sepLen && string(sRunes[i:i+sepLen]) == sep {
			result = append(result, currentWord)
			currentWord = ""
			i += sepLen
		} else {
			currentWord += string(sRunes[i])
			i++
		}
	}
	result = append(result, currentWord)
	fmt.Println(result)
	return result

}

func SubstrIndex(s string, toFind string) int {


    if (toFind == "") {
		return 0
	}
	
	for i,_ := range s{
		if i+len(toFind) <= len(s) {
            if s[i:i+len(toFind)] == toFind {
				fmt.Println(i)
                return i
            }
        }
	}
	fmt.Println(-1)
	return -1

}