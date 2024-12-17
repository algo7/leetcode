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
