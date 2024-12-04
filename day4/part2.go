package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	fmt.Println("Hello")

	lines := readFile()
	//fmt.Println(lines)
	canvas := createCanvasAround(lines)
	//fmt.Println(canvas)

	count := searchXmas(canvas)
	fmt.Println("COUNT OF XMAS IS:", count)
}


func dump(canvas [][]byte, i int, j int) int{
	manyM := 0
	manyS := 0
	if canvas[i][j] == 'A' {
		if canvas[i-1][j-1] != 'S' && canvas[i-1][j-1] != 'M' {return 0}
		if canvas[i-1][j-1] == 'M' {manyM++}else{manyS++}

		if canvas[i-1][j+1] != 'S' && canvas[i-1][j+1] != 'M' {return 0}
		if canvas[i-1][j+1] == 'M' {manyM++}else{manyS++}

		if canvas[i+1][j+1] != 'S' && canvas[i+1][j+1] != 'M' {return 0}
		if canvas[i+1][j+1] == 'M' {manyM++}else{manyS++}

		if canvas[i+1][j-1] != 'S' && canvas[i+1][j-1] != 'M' {return 0}
		if canvas[i+1][j-1] == 'M' {manyM++}else{manyS++}

		if (manyM != 2) || (manyS != 2) {return 0}
		fmt.Println("i: ", i, " , j:",j )
		return 1
	}

	return 0
}

func searchXmas(canvas [][]byte) int{
	count := 0;

	for i:=3; i<=len(canvas)-4; i++ {
		for j:=3; j<=len(canvas[i])-4; j++ {
			fmt.Println("-----------[",i,",",j,"]");
			count += dump(canvas, i, j)
			fmt.Println("=============");
			//fmt.Printf("%c",canvas[i][j])
		}
		//fmt.Println()
	}
	return count
}

func createCanvasAround(lines []string) [][]byte{
	canvas := make([][]byte, len(lines)+6)
	
	by := make([]byte, len(lines[0])+6)
	for i:=0; i<=len(lines[0])+5; i++{
		by[i] = '.'
	}

	canvas[0] = by
	canvas[1] = by
	canvas[2] = by
	canvas[len(lines[0])+3] = by
	canvas[len(lines[0])+4] = by
	canvas[len(lines[0])+5] = by

	i := 3
	for _, str := range lines{
		by := make([]byte, len(str)+6)
		by[0]='.';by[1]='.';by[2]='.';
		k := 3;
		for j:=0; j<=len(str)-1; j++ {
			if str[j] == 'X' {
				by[k] = '.'
			}else{
				by[k] = str[j]
			}
			k++
		}
		by[len(str)+3]='.';by[len(str)+4]='.';by[len(str)+5]='.';
		canvas[i] = by
		i++
	}

	return canvas
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