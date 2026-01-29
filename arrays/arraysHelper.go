package arrays

import "fmt"


func GenerateRange(min, max int) []int {
	result := []int{}

	if min > max{
		return nil
	}

	for i:=min ; i<=max; i++{
		result = append(result, i)
	}

	fmt.Println(result)

	return result

}

func RemoveElementsInRange(arr []float64, from, to int) (result []float64 ){

	if from >to {
		from, to = to, from
	}

	if from <0 {
		from =0
	}
	
	if to > len(arr) {
		to = len(arr)
	}

	for i,v := range arr {
		if i < from || i >= to {
			fmt.Println(v)
			result = append(result, v)
		}
	}
	return result

}

func BalanceOut(arr []bool)  []bool {
	trueCount :=0
	falseCount :=0

	for _,v := range arr {
		if v {
			trueCount++
		} else {
			falseCount++
		}
	}

	var diff int = trueCount - falseCount

	if trueCount  > falseCount{
		for i := 0; i < diff; i++ {
			arr = append(arr, false)
		}

	}


	if falseCount > trueCount{
		for i := 0; i < -diff; i++ {
			arr = append(arr, true)
		}
	}

	fmt.Println(arr)
	return arr

}

func SortIntegerTable(table []int) []int {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table) -i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}		
	}
	return table
}