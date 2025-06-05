package summator

func SumAll(slcs ...[]int) []int {
	result := make([]int, len(slcs))

	for idx, slc := range slcs {
		result[idx] = sumOfIntSlice(slc)
	}

	return result
}

func sumOfIntSlice(slc []int) int {
	var sum int

	for _, num := range slc {
		sum += num
	}

	return sum
}
