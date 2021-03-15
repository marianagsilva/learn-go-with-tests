package iteration

//const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	repeatCount := 10

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
