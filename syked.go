package main

import (
	"fmt"
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func BasicAtoi(s string) int {
	result := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		}
	}
	return result
}

func BasicAtoi2(s string) int {
	started := false
	result := 0
	for _, char := range s {

		if char >= '0' && char < '9' && !started {

			result = result*10 + int(char-'0')
			started = true
		} else if char >= '0' && char < '9' && started {

			result = result*10 + int(char-'0')

		} else {
			return 0
		}
	}
	return result
}

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	started := false

	for _, char := range s {
		if char == '-' && !started {
			sign = -1
			started = true
		} else if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
			started = true
		}
	}
	return sign * result
}

func Atoi(s string) int {

	started := false
	result := 0
	sign := 1

	for _, char := range s {

		if char == ' ' && !started {
			continue
		}

		if char == '+' && !started {
			sign = 1
			started = true
		}
		if char == '-' && !started {
			sign = -1
			started = true
			continue
		}
		if char < '0' || char > '9' {

			break
		}
		result = result*10 + int(char-'0')

	}
	return result * sign
}

func SortIntegerTable(table []int) {
	n := len(table)
	for i := 0; i < n; i++ {
		swaped := false
		for j := 0; j < n-1; j++ {

			if table[j] < table[j+1] {

				table[j], table[j+1] = table[j+1], table[j]
				swaped = true
			}

		}
		if !swaped {
			break
		}

	}

}

func IterativeFactorial(nb int) int {
	r := 1
	for i := nb; i > 0; i-- {
		r = r * i

	}
	return r
}

func RecursiveFactorial(nb int) int {
	r := 1
	if nb == 0 {
		return 1
	}
	if nb > 0 {
		r = RecursiveFactorial(nb-1) * nb
	}
	return r
}

func IterativePower(nb int, power int) int {
	r := 1
	for i := 0; i < power; i++ {

		r = r * nb
	}
	return r
}
func IsPrime(nb int) bool {

	if nb < 0 {
		return false
	}

	for i := 1; i <= nb; i++ {

		if nb%i == 0 && i != nb && i != 1 {
			return false
			break
		}

	}
	return true

}

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	prime := nb
	// If number is even, make it odd as even numbers greater than 2 are not prime
	if prime%2 == 0 {
		prime++
	}
	// Check only odd numbers for primality
	for !IsPrime(prime) {
		prime += 2
	}
	return prime
}

func Index(s string, toFind string) int {

	for i := 0; i < len(s)-len(toFind); i++ {

		if s[i:i-len(toFind)] == toFind {
			return i
		}

	}
	return -1
}

func Capitaliza(s string) string { //print fisrt car upper and rest lower

	result := []rune(s) // Convert the string to a slice of runes to handle each character
	started := true

	for i, char := range result {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			if started {
				if char >= 'a' && char <= 'z' {
					result[i] = char - 32
				}
				started = false

			} else {
				if char >= 'A' && char <= 'Z' {
					result[i] = char + 32
				}
			}
		} else {

			started = true
		}
	}
	return string(result)
}

// Itoa converts an integer to a string.
func Itoa(n int) string {

	if n == 0 {
		return "0"
	}

	var table []byte
	isnegativ := false

	if n < 0 {
		n = -n
		isnegativ = true

	}

	for n > 0 {
		digit := byte(n%10) + '0'
		table = append(table, digit)
		n /= 10
	}

	if isnegativ {
		table = append(table, '-')
	}

	for j := 0; j < len(table); j++ {

		for i := 0; i < len(table)-1; i++ {

			if table[i] > table[i+1] {

				table[i], table[i+1] = table[i+1], table[i]
			}
		}
	}

	return string(table)
}
func PrintNbrInOrder(n int) {
	if n == 0 {
		fmt.Print("0")
	}

	var table []byte

	for n > 0 {
		num := byte(n%10) + '0'
		table = append(table, num)
		n = n / 10
	}

	for j := 0; j < len(table); j++ {

		for i := 0; i < len(table)-1; i++ {

			if table[i] > table[i+1] {

				table[i], table[i+1] = table[i+1], table[i]
			}
		}
	}
	fmt.Print(string(table))

}

func print_param() {
	//choumi is best cat
	for _, args := range os.Args[1:] {

		for _, w := range args {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}

func revparams() {
	//cat best the choumi
	arc := len(os.Args) - 1
	for i := arc; i > 0; i-- {

		for _, w := range os.Args[i] {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}

func sort_param() {

	arg := os.Args
	for i := 0; i < len(os.Args)-1; i++ {
		for j := 0; j < len(os.Args)-1; j++ {
			if arg[j] > arg[j+1] {

				arg[j], arg[j+1] = arg[j+1], arg[j]
			}
		}
	}
	for _, args := range os.Args[1:] {

		for _, w := range args {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}
//////////////////////////////////////////////////////////////////////////////////
func SplitWhiteSpaces(s string) []string {
	var result []string
	var word string

	for i, char := range s {
		if char != ' ' && char != '\n' && char != '\t' {
			word = word + string(char)
		}
		if ((char == ' ' || char == '\n' || char == '\t') && word != "") || i == len(s)-1 {

			result = append(result, word)
			word = ""
		}
	}
	return result
}

///////////////////////////* inter */ //////////////////////////////////////////////////
func main() {
	
	if len(os.Args) != 3 {
		return // if not exactly two arguments, exit the program without printing anything
	}
	str1 := os.Args[1]
	str2 := os.Args[2]

	result := ""

	for _, c1 := range str1 {
		if !stringContains(result, c1) && stringContains(str2, c1) {
			result += string(c1)
		}
	}

	fmt.Println(result)
}

// stringContains checks if the string `s` contains the character `r`.
func stringContains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}
/* /////////////////////////////// FROM TO ///////////////////////////////////////// */
func FromTo(from int, to int) string {
	
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	result := ""
	var i int

	// Check if we are counting up or down.
	if from <= to {
		for i = from; i <= to; i++ {
			result += formatNumber(i)
			
			if i < to {
				result += ", "
			}
		}
	} else { // from is greater than to, so we count down
		for i = from; i >= to; i-- {
			result += formatNumber(i)
			
			if i > to {
				result += ", "
			}
		}
	}

	result += "\n"

	return result
}

func formatNumber(number int) string {
	if number < 10 {
		return fmt.Sprintf("0%d", number)
	}
	return fmt.Sprintf("%d", number)
}
/* ////////////////////////////////////////////////////////*/

func main() {

	fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))

}

/*
s := strconv.Itoa(n)
	digits := []rune(s)
*/

