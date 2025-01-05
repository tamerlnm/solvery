package utils

func IsValidArgument(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s[0] >= '1' && s[0] <= '9' {
		return false
	}

	for i := 1; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= '9' && s[i-1] >= 48 && s[i-1] <= '9' {
			return false
		}
	}

	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '1' && s[i] <= '9' {
			count++
		}
	}
	if count == len(s) {
		return false
	}
	return true
}
