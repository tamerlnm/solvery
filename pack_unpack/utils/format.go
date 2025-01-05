package utils

func Format(output string) string {
	outputWithEscapedNewline := ""
	for _, r := range output {
		if r == '\n' {
			outputWithEscapedNewline += "\\n"
		} else {
			outputWithEscapedNewline += string(r)
		}
	}
	return outputWithEscapedNewline
}
