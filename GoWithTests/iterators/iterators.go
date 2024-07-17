package iterators

func Repeat(character string, numberOfRepetitions int) string {
	var repetitions string
	for i := 0; i < numberOfRepetitions; i++ {
		repetitions += string(character)
	}
	return repetitions
}
