package twosum

// time complexity O(n^2)
func TwoSumDoubleLoop(nums []int, target int) []int {

	for i, n := range nums {
		for o, m := range nums {

			if n+m == target && i != o {

				return []int{i, o}
			}
		}
	}

	return []int{}
}

// time complexity O(n)
func TwoSumMap(nums []int, target int) []int {

	myMap := map[int]int{}

	for i, n := range nums {

		missing := target - n

		value, found := myMap[missing]
		if found {
			return []int{i, value}
		}

		myMap[n] = i
	}

	return []int{}
}
