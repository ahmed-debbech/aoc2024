package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main(){
	fmt.Println("Hello world")

	stones_str := readFile()
	fmt.Println(stones_str)

	stones_int := toInt(stones_str)
	//fmt.Printf("%T\n",stones_int)

	stones_str = toString(stones_int)
	//fmt.Printf("%T\n",stones_str)

	for i:=1; i<=10; i++{
		blink(&stones_str);
		fmt.Println("it", i, stones_str)
		//fmt.Println("len", len(stones_str))
		stones_int = toInt(stones_str)
		stones_str = toString(stones_int)
	}
	fmt.Println("FINAL NUMBER STONES:", len(stones_str))
}


func blink(stones_str *[]string){ 
	str := make([]string, 0)

	for  _, l := range *stones_str{
		if l == "0"{
			str = append(str, "1")
		}else{
			if len(l) % 2 == 0{
				str = append(str, l[:(len(l)/2)])
				str = append(str, l[(len(l)/2):])
			}else{
				ii, _ := strconv.Atoi(l)
				ii = ii * 2024
				str = append(str, strconv.Itoa(ii))
			}
		}
	}
	*stones_str= str
}

//func to string
func toString(stones_int []int64) []string{
	strr := make([]string, 0)
	
	for _, x := range stones_int {
		s := strconv.Itoa(int(x))
		strr = append(strr, s)
	}

	return strr
}
//func to int
func toInt(stones_str []string) []int64{
	intt := make([]int64, 0)
	
	for _, x := range stones_str {
		s, _ := strconv.Atoi(x)
		intt = append(intt, int64(s))
	}

	return intt
}

func readFile() []string{
	data, _ := os.ReadFile("input")
	return strings.Split(string(data), " ")
}