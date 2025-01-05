package utils

func AppendToSliceOfRune(result *[]rune, prevRune rune, n int) {
	for i := 1; i < n; i++ {
		*result = append(*result, prevRune)
	}
}
