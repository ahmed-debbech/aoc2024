package main

import (
	"fmt"
	"os"
	"bufio"
)

type Machine struct{
	XIncrement int
	YIncrement int
	PrizeX int
	PrizeY int
}

func main(){
	fmt.Println("Hello")

	lines := readFile()
	fmt.Println(lines)

	machines := getMachines(lines)
	fmt.Println(machines)
}

func getMachines(lines []string) []Machine{

	mchns := make([]Machine, 0)

	for i:=0; i<=len(lines)-1; i++{
		if lines[i] == "" {continue}
		
		for j:=0; j<=len(lines[i])-1; j++{
			mchns = append(mchns,  lines[i])
		}
	}
	return mchns
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
