package main

import (
	"fmt"
	"os"
	"bufio"
)

func getMul(lines []string) []string{
	var st []string
	for _, val := range lines {
		for i:=0; i<=len(val)-12; i++{
			if val[i:i+4] == "mul(" {
				fmt.Println(val[i:i+12])
				st = append(st, val[i:i+12])
			}
		}
		fmt.Println(st)
	}
	return st
}

func verifyMuls(uncheckedMuls []string) []string{
	var st []string
	for _, vl := range uncheckedMuls {
		if (vl[5] >= '0') && (vl[5] <= '9') {
			if 
		}
	}
	return st
}

func main(){
	file, _ := os.Open("input")
    defer file.Close()

    scanner := bufio.NewScanner(file)
	
	var lines []string
    for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//fmt.Println(lines[0])

	muls := getMul(lines)
	muls = verifyMuls(muls)
	fmt.Println(muls)
}