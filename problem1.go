package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func removeDupe(slc []int) []int {
    var tmpSlice []int
    chkMap := make(map[int]bool)

    for _, val := range slc {
        chkMap[val] = true
    }

    for k, _ := range chkMap {
        tmpSlice = append(tmpSlice, k)
    }
    sort.Ints(tmpSlice)
    return tmpSlice
}

func main() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	maxCal := 0
	curElfCalories := 0
	input := make([]int, 1)

	for scan.Scan(){
		food, err := strconv.Atoi(scan.Text())
		curElfCalories += food

		input = append(input, curElfCalories)
		if err != nil{
			if curElfCalories>maxCal{
				maxCal = curElfCalories
			}
			curElfCalories = 0
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(input)))
	unique := removeDupe(input)
	fmt.Println(maxCal)
	first:= unique[len(unique)-1]
	second:= unique[len(unique)-2]
	third:= unique[len(unique)-3]
	fmt.Println(first+second+third)
}
