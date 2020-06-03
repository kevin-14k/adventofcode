package main

import ( 
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"regexp"
	"strconv"
)

func ReadData() string {
	data, err := ioutil.ReadFile("data")

	if err != nil {
		os.Exit(-1)
	}

	return string(data)
}

func ParseData(data string) ([]string, []string) {
	wires_path  := strings.Split(data, "\n")

	return strings.Split(wires_path[0], ","), strings.Split(wires_path[1], ",")
}

func ParseMove(step string) (string, int) {
	reg_num := regexp.MustCompile("[0-9]+")
	reg_cha := regexp.MustCompile("[A-Z]+")

	number := reg_num.FindAllString(step, -1)[0]
	letter := reg_cha.FindAllString(step, -1)[0]

	casted_number, _ := strconv.Atoi(number)

	return letter, casted_number
}

func ProducePath(path []string, tda [20][20]int) ([20][20]int) {
	x_c := 10
	y_c := 10

	for i := 0; i < len(path); i++ {
		letter, number := ParseMove(path[i])

		for n := 0; n < number; n++ {
			if letter == "R" {
				x_c++
				tda[y_c][x_c] += 1
			}

			if letter == "U" {
				y_c--
				tda[y_c][x_c] += 1
			}

			if letter == "D" {
				y_c++
				tda[y_c][x_c] += 1
			}

			if letter == "L" {
				x_c--
				tda[y_c][x_c] += 1
			}
		}
	}

	return tda
}

func ProcessData(first_path []string, second_path []string) ([20][20]int) {
	tda := [20][20]int{}
	tda[10][10] = 4

	tda = ProducePath(first_path, tda)
	tda = ProducePath(second_path, tda)
	
	return tda
}

func Draw(tda [20][20]int) {
	for  i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			fmt.Print(tda[i][j])
		}
		fmt.Print("\n")		
	}
}

func main() {
	Draw(ProcessData(ParseData(ReadData())))
	fmt.Println("_")
}