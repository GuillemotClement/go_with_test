package iteration

import "strings"

//const repeatCount = 5

func Repeat(character string, numberRepeat int) string {
	var repeated strings.Builder
	for i := 0; i < numberRepeat; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
