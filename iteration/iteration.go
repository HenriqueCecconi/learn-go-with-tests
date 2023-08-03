package iteration

func Repeat(toBeRepeated string) string {
	var concatenated string
	for i := 0; i < 5; i++ {
		concatenated += toBeRepeated
	}
	return concatenated
}
