// https://leetcode.com/problems/find-the-duplicate-number/
package main

import (
	"fmt"
	"sort"
)

func main() {
	var numbers = []int{8, 2, 1, 9, 9}
	var numbers2 = []int{8, 2, 1, 9, 9}
	fmt.Println("The duplicate number is:", findDuplicate1(numbers))
	fmt.Println("The duplicate number is:", findDuplicate2(numbers2))
}

// findDuplicate1 takes in a slice of integers 'nums' and returns the duplicate number found in the slice.
// The function first sorts the slice in ascending order using the sort.Ints() function from the 'sort' package.
// It then iterates through the sorted slice starting from index 1, comparing each element with the previous element.
// If a duplicate number is found, it is assigned to the variable 'duplicateNumber', and the loop is exited.
// Finally, the function returns the 'duplicateNumber' found.
func findDuplicate1(nums []int) int {
	var duplicateNumber int

	sort.Ints(nums) // Since the sort.Ints() has to use a slice, we insert a semicolon(:).

	for i := 1; i < len(nums); i++ { // We start at 1, and then compare with the element before it, to avoid comparing outside the array.
		if nums[i] == nums[i-1] {
			duplicateNumber = nums[i]
			break
		}
	}

	return duplicateNumber
}

// findDuplicate2 takes in a fixed-size array of integers 'nums' and returns the duplicate number found in the array.
// It iterates through each element of the array and compares it with the remaining elements starting from the next index.
// If a duplicate number is found, it is assigned to the variable 'duplicateNumber' and returned immediately.
// If no duplicate is found, the function returns 0.
// Note that findDuplicate2 does not sort the array before finding the duplicate, unlike findDuplicate1.
func findDuplicate2(nums []int) int {
	var duplicateNumber int

	for i := range nums {

		for j := i + 1; j < len(nums); j++ { // To make sure we dont compare with the same index, we add 1 to j.
			if nums[i] == nums[j] {
				duplicateNumber = nums[i]
				return duplicateNumber
			}
		}
	}

	return duplicateNumber
}

/* ToDo: make this work with a stack instead.
func findDuplicate3(nums [5]int) int {

}*/
