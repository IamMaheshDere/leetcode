package main

import "fmt"

func main() {
	fmt.Println("adding duplicate zeros")
	inputArr := []int{1, 5, 2, 0, 6, 8, 0, 6, 0}

	fmt.Println("input array is ", inputArr)

	duplicateZeros(inputArr)

	fmt.Println("output ", inputArr)
}

func duplicateZeros(inputArr []int) []int {
	for index := 0; index < len(inputArr)-1; index++ {
		if inputArr[index] != 0 {
			continue
		}

		rotateArrayRight(inputArr[index+1:])
		index = index + 1

	}
	return inputArr
}

func rotateArrayRight(inputArr []int) {
	if len(inputArr) == 0 {
		return
	}

	j := len(inputArr) - 1
	for ; j > 0; j-- {
		inputArr[j] = inputArr[j-1]
	}

	inputArr[0] = 0
}
