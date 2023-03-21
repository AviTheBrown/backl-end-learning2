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

func makeEven() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i // this is 0
		i += 2
		return ret
	}
}
	

func main() {
	theSlice := []int{1,2,3,4,5,6,7,8,9,10}
	averageNums(theSlice)
	add(theSlice...)
	theNext := makeEven()
	fmt.Println(theNext())
	fmt.Println(theNext())
	fmt.Println(theNext())


	}
