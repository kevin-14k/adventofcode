package main

import ( 
	"fmt"
	"strings"
	"strconv"
)

func ConvertIntToArray(n int) ([]int) {
	var ari []int

	s   := strconv.Itoa(n)
	ars := strings.Split(s, "")
	for _, i := range ars {
		ie, _ := strconv.Atoi(i)
		ari = append(ari, ie)
	}

	return ari
}

func OnlyOneDuplcate(m map[string]int) bool {
	for _, value := range m {
		if (value == 2) { return true }
	}

	return false
}

func main() {
	nmatch := 0

	for input := 246540; input <= 787419; input++ {
		ari := ConvertIntToArray(input)
		m := map[string]int{"1": 0, "2": 0, "3": 0, "4": 0, "5": 0, "6": 0, "7": 0, "8": 0, "9": 0}

		prev := 0
		for i, v := range ari {
			if (prev > v) { break }

			m[strconv.Itoa(v)] += 1
			if (i == 5 && OnlyOneDuplcate(m)) { nmatch++ }

			prev = v
		}
	}

	fmt.Println(nmatch)
}