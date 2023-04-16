package main

import "fmt"

func main() {
	var arrNum [5]int

	arrNum[0] = 10
	arrNum[1] = 9
	arrNum[2] = 8
	arrNum[3] = 7
	arrNum[4] = 6

	fmt.Println(arrNum)
	fmt.Println(arrNum[3])
	fmt.Println()
	arrNum2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrNum2)
	fmt.Println()
	for i, value := range arrNum2 {
		fmt.Printf("Value : %d, index of : %d\n", value, i)
	}
	fmt.Println()

	sliced := arrNum2[2:5]
	fmt.Println(sliced)

	sliced2 := make([]int, 5, 10)
	fmt.Println(sliced2)
	fmt.Println(len(sliced2))
	fmt.Println(cap(sliced2))
	copy(sliced2, arrNum2[:])
	fmt.Println(sliced2)

}
