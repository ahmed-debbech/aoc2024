package main

import (
	"fmt"
	"os"
	"strconv"
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

	moveFrom := getRightMostBlock(disk)
	for i:=0; i<=len(disk)-1; i++{
		if (disk[i] == '.') && (disk[moveFrom] != '.') {
			disk[i] = disk[moveFrom]
			disk[moveFrom] = '.'
			moveFrom = getRightMostBlock(disk)
			if hasFinishMoving(disk) {return}
		}
	}
}

func getRightMostBlock(disk []byte) int{
	for i:=len(disk)-1; i>=0; i--{
		if(disk[i] >= '0') && (disk[i] <= '9'){
			return i
		}
	}
	return -1
}

func hasFinishMoving(disk []byte) bool{

	tolerate := true
	for i:=0; i<=len(disk)-1; i++{
		if disk[i] == '.' {
			tolerate = false
		}
		if((!tolerate) && ((disk[i] >= '0') && (disk[i] <= '9'))){
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
	return dat
}
