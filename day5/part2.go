package main 

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"math"
)

func main(){
	fmt.Println("Hello")

	lines := readFile()
	rules, updates := extract(lines)
	fmt.Println(updates)
	beforeMap, afterMap := prepare(rules)
	fmt.Println(beforeMap,".......", afterMap)

	count := 0;
	for i:=0; i<=len(updates)-1; i++{
		_, up := isOrdered(beforeMap, afterMap, updates[i]);		
		if len(up) > 0 {
			fmt.Println("Correction:", up)
			middle, _ := strconv.Atoi(up[int(math.Round(float64(len(up)/2)))])
			count += middle
		}	
	}
	fmt.Println(count)
}

func isOrdered(mbefore map[string][]string, mafter map[string][]string, updates []string) (int, []string){

	mod:=false
	for i:=0; i<=len(updates)-1; i++ {

		//fmt.Println("number : ", updates[i])
		//before
		for j:=i-1; j>=0; j-- {
			//fmt.Println("looking for", updates[j], "against map after")
			nums := mafter[updates[j]]
			//fmt.Println("array from map before", nums)
			if nums != nil {
				for _, a:=range nums{
					if a == updates[i]{
						xx := updates[i]
						updates[i] = updates[j]
						updates[j] = xx
						mod = true
						//fmt.Println("correction: ", updates)
					}
				}
			}
		}

		//after
		//fmt.Println("after")
		for j:=i+1; j<=len(updates)-1; j++{
			//fmt.Println("looking for", updates[j], "against map before")
			nums := mbefore[updates[j]]
			//fmt.Println("array from map before", nums)
			if nums != nil {
				for _, a:=range nums{
					if a == updates[i]{
						xx := updates[i]
						updates[i] = updates[j]
						updates[j] = xx
						mod = true
						//fmt.Println("correction: ", updates)
					}
				}
			}
		}
	}
	middle, _ := strconv.Atoi(updates[int(math.Round(float64(len(updates)/2)))])
	//fmt.Println("found element that should be after", middle)
	if mod {
		return middle, updates
	}
	return middle, []string{}

}

func prepare(rules []string) (map[string][]string, map[string][]string){
	ma := make(map[string][]string)
	mb := make(map[string][]string)
	for _, rule := range rules{
		d := strings.Split(rule, "|")
		mb[d[0]] = append(mb[d[0]], d[1]);
	}
	for _, rule := range rules{
		d := strings.Split(rule, "|")
		ma[d[1]] = append(ma[d[1]], d[0]);
	}
	return mb, ma
}

func extract(lines []string) ([]string, [][]string){
	var rules []string 
	var updates []string 
	var spltdd []string
	var spltd [][]string

	is_updates := false
	for _, l := range lines {

		if l == "" {
			is_updates = true
			continue;
		}
		if is_updates {
			updates = append(updates, strings.Split(l,",")...)
			spltdd = append(spltdd, updates...)
			spltd = append(spltd, spltdd)
			updates = make([]string, 0)
			spltdd = make([]string, 0)
		}else{
			rules = append(rules, l)
		}
	}
	return rules, spltd
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