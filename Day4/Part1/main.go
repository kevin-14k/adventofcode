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

func main() {
	nmatch := 0

	for input := 246540; input <= 787419; input++ {
		ari := ConvertIntToArray(input)

		prev := 0
		dup  := false
		for i, v := range ari {
			if (prev > v) { break }
			if (prev == v) { dup = true }

			if (i == 5 && dup == true) { nmatch++ }

			prev = v
		}
	}

	fmt.Println(nmatch)
}