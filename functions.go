package main

import (
	"fmt"
)


func averageNums(intSlice []int) float64 {
	var total float64
	var sum float64 = 0.0
	for i := 0; i < len(intSlice); i++ {
		sum += float64(intSlice[i]) 
		}
	total = sum / float64(len(intSlice))
	fmt.Println(total)
	return total
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	fmt.Println(total)
	return total
}
	

func main() {
	theSlice := []int{1,2,3,4,5,6,7,8,9,10}
	averageNums(theSlice)
	add(theSlice...)

}
