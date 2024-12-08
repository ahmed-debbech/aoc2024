package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"math"
)

type Equation struct {
	Result int64
	Vals []int64
}

func main(){
	fmt.Println("hello world")

	lines := readFile()
	//fmt.Println(lines)

	equations := structureIt(lines)
	fmt.Println(equations)

	//now logic starts below

	all_calbrations := calculate(equations)
	fmt.Println(all_calbrations)

	var result int64 = 0
	for _, l := range all_calbrations{
		result += l
	}
	fmt.Println("final result IS:", result)
}


func calculate(eqs []Equation) []int64{
	all_calis := make([]int64, 0)

	for _, eq := range eqs{

		num_of_ops := len(eq.Vals)-1
		possible_combs := math.Pow(float64(2), float64(num_of_ops))
		fmt.Println("possible combinations:",possible_combs)
		fmt.Println("counting eq:", eq.Result )
		for i:=0;i<=int(possible_combs)-1; i++{
			rep := strconv.FormatInt(int64(i), 2)
			
			final_arr_op := make([]rune, 0)
			for j:=len(rep)-1; j>=0; j-- {
				if rep[j] == '0'{
					final_arr_op = append(final_arr_op, '+')
				}else{
					final_arr_op = append(final_arr_op, '*')
				}
			}

			for j:=len(rep); j<=num_of_ops-1; j++{
				final_arr_op = append(final_arr_op, '+')
			}

			fmt.Println(string(final_arr_op))
			
			k :=0
			s := eq.Vals[k]
			k++
			for j:=0; j<=len(final_arr_op)-1; j++{
				if final_arr_op[j] == '*'{
					s *= eq.Vals[k]
				}else{
					s += eq.Vals[k]
				}
				k++
			}
			fmt.Println("s: ", s)

			if s == eq.Result {
				all_calis = append(all_calis, s)
				break;
			}
		}
		fmt.Println("========")

		
		//fmt.Println(num_of_ops, bin)
	}

	return all_calis
}

func structureIt(lines []string) []Equation{
	equations := make([]Equation, 0)

	for _, l := range lines{
		
		v := strings.Split(l, ":")[0]
		nums := strings.Split(strings.Trim(strings.Split(l, ":")[1], " "), " ")
		res, _ := strconv.Atoi(v)
		var ii []int64
		for _, n := range nums{
			oo, _:=strconv.Atoi(n);
			ii = append(ii, int64(oo))
		}
		eq := Equation{
			Result : int64(res),
			Vals: ii,
		}
		equations = append(equations, eq)
	}
	return equations
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
