package iteration

import "strings"

func Repeat(s string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += s
	}
	return repeated
}

func RepeatWithBuilder(s string) string {
	var sb strings.Builder
	for i := 0; i < 5; i++ {
		sb.WriteString(s)
	}
	return sb.String()
}

func RepeatWithStringsRepeat(s string) string {
	return strings.Repeat(s, 5)
}
