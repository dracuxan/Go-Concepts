package arrays

func Sum(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

func SumAll(numsToSum ...[]int) []int {
	var sums []int
	for _, nums := range numsToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}

func SumAllTails(numsToSum ...[]int) []int {
	var sums []int
	for _, nums := range numsToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(nums[1:]))
		}
	}

	return sums
}
