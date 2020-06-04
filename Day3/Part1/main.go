package main

import ( 
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"regexp"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ReadData() string {
	data, err := ioutil.ReadFile("data")

	if err != nil {
		os.Exit(-1)
	}

	return string(data)
}

func ParseData(data string) ([]string, []string) {
	wires_path := strings.Split(data, "\n")

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

func FindLowestDistance(first_path_points [][]int, second_path_points [][]int) (int) {
	lowest_distance := 0

	for i := 0; i < len(first_path_points); i++ {
		for j := 0; j < len(second_path_points); j++ {
			if (first_path_points[i][0] == second_path_points[j][0] && first_path_points[i][1] == second_path_points[j][1] ) {
				distance := Abs(first_path_points[i][0]) + Abs(first_path_points[i][1])
				if (distance < lowest_distance || lowest_distance == 0) {
					lowest_distance = distance
				}
			}
		}
	}

	return lowest_distance
}

func GetPoints(path []string) ([][]int) {
	var points [][]int

	x := 0
	y := 0
	drx := map[string]int{ "U": 0, "R": 1, "D": 0, "L": -1, }
	dry := map[string]int{ "U": 1, "R": 0, "D": -1, "L": 0, }
	for i := 0; i < len(path); i++ {
		letter, number := ParseMove(path[i])

		for n := 0; n < number; n++ {
			x += drx[letter]
			y += dry[letter]
			points = append(points, []int{x,y})
		}
	}

	return points
}

func main() {
	data_file				:= ReadData()
	first_path, second_path := ParseData(data_file)
	first_path_points		:= GetPoints(first_path)
	second_path_points		:= GetPoints(second_path)
	lowest_distance			:= FindLowestDistance(first_path_points, second_path_points)

	fmt.Println(lowest_distance)
}