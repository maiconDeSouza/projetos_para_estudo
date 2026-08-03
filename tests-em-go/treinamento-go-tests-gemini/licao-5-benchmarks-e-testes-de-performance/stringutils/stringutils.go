package stringutils

import "strings"

func ConcatenarComBuilder(n int) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString("a")
	}
	return builder.String()
}
