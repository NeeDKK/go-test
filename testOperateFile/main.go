package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	open, _ := os.Open("device_info.conf")
	defer open.Close()
	all, _ := ioutil.ReadAll(open)
	split := strings.Split(string(all), "\n")
	i := strings.Split(split[0], "sn=")
	fmt.Println(i[1])
	os.Create("aaa-zzz.json")
}
