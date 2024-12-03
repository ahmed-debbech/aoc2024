package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strings"
)

func getMul(lines []string) []string{
	var st []string
	for _, val := range lines {
		for i:=0; i<=len(val)-4; i++{
			if val[i:i+4] == "mul(" {
				//fmt.Println(val[i:i+12])
				if i+12 > len(val)-1 {
					st = append(st, val[i:len(val)-1])
				}else{
					st = append(st, val[i:i+12])
				}
			}
		}
		//fmt.Println(st)
	}
	return st
}

func verifyMuls(uncheckedMuls []string) []string{
	var st []string
	for _, vl := range uncheckedMuls {
		match, _ := regexp.MatchString(`mul\((\d{1,3}),(\d{1,3})\)`, vl)
		if match {
			st = append(st, vl)
		}
	}
	
	//trim
	for i:=0; i<=len(st)-1; i++ {
		st[i] = st[i][:strings.Index(st[i], ")")+1]
		fmt.Println(st[i])

	}
	return st
}

func countMuls(muls []string) int64{
	
	var count int64 = 0
	for _, vl := range muls {
		var left int64
		var right int64
		fmt.Sscanf(vl, "mul(%d,%d)", &left, &right)
		fmt.Println("ll : ", left)
		count += (left * right)
	}
	return count
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
	countFinal := countMuls(muls)
	fmt.Println("THE FINAL COUNT OF MULS IS:",countFinal)
}