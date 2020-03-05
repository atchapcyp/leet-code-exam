package palindrome

import (
	"fmt"
)

func IsPalindrome(x int) bool {
	var intArr []int
	var size int
	size = findIntSize(x)
	fmt.Println("size", size)
	fmt.Println("input", x)
	if size == 1 {
		return x >= 0
	}
	for i := 1; i <= size; i++ {
		digit := x % 10
		if i <= size/2 {
			intArr = append(intArr, digit)
		}

		if i == size/2+1 && size%2 == 1 {
			x = x / 10
			continue
		}

		if i > size/2 {

			if digit < 0 {
				return false
			}

			fmt.Printf("size(%d) - i(%d) %d ", size, i, size-i)
			fmt.Printf("compare %d  %d\n", digit, intArr[size-i])
			if digit != intArr[size-i] {
				return false
			}

		}
		x = x / 10
		fmt.Println("intArr", intArr)
	}
	return true
}

func findIntSize(x int) int {
	var size int
	for {
		if x/10 == 0 {
			return size + 1
		}
		x = x / 10
		size++
	}
	return 0
}

func IsPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	if x >= 0 && x < 10 {
		return true
	}
	var result int = 0
	var org int = x
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}
	return result == org
}
