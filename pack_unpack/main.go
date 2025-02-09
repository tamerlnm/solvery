package main

import (
	"coder/utils"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var arg string
	input := flag.Bool("input", false, "input argument")
	daemon := flag.Bool("daemon", false, "daemon argument")
	pack := flag.Bool("pack", false, "pack argument")
	unpack := flag.Bool("unpack", false, "unpack argument")
	flag.Parse()

	if *input {
		arg = os.Args[3]
	} else if *daemon {
		fmt.Print("Введите строку: ")
		fmt.Scanln(&arg)
	}

	if !utils.IsValidArgument(arg) {
		fmt.Println(`""`)
		return
	}

	if *pack {
		fmt.Println(packValue(arg))
	} else if *unpack {
		fmt.Println(unpackValue(arg))
	}

}

func unpackValue(arg string) string {
	arg, _ = strconv.Unquote(`"` + arg + `"`)
	var builder strings.Builder

	prevRune := rune(0)
	for _, r := range arg {
		if r < '0' || r > '9' {
			builder.WriteRune(r)
			prevRune = r
			continue
		}

		n, _ := strconv.Atoi(string(r))
		if n == 0 {
			str := builder.String()
			if len(str) > 0 {
				builder.Reset()
				builder.WriteString(str[:len(str)-1])
			}
			continue
		}
		utils.AppendToBuilder(&builder, prevRune, n)
	}
	return utils.Format(builder.String())
}

func packValue(arg string) string {
	result := ""
	count := 1
	runes := []rune(arg)

	for i, r := range runes {
		if i+1 < len(runes) && r == runes[i+1] {
			count++
			continue
		}

		result += string(r)
		if count > 1 {
			result += strconv.Itoa(count)
		}
		count = 1
	}
	return result
}
