package main

import (
		"fmt"
        "os"
        "log"
        "bufio"
        "strconv"
)

func ReadData() ([]int) {
	file, err := os.Open("data")
	
	if err != nil {
    	log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data_array := []int{}
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())

		if err == nil {
	    	data_array = append(data_array, n)
	    }
	}

	if err := scanner.Err(); err != nil {
    	log.Fatal(err)
	}

	return data_array
}

func CalculateMass(mass int) int {
	return mass / 3 - 2
}

func ProcessData(mass []int) int {
	fuel := 0

	for i := 0; i < len(mass); i++ {
		added_mass := mass[i]

		for (added_mass >= 0) {
			added_mass = CalculateMass(added_mass)

			if (added_mass < 0) {
				break
			}

			fuel += added_mass
		}
	}

	return fuel
}

func main() {
    fmt.Println(ProcessData(ReadData()))
}