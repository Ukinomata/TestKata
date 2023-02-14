package main

import (
	"fmt"
	romans "github.com/summed/goromans" // go get -u "github.com/summed/goromans"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Введите уравнение: ")
	var str, sign, per2 string
	fmt.Scanln(&str, &sign, &per2)
	add := Index(str, '+')
	substract := Index(str, '-')
	multiply := Index(str, '*')
	divide := Index(str, '/')
	ind := -1
	k := 0
	result := 0
	if sign == "" {
		if add == LastIndex(str, '+') {
			if add > -1 {
				k += 1
				ind = add
			}
		}
		if substract == LastIndex(str, '-') {
			if substract > -1 {
				k += 1
				ind = substract
			}
		}
		if multiply == LastIndex(str, '*') {
			if multiply > -1 {
				k += 1
				ind = multiply
			}
		}
		if divide == LastIndex(str, '/') {
			if divide > -1 {
				k += 1
				ind = divide
			}
		}
		if k == 1 {
			var ka int
			var ko int
			rpp1, ka := cases(str[:ind])
			rpp2, ko := cases(str[ind+1:])
			if rpp1 > 10 || rpp2 > 10 {
				fmt.Println("Invalid Input1")
				os.Exit(1024)
			}
			if ka != ko {
				fmt.Println("Invalid Input2")
				os.Exit(1024)
			}
			sign += str[ind : ind+1]
			if sign == "+" {
				result = rpp1 + rpp2
			} else if sign == "-" {
				result = rpp1 - rpp2
			} else if sign == "*" {
				result = rpp1 * rpp2
			} else if sign == "/" {
				result = rpp1 / rpp2
			} else {
				fmt.Println("Invalid Input3")
				os.Exit(1024)
			}
			if ka == 1 {
				if result <= 0 {
					fmt.Println("Invalid Input4")
					os.Exit(1024)
				} else {
					fmt.Println(romans.AtoR(uint(result)))
				}
			} else if ka == 2 {
				fmt.Println(result)
			} else {
				fmt.Println("Invalid Input5")
				os.Exit(1024)
			}

		} else {
			fmt.Println("Invalid Input6")
			os.Exit(1024)
		}
	} else if per2 == "" {
		fmt.Println("Invalid Input7")
		os.Exit(1024)
	} else {
		var ka int
		var ko int
		rpp1, ka := cases(str)
		rpp2, ko := cases(per2)
		if rpp1 > 10 || rpp2 > 10 {
			fmt.Println("Invalid Input8")
			os.Exit(1024)
		}
		if ka != ko {
			fmt.Println("Invalid Input9")
			os.Exit(1024)
		}
		if sign == "+" {
			result = rpp1 + rpp2
		} else if sign == "-" {
			result = rpp1 - rpp2
		} else if sign == "*" {
			result = rpp1 * rpp2
		} else if sign == "/" {
			result = rpp1 / rpp2
		} else {
			fmt.Println("Invalid Input10")
			os.Exit(1024)
		}
		if ka == 1 {
			if result <= 0 {
				fmt.Println("Invalid Input11")
				os.Exit(1024)
			} else {
				fmt.Println(romans.AtoR(uint(result)))
			}
		} else if ka == 2 {
			fmt.Println(result)
		} else {
			fmt.Println("Invalid Input12")
			os.Exit(1024)
		}
	}
}

func cases(rpp string) (int, int) {
	k := 0
	result := 0
	switch rpp {
	case "I":
		result = 1
	case "II":
		result = 2
	case "III":
		result = 3
	case "IV":
		result = 4
	case "V":
		result = 5
	case "VI":
		result = 6
	case "VII":
		result = 7
	case "VIII":
		result = 8
	case "IX":
		result = 9
	case "X":
		result = 10
	}
	if result == 0 {
		k = 2
		result, _ = strconv.Atoi(rpp)
	} else {
		k = 1
	}
	return result, k
}

func Index(str string, ch uint8) int {
	ind := -1
	for i := 0; i < len(str); i++ {
		if str[i] == ch {
			ind = i
			break
		}
	}
	return ind
}

func LastIndex(str string, ch uint8) int {
	ind := -1
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ch {
			ind = i
			break
		}
	}
	return ind
}
