package utils

import "strings"

func AppendToSliceOfRune(result *[]rune, prevRune rune, n int) {
	for i := 1; i < n; i++ {
		*result = append(*result, prevRune)
	}
}

func AppendToBuilder(builder *strings.Builder, prevRune rune, n int) {
	for i := 1; i < n; i++ {
		builder.WriteRune(prevRune)
	}
}
