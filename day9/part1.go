package main

import (
	"fmt"
	"os"
	"strconv"
	_"strings"
	_"unicode/utf8"
)


func main(){
	fmt.Println("hello world")

	blocks := readFile()
	//fmt.Println(string(line))
	
	disk := convertToDiskLayout(blocks)
	fmt.Println("disk layout: ", string(disk))

	move(disk)
	fmt.Println("modified disk layout: ", string(disk))
	sum := checksum(disk)
	fmt.Println("FINAL RESULT IS:" , sum)
}

func checksum(disk []byte) uint64{
	var sum uint64 = 0
	for i :=0; i<=len(disk)-1; i++{
		if disk[i]=='.' {return sum}
		value := rune(disk[i]-'0')
		v := uint64(value)
		sum += (uint64(i)*v)
	}
	return sum
}

func move(disk []byte){

	for i:=0; i<=len(disk)-1; i++{
		moveFrom := getRightMostBlock(disk)
		if (rune(disk[i]) == '.') && (moveFrom != -1) {
			disk[i] = disk[moveFrom]
			disk[moveFrom] = byte('.')
			//fmt.Println(string(disk))
			if hasFinishMoving(disk) {return}
		}
	}
}

func getRightMostBlock(disk []byte) int{
	for i:=len(disk)-1; i>=0; i--{
		if(rune(disk[i]) >= '0') && (rune(disk[i]) <= '9'){
			return i
		}
	}
	return -1
}

func hasFinishMoving(disk []byte) bool{

	tolerate := true
	for i:=0; i<=len(disk)-1; i++{
		if rune(disk[i]) == '.' {
			tolerate = false
		}
		if((!tolerate) && ((rune(disk[i]) >= '0') && (rune(disk[i]) <= '9'))){
			return false
		}
	}
	return true
}

func convertToDiskLayout(blocks []byte) []byte{
	disk := make([]byte, 0)
	files := 0
	for i:=0; i<=len(blocks)-1; i++{
		if i % 2 == 0 {
			for j:=1; j<=int(blocks[i] - '0'); j++{
				disk = append(disk, byte(strconv.Itoa(files)[0]))
			}
			files++
		}else{
			//free space
			for j:=1; j<=int(blocks[i] - '0'); j++{
				disk = append(disk, byte('.'))
			}
		}
	}
	return disk
}

func readFile() []byte{
	dat, _ := os.ReadFile("input")
	dd := make([]byte, 0)
	for i:=0; i<=len(dat)-1; i++{
		if((rune(dat[i]) >= '0') && (rune(dat[i]) <= '9')) || (rune(dat[i]) == '.'){
			dd = append(dd, dat[i])
		}
	}
	//ss := (strings.Split(string(dd), "\n"))
	/*lines := (strings.Split(string(dat), "\n"))
	for _, str := range lines {
		dd = append(dd, []byte(str)...)
	}*/
	return dd
}
