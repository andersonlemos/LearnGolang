package slicesAndArrays

func Sum(array []int) (sum int) {

	for _, value := range array {
		sum += value
	}
	return sum
}

/*func SumAll(array ...[]int) (sum []int) {
	arrCount := len(array)
	sumOf := make([]int, arrCount)

	for idx, value := range array {
		sumOf[idx] = Sum(value)
	}

	return
}*/

func SumAll(arrays ...[]int) (sums []int) {
	var sum []int
	for _, values := range arrays {
		sum = append(sum, Sum(values))
	}
	return sum
}

func SumEvereythingElse(arrays ...[]int) []int {
	var sums []int
	for _, values := range arrays {
		if len(values) == 0 {
			sums = append(sums, 0)
		} else {
			final := values[1:]
			sums = append(sums, Sum(final))
		}
	}
	return sums
}
