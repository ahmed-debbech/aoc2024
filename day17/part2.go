package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
	"sync"
)


var (
	ops_names = []string{"adv", "bxl", "bst", "jnz", "bxc", "out", "bdv", "cdv"}
)

func main(){

	var wg sync.WaitGroup
	wg.Add(3)

	thea := 101000000
	go work(thea, &wg)
	thea = 500000000
	go work(thea, &wg)
	thea = 1000000000
	go work(thea, &wg)

	wg.Wait()

}

func work(thea int,  wg *sync.WaitGroup) {
	defer wg.Done()

	lines := readFile()
	a,b,c, ops := makeItReal(lines)

	cpu := Cpu{}

	for  {
		a = thea
		cpu.Init(a,b,c,ops)

		//fmt.Println(cpu)

		for cpu.Execute() {
			//cpu.Dump()
		} 

		fmt.Println("SEQ: ", thea)
		if exact(cpu.Log, intSliceToStringSlice(ops)) {break}
		thea++
	}

	fmt.Println("RESULT IS", thea)

}

func intSliceToStringSlice(intSlice []int) []string {
	var stringSlice []string
	for _, num := range intSlice {
		stringSlice = append(stringSlice, strconv.Itoa(num))
	}
	return stringSlice
}

func exact(a []string, b []string) bool{
	
	if len(a) != len(b) {return false}

	for k, c := range a {
		if c != b[k] {
			return false
		}
	}
	return true
}

type Cpu struct {
	A int
	B int
	C int 
	Pointer int
	PIncr int 
	Ops []int
	Log []string
}

func (cpu *Cpu) Dump(){
	fmt.Println("A:", cpu.A)
	fmt.Println("B:", cpu.B)
	fmt.Println("C:", cpu.C)
}
func (cpu *Cpu) Init(a int, b int, c int, ops []int){
	cpu.A = a
	cpu.B = b
	cpu.C = c
	cpu.Ops = ops
	cpu.Pointer = 0
	cpu.PIncr = 2
	cpu.Log = make([]string, 0)
}
func (cpu *Cpu) Execute() bool{
	if (cpu.Pointer >= len(cpu.Ops))  {return false}
	
	operation := cpu.Ops[cpu.Pointer]
	operand := cpu.Ops[cpu.Pointer+1]
	
	cpu.Do(operation, operand)

	if cpu.PIncr == 0 {
		cpu.PIncr = 2
		return true
	}
	cpu.Pointer += cpu.PIncr
	return true
}

func (cpu *Cpu) Do(operation int, operand int){
	//fmt.Println(operation, operand)
	unit := cpu.KnowOperation(operation)
	unit(operation, operand)
}

func (cpu *Cpu) KnowOperation(operation int) func(int, int) {
	ss := ops_names[operation]

	switch(ss){
	case "adv":
		return func(op int, oper int){
			oper = cpu.setOperandRealValue(oper)
			//fmt.Println("adv ecuteing")
			res := (cpu.A / int(math.Pow(2, float64(oper))))
			cpu.A = int(res)
		}
	case "bxl":
		return func(op int, oper int){
			//fmt.Println("bxl ecuteing")
			res := (cpu.B ^ oper)
			cpu.B = res
		}
	case "bst":
		return func(op int, oper int){
			oper = cpu.setOperandRealValue(oper)
			//fmt.Println("bst ecuteing")
			g := oper % 8
			res := 7 & g
			cpu.B = res
		}
	case "jnz":
		return func(op int, oper int){
			//fmt.Println("jnz ecuteing")
			if cpu.A == 0 {return}
			cpu.Pointer = oper
			cpu.PIncr = 0
		}
	case "bxc":
		return func(op int, oper int){
			//fmt.Println("bxc ecuteing")
			res := (cpu.B ^ cpu.C)
			cpu.B = res
		}
	case "out":
		return func(op int, oper int){
			//fmt.Println("out ecuteing")
			oper = cpu.setOperandRealValue(oper)
			res := oper % 8
			fop := strconv.Itoa(res)
			cpu.Log = append(cpu.Log, (fop))
		}
	case "bdv":
		return func(op int, oper int){
			//fmt.Println("bdv ecuteing")
			oper = cpu.setOperandRealValue(oper)
			//fmt.Println("bdv ecuteing")
			res := (cpu.A / int(math.Pow(2, float64(oper))))
			cpu.B = int(res)
		}
	case "cdv":
		return func(op int, oper int){
			//fmt.Println("cdv ecuteing")
			oper = cpu.setOperandRealValue(oper)
			//fmt.Println("cdv ecuteing")
			res := (cpu.A / int(math.Pow(2, float64(oper))))
			cpu.C = int(res)
		}
	}
	return nil
}

func (cpu *Cpu) setOperandRealValue(oper int) int{
	if (oper <= 3) && (oper >= 0) {
		return oper
	}
	if oper == 4 {
		return cpu.A
	}
	if oper == 5 {
		return cpu.B
	}
	if oper == 6 {
		return cpu.C
	}
	return -1
}


func makeItReal(lines []string) (int,int,int, []int){
	a, _ := strconv.Atoi(lines[0])
	b, _ := strconv.Atoi(lines[1])
	c, _ := strconv.Atoi(lines[2])

	ops := make([]int, 0)
	for _ , x := range strings.Split(lines[3], ",") {
		xx , _ := strconv.Atoi(x)
		ops = append(ops, xx)
	}

	return a,b,c, ops
}

func readFile() []string{
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	scanner.Scan()
	lines = append(lines, strings.Trim(strings.Split(scanner.Text(), ":")[1], " "))
	scanner.Scan()
	lines = append(lines, strings.Trim(strings.Split(scanner.Text(), ":")[1], " "))
	scanner.Scan()
	lines = append(lines, strings.Trim(strings.Split(scanner.Text(), ":")[1]," "))
	scanner.Scan()
	scanner.Scan()
	lines = append(lines, strings.Trim(strings.Split(scanner.Text(), ":")[1], " "))

	//fmt.Println(lines)
	return lines
}
