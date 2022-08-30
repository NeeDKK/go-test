package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Test1 interface {
	Func1()
}
type People struct {
	Name string
}

type Test2 interface {
	Func2()
}

func main() {
	a := "SN=4ZxQX80s3qkHllngZ52GmOHSAbfb0csODz89ZsSvvC+X45o01bjAT3Lex0kLzfwgBCCLFxxFIck=#prod=IPCDefSys;project=IPCDefSys001;pid=;vid=;sku=;KillTool=1;ExternalDevice=1;Whitelist=1;AppGuard=1;apply_date=1659057092;expire_date=1690593092;effective_time=31536000;type=2;uuid=c9f97248-d887-4252-a5e9-7bf1518f1acf;#signature=cbf219979b0b7d96ff9430b717e61f86440fd25a1ea277a71244f6d817d2bd7d"
	split := strings.Split(a, "#")
	requestCode := strings.Split(split[0], "SN=")[1]
	typeString := strings.Split(split[1], "type=")[1]
	typeString = typeString[:1]
	num := ""
	if strings.Contains(split[1], "Whitelist=1;") {
		num += "1"
	} else {
		num += "0"
	}
	if strings.Contains(split[1], "ExternalDevice=1;") {
		num += "1"
	} else {
		num += "0"
	}
	if strings.Contains(split[1], "KillTool=1;") {
		num += "1"
	} else {
		num += "0"
	}
	if strings.Contains(split[1], "AppGuard=1;") {
		num += "1"
	} else {
		num += "0"
	}
	fmt.Println(requestCode)
	fmt.Println(typeString)
	fmt.Println(num)
	c := getInput("0100")
	out := sq(c)
	sum := 0
	for o := range out {
		// fmt.Println(o)
		sum += o
	}
	timeStr := strings.Split(split[1], "effective_time=")[1]
	timeUsestr, _ := strconv.Atoi(strings.Split(timeStr, ";")[0])
	timeUse := timeUsestr / 3600 / 24
	commondstr := "/go/src/netvine/server/generate " + requestCode + " " + typeString + " " + strconv.Itoa(sum) + " " + strconv.Itoa(timeUse)
	fmt.Println(sum)
	fmt.Println(commondstr)
}

func StringToIntArray(input string) []int {
	output := []int{}
	for _, v := range input {
		output = append(output, int(v))
	}
	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}
	return output
}

func getInput(input string) <-chan int {
	out := make(chan int)
	go func() {
		for _, b := range StringToIntArray(input) {
			out <- b
		}
		close(out)
	}()

	return out
}
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	var base, i float64 = 2, 0
	go func() {
		for n := range in {
			out <- (n - 48) * int(math.Pow(base, i))
			i++
		}
		close(out)
	}()
	return out
}

func Test(interface{}) {
	fmt.Println("Test========interface")

}

func (a People) Func1() {
	fmt.Println("Func1========people")
}

func (a People) Func2() {
	fmt.Println("Func2========people")

}
