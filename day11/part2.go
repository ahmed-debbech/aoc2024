package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"time"
)

func main(){
	fmt.Println("Hello world")

	stones_str := readFile()
	fmt.Println(stones_str)

	stones_int := toInt(stones_str)
	//fmt.Printf("%T\n",stones_int)

	stones_str = toString(stones_int)
	//fmt.Printf("%T\n",stones_str)

	//for i:=1; i<=75; i++{
		stones_str = []string{"0"}
		stones_str = blink(stones_str, 1);
		fmt.Println("it", 1, stones_str)
		//fmt.Println("len", len(stones_str))
		stones_int = toInt(stones_str)
		stones_str = toString(stones_int)
	//}
	fmt.Println("FINAL NUMBER STONES:", len(stones_str))
}

/*func blnk(str *[]string ,stones_str *[]string, i int){
	if i<0 {
		return
	}
	if (*stones_str)[i] == "0"{
		*str = append(*str, "1")
	}else{
		if len((*stones_str)[i]) % 2 == 0{
			*str = append(*str, (*stones_str)[i][:(len((*stones_str)[i])/2)])
			*str = append(*str,(*stones_str)[i][(len((*stones_str)[i])/2):])
		}else{
			ii, _ := strconv.Atoi((*stones_str)[i])
			ii = ii * 2024
			*str = append(*str, strconv.Itoa(ii))
		}
	}
	blnk(str, stones_str,i-1)

}*/

func blink(stones_str []string, kk int) []string { 

	if kk > 75 {return stones_str}

	str := make([]string, 0)

	for _, l := range stones_str{
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
	fmt.Println("it", kk , stones_str)
	time.Sleep(time.Second * 2)
	return blink(str, kk+1)
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