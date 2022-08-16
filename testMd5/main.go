package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := os.Open("/Users/wuzijie/Downloads/audit-blacklist-upgrade/归档.zip")
	content, _ := ioutil.ReadAll(f)
	h := md5.New()
	h.Write(content)
	toString := hex.EncodeToString(h.Sum(nil))
	fmt.Println(toString)
	//fileName := "d41d8cd98f00b204e9800998ecf8427e.zip"
	//s := strings.Split(fileName, path.Ext(fileName))[0]
	//fmt.Println(s)
}
