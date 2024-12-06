package main 

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	fmt.Println("hello world")

	map_sit := readFile()

	strt_pos_x, strt_pos_y := getStartingPosition(map_sit) 
	fmt.Println("starting pos:", strt_pos_x, strt_pos_y)

	map_sit = loopup(map_sit, strt_pos_y, strt_pos_x)
	
	show(map_sit)

	count := unicVisits(map_sit)
	fmt.Println("TOTAL VISITED PLACES IS:", count)
}

func unicVisits(mapp []string) int{
	sum := 0
	for i, l := range mapp {
		for j:=0; j<=len(l)-1; j++{
			if mapp[i][j] == 'X'{
				sum++
			} 
		}
	}
	return sum
}

func loopup(mapp []string, x int, y int) []string{

	i := y
	j := x
	dir := 'u'
	for {
		fmt.Println("ind:",i,j)
		switch dir {
			case 'u':
				if i-1 < 0 {return mapp;}
				if mapp[i-1][j] == '#'{
					j++
					dir = 'r'
				}else{
					i--
				}
			case 'd':
				if i+1 > len(mapp)-1  {return mapp;}

				if mapp[i+1][j] == '#'{
					j--
					dir = 'l'
				}else{
					i++
				}
			case 'l':
				if j-1 < 0  {return mapp;}

				if mapp[i][j-1] == '#'{
					i--
					dir = 'u'
				}else{
					j--
				}
			case 'r':
				if j+1 > len(mapp[0])-1  {return mapp;}

				if mapp[i][j+1] == '#'{
					i++
					dir = 'd'
				}else{
					j++
				}
		}
		st := mapp[i]
		p := []byte(st)
		p[j] = 'X'
		mapp[i] = string(p)
	}
	return mapp
}

func show(mapp []string){
	for _, l := range mapp{
		fmt.Println(l)
	}
}



func getStartingPosition(mapp []string) (int, int){
	for kk, l := range mapp{
		for i:=0; i<=len(l)-1; i++{
			if l[i] == '^'{
				return kk, i
			}
		}
	}
	return -1,-1
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