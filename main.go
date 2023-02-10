package main

import (
	"fmt"
	romans "github.com/summed/goromans" // go get -u "github.com/summed/goromans" в терминал
	"os"
	"strconv"
)

func main() {
	fmt.Println("Go calculator!")
	ent := readLine("Enter command: [a]rabic, [r]oman:")
	if ent == "a" {
		cmd := readLine("Enter command: [a]dd, [s]ubtract, [m]ultiply, [d]ivide: ")
		fmt.Print(cmd)
		if cmd == "a" {
			num1, num2 := getUserNumbers()
			if num1 <= 10 {
				if num2 <= 10 {
					result := num1 + num2
					fmt.Println(fmt.Sprintf("%d", result))
				}
			} else {
				fmt.Println("Invalid input")
				os.Exit(13)
			}
		} else if cmd == "s" {
			num1, num2 := getUserNumbers()
			if num1 <= 10 {
				if num2 <= 10 {
					result := num1 - num2
					fmt.Println(fmt.Sprintf("%d", result))
				}
			} else {
				fmt.Println("Invalid input")
				os.Exit(13)
			}
		} else if cmd == "m" {
			num1, num2 := getUserNumbers()
			if num1 <= 10 {
				if num2 <= 10 {
					result := num1 * num2
					fmt.Println(fmt.Sprintf("%d", result))
				}
			} else {
				fmt.Println("Invalid input")
				os.Exit(13)
			}
		} else if cmd == "d" {
			num1, num2 := getUserNumbers()
			if num1 <= 10 {
				if num2 <= 10 {
					result := num1 / num2
					fmt.Println(fmt.Sprintf("%d", result))
				}
			} else {
				fmt.Println("Invalid input")
				os.Exit(13)
			}
		} else {
			fmt.Println("Invalid input")
			os.Exit(13)
		}
		readLine("press 'enter' to exit")
	} else if ent == "r" {
		cmd := readLine("Enter command: [a]dd, [s]ubtract, [m]ultiply, [d]ivide: ")
		fmt.Print(cmd)
		if cmd == "a" {
			num1, num2 := getUserNumbersRoman()
			result := num1 + num2
			fmt.Println(romans.AtoR(uint(result)))
		} else if cmd == "s" {
			num1, num2 := getUserNumbersRoman()
			result := num1 - num2
			fmt.Println(romans.AtoR(uint(result)))
		} else if cmd == "m" {
			num1, num2 := getUserNumbersRoman()
			result := num1 * num2
			fmt.Println(romans.AtoR(uint(result)))
		} else if cmd == "d" {
			num1, num2 := getUserNumbersRoman()
			result := float32(num1) / float32(num2)
			fmt.Println(romans.AtoR(uint(result)))
		} else {
			fmt.Println("Invalid input")
			os.Exit(13)
		}
		readLine("press 'enter' to exit")
	} else {
		fmt.Println("Invalid input")
		os.Exit(13)
	}
}
func readLine(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func getUserNumbers() (int, int) {
	num1String := readLine("First Number: ")
	num1, _ := strconv.Atoi(num1String)
	num2String := readLine("Second Number: ")
	num2, _ := strconv.Atoi(num2String)
	return num1, num2
}

func getUserNumbersRoman() (int, int) {
	num1String := readLine("First Number: ")
	num1 := test(num1String)
	num2String := readLine("Second Number: ")
	num2 := test(num2String)
	return num1, num2
}

func test(num string) int {
	k := 0
	romanmap2 := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	romanmap := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X"}
	for i := 1; i <= 10; i++ {
		if num == romanmap[i] {
			k += 1
		}
	}
	if k == 0 {
		fmt.Println("Invalid Input")
		os.Exit(13)
	}
	m := romanmap2[num]
	return m
}
