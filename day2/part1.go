package main;


import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"math"
	"strconv"
)

func main(){

	fmt.Println("Hello world")

	readFile, err := os.Open("input")

    if err != nil {
        //fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string

    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
		//fmt.Println(fileLines)
	}

	var levelsList [][]int
	for _, ele := range fileLines {
		s := strings.Split(ele, " ")
		k := make([]int, 0)
		for _, ss := range s {
			ee, _ := strconv.Atoi(ss);
			k = append(k, ee)
		}
		levelsList = append(levelsList, k)
	}
	fmt.Println(levelsList)

	safeCount := 0;
	for _, lvl := range levelsList {
		decreasing := false
		if lvl[0] > lvl[1] {
			decreasing = true
		}
		fmt.Println("decreasing?", decreasing)
		safe := inTheCorrectWay(lvl, 0, decreasing)
		fmt.Println("is safe :", safe)
		if safe {safeCount++}
	}
	fmt.Println("SAFE COUNT: ", safeCount)
}

func inTheCorrectWay(lvl []int, i int, decreasing bool) bool {
	if i == len(lvl)-1{
		return true
	}
	if decreasing {
		if (lvl[i] > lvl[i+1]) && ((math.Abs(float64(lvl[i]) - float64(lvl[i+1])) <= 3) && (math.Abs(float64(lvl[i]) - float64(lvl[i+1])) >= 1)) {
			return inTheCorrectWay(lvl, i+1, decreasing)
		}
		return false	
	}else{
		if (lvl[i] < lvl[i+1]) && ((math.Abs(float64(lvl[i]) - float64(lvl[i+1])) <= 3) && (math.Abs(float64(lvl[i]) - float64(lvl[i+1])) >= 1)) {
			return inTheCorrectWay(lvl, i+1, decreasing)
		}
		return false
	}
}



