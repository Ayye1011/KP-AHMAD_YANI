package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	freq := make(map[int]int)
	for _, char := range angka {
		digit, _ := strconv.Atoi(string(char))
		freq[digit]++
	}

	var result []int
	for _, char := range angka {
		digit, _ := strconv.Atoi(string(char))
		if freq[digit] == 1 {
			result = append(result, digit)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
