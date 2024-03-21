package challenges

// solve more problems
import (
	"fmt"
	"sort"
)

// IndexOf finds the index of an element in a slice
func IndexOf(slice []int, target int) int {
	for i, element := range slice {
		if element == target {
			return i
		}
	}
	return -1 // Return -1 if the element is not found
}

// RemoveAtIndex removes an element at a given index from a slice
func RemoveAtIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func getMaxElementInLimit(slice []int, limit int) int {
	answer := 0
	for _, element := range slice {
		if element <= limit {
			if element > answer {
				answer = element
			}
		}
	}
	return answer
}

func LatestClock(a int, b int, c int, d int) string {
	numList := []int{a, b, c, d}
	sort.Ints(numList)
	two_index := IndexOf(numList, 2)
	one_index := IndexOf(numList, 1)
	zero_index := IndexOf(numList, 0)

	first := 0
	second := 0
	if two_index != -1 {
		first = 2
		numList = RemoveAtIndex(numList, two_index)
		second = getMaxElementInLimit(numList, 3)
	} else if one_index != -1 {
		first = 1
		numList = RemoveAtIndex(numList, one_index)
		second = getMaxElementInLimit(numList, 9)
	} else {
		first = 0
		numList = RemoveAtIndex(numList, zero_index)
		second = getMaxElementInLimit(numList, 9)
	}

	// now remove second element too
	numList = RemoveAtIndex(numList, IndexOf(numList, second))

	// find mins
	third := getMaxElementInLimit(numList, 5)
	numList = RemoveAtIndex(numList, IndexOf(numList, third))
	fourth := numList[0]
	latest_time := fmt.Sprintf("%d%d:%d%d", first, second, third, fourth)

	return latest_time
}
