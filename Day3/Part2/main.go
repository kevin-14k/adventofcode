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

func FindLowestDistance(first_path_points map[string]int, second_path_points map[string]int) (int) {
	lowest_distance := 0
	var keys1 []string
	var keys2 []string

	for k := range first_path_points { keys1 = append(keys1, k) }
	for k := range second_path_points { keys2 = append(keys2, k) }

	for i := 0; i < len(keys1); i++ {
		for j := 0; j < len(keys2); j++ {
			if (keys1[i] == keys2[j]) {
				distance := first_path_points[keys1[i]] + second_path_points[keys2[j]]
				
				if (distance < lowest_distance || lowest_distance == 0) {
					lowest_distance = distance
				}
			}
		}
	}

	return lowest_distance
}

func GetPoints(path []string) (map[string]int) {
	points := make(map[string]int)

	x := 0
	y := 0
	l := 0
	drx := map[string]int{ "U": 0, "R": 1, "D": 0, "L": -1, }
	dry := map[string]int{ "U": 1, "R": 0, "D": -1, "L": 0, }
	for i := 0; i < len(path); i++ {
		letter, number := ParseMove(path[i])

		for n := 0; n < number; n++ {
			x += drx[letter]
			y += dry[letter]
			key := strconv.Itoa(x) + "." + strconv.Itoa(y)
			l++
			points[key] = l
		}
	}

	return points
}

func main() {
	data_file				:= ReadData()
	first_path, second_path := ParseData(data_file)
	first_path_points		:= GetPoints(first_path)
	second_path_points		:= GetPoints(second_path)
	lowest_distance         := FindLowestDistance(first_path_points, second_path_points)

	fmt.Println(lowest_distance)
}