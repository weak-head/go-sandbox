package arr

func Sum(a []int) int {
	sum := 0
	for _, number := range a {
		sum += number
	}
	return sum
}

func SumAll(ars ...[]int) (sums []int) {
	for _, numbers := range ars {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(nums ...[]int) (sums []int) {
	for _, numbers := range nums {
		if len(numbers) > 0 {
			sums = append(sums, Sum(numbers[1:]))
		} else {
			sums = append(sums, 0)
		}
	}
	return
}
