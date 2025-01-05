package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Wrong data inputs")
	}
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

	if !isValidArgument(arg) {
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
	result := make([]rune, 0)
	for i := 0; i < len(arg); i++ {
		if arg[i] >= 48 && arg[i] <= 57 {
			n, err := strconv.Atoi(string(arg[i]))
			if err != nil {
				log.Fatal("error to convert string to integer")
			}
			if n == 0 && len(result) > 0 {
				result = result[:len(result)-1]
			} else {
				for n > 1 {
					if i > 1 && arg[i-2] == 92 {
						result = append(result, rune(92), rune(arg[i-1]))
					} else {
						result = append(result, rune(arg[i-1]))
					}
					n--
				}
			}
		} else {
			result = append(result, rune(arg[i]))
		}
	}
	return string(result)
}

func packValue(arg string) string {
	if len(arg) == 0 {
		return ""
	}

	result := ""
	count := 1
	for i := 0; i < len(arg); i++ {
		if i+1 < len(arg) && arg[i] == arg[i+1] {
			count++
		} else {
			result += string(arg[i])
			if count > 1 {
				result += strconv.Itoa(count)
			}
			count = 1
		}
	}
	return result
}

func isValidArgument(s string) bool {
	if s[0] >= 49 && s[0] <= 57 {
		return false
	}

	for i := 1; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= 57 && s[i-1] >= 48 && s[i-1] <= 57 {
			return false
		}
	}

	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 49 && s[i] <= 57 {
			count++
		}
	}
	if count == len(s) {
		return false
	}
	return true
}
