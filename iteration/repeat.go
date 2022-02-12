package iteration

func Repeat(char string, j int) string {
	var repeated string
	for i := 0; i < j; i++ {
		repeated += char
	}
	return repeated
}
