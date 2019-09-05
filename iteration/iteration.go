package iteration

func Repeat(character string) string {
	var repected string
	for i := 0; i < 5; i++ {
		repected += character
	}
	return repected
}
