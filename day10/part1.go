package main

import (
	"fmt"
	"os"
	"bufio"
	_"time"
)

type Trail struct{
	X int
	Y int
}

var (
	max_line int
	max_col int	
)

func main(){
	fmt.Println("hello world")

	map_trail := readFile()
	fmt.Println(map_trail)

	max_line = len(map_trail)
	max_col = len(map_trail[0])
	fmt.Println("max col", max_line, max_col)
	trails := getTrailHeads(map_trail)
	fmt.Println(trails)
	
	var p []string
	copy(p, map_trail)
	countSingleTrail(trails[0].Y, trails[0].X, map_trail, p)
}

func countSingleTrail(i int, j int, map_trail []string, p []string){
	fmt.Println("counting trail in:", i, j)
	if ((i >= max_line) || (i < 0)) || ((j >= max_col) || (j < 0) ) || (p[i],) {
		fmt.Println("discarding.. ",i, j)
		return
	}else{
		p = append(p, '.')
	}
	countSingleTrail(i-1, j, map_trail, p) // up
	countSingleTrail(i, j+1, map_trail, p) // right
	countSingleTrail(i+1, j, map_trail, p) // down
	countSingleTrail(i, j-1, map_trail, p) // left
	//time.Sleep(2*time.Second)
}

func getTrailHeads(lines []string) []Trail{
	trails := make([]Trail, 0)

	for i:=0; i<=len(lines)-1; i++{
		for j:=0; j<=len(lines[i])-1; j++{
			if lines[i][j] == '0' {
				trails = append(trails, Trail{X:j, Y:i,})		
			}
		}
	}
	return trails
}

func readFile() []string{
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//fmt.Println(lines)
	return lines
}