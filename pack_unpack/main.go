package main

import (
	"coder/utils"
	"flag"
	"fmt"
	"os"
	"strconv"
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
	result := make([]rune, 0)
	prevRune := rune(0)
	for _, r := range arg {
		if r < '0' || r > '9' {
			result = append(result, r)
			prevRune = r
			continue
		}

		n, _ := strconv.Atoi(string(r))
		if n == 0 && len(result) > 0 {
			result = result[:len(result)-1]
		}
		utils.AppendToSliceOfRune(&result, prevRune, n)
	}
	return utils.Format(string(result))
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
