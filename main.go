package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("- TASK 1 -")
	printSegitiga(5)

	println()

	fmt.Println("- TASK 2 -")
	genPass("fazztract", "Low")
	genPass("fazztract", "Medium")
	genPass("fazztract", "Strong")

	println()

	fmt.Println("- TASK 3 -")
	fmt.Println(choosedFilm(6))

}

func printSegitiga(rows int) {

	for i := rows; i >= 1; i-- {
		for k := 1; k <= (i*2)-1; k++ {
			fmt.Print("*")
		}
		println()

		for j := rows; j >= i; j-- {
			fmt.Print(" ")
		}
	}
}

func genPass(password string, level string) string {
	var randomChar = strconv.Itoa(rand.Intn(900) + 100)

	if len(password) >= 6 {
		if level == "Low" {
			fmt.Println(password)

			return password

		} else if level == "Medium" {
			password += randomChar

			slice := make([]string, len(password))
			for i := range password {
				if rand.Int()%2 == 0 {
					slice[i] = strings.ToLower(string(password[i]))
				} else {
					slice[i] = strings.ToUpper(string(password[i]))
				}
			}
			result := strings.Join(slice, "")
			fmt.Println(result)

			return result

		} else if level == "Strong" {
			slice := make([]string, len(password))
			for i := range password {
				if rand.Int()%2 == 0 {
					slice[i] = strings.ToLower(string(password[i]))
				} else {
					slice[i] = strings.ToUpper(string(password[i]))
				}
			}
			result := strings.Join(slice, "")

			specialChars := "!@#$%&*"
			for i := 0; i < 2; i++ {
				result += string(specialChars[rand.Intn(len(specialChars))])
			}

			result += randomChar
			fmt.Println(result)

			return result

		} else {
			return "level selection is't correct"
		}
	} else {
		return "Password min 6 characters"
	}
}

func choosedFilm(duration int) string {

	var data = [...]int{1, 7, 3, 4, 8, 9, 2}

	for i, value1 := range data {
		for j, value2 := range data {
			if value1+value2 == duration && i != j {
				result := fmt.Sprintf("%d and %d", value1, value2)
				return result
			}
		}
	}

	return "Duration not Found"
}
