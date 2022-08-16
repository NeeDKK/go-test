package main

import (
	"bytes"
	"fmt"
	"github.com/ErmaiSoft/GoOpenXml/word"
	"io/ioutil"
	"os"
	"text/template"
)

type AAAA struct {
	Aaa string `json:"aaa"`
}

var (
	base64String = "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"
)

func main() {
	reportData := AAAA{
		Aaa: base64String,
	}
	open, _ := os.Open("/Users/wuzijie/Downloads/Doc1(1).xml")
	templateContent, err := ioutil.ReadAll(open)
	if err != nil {
		fmt.Println(err.Error())
	}
	reportPath := "b.xml"
	temp := template.New(reportPath)
	parse, err := temp.Parse(string(templateContent))
	if err != nil {
		fmt.Println(err.Error())
	}
	outputBuf := new(bytes.Buffer)
	err = parse.Execute(outputBuf, &reportData)
	if err != nil {
		fmt.Println(err.Error())
	}
	docx := word.CreateDocx()
	docx.WriteToFile("c.docx")
	err = ioutil.WriteFile("c.doc", outputBuf.Bytes(), 0777)
	if err != nil {
		fmt.Println(err.Error())
	}
}
