package array

func Sum(a []int) int  {
	sum := 0
	for _, num := range a {
		sum += num
	}
	return sum
}

func SumEach(a ...[]int) []int {
	var sums []int
	for _, nums := range a {
		sums = append(sums, Sum(nums))
	}
	return sums
}
