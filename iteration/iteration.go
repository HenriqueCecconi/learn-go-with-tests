package iteration

const repeatCount = 5

func Repeat(toBeRepeated string, repeatAmount int) string {
	var concatenated string
	if repeatAmount == 0 {
		repeatAmount = repeatCount
	}
	for i := 0; i < repeatAmount; i++ {
		concatenated += toBeRepeated
	}
	return concatenated
}
