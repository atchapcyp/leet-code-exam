package twosum

import "fmt"

func TwoSum(nums []int, target int) []int {
	mapInput := make(map[int]int, len(nums))
	for _, n := range nums {
		mapInput[n]++
	}

	for numKey, val := range mapInput {
		numFind := target - numKey
		if numKey == numFind && val == 1 {
			continue
		}
		if mapInput[numFind] > 0 {
			fmt.Printf("findIndexOf numKey %d , numFind %d\n", numFind, numKey)
			for k, v := range mapInput {
				fmt.Printf("k-> %d,v-> %d\n", k, v)
			}
			return findIndexOf(numKey, numFind, nums)
		}
	}
	return []int{-1, -1}
}

func findIndexOf(f, s int, arr []int) []int {
	indexRes := []int{-1, -1}

	for i, ar := range arr {

		if f == ar && indexRes[0] == -1 {
			indexRes[0] = i
			continue
		}
		if s == ar && indexRes[1] == -1 {
			indexRes[1] = i
		}

		if indexRes[0] != -1 && indexRes[1] != -1 {
			break
		}
	}

	if indexRes[0] > indexRes[1] {
		fmt.Printf("swap %d <-> %d\n", indexRes[0], indexRes[1])
		indexRes[0], indexRes[1] = indexRes[1], indexRes[0]
	}

	return indexRes
}
