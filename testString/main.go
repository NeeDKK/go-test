package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AllowListModbus struct {
	AllowListGroupId int    `json:"allowListGroupId"`
	DeviceName       string `json:"deviceName"`
	SrcAndDestIp     string `json:"srcAndDestIp"`
	SrcAndDestMac    string `json:"srcAndDestMac"`
	Func             string `json:"func"`
	StartAddr        string `json:"startAddr"`
	EndAddr          string `json:"endAddr"`
	Subfunc          string `json:"subfunc"`
}

func main() {
	//resp, err := http.Get("https://10.25.10.120:1443/state/getRegisterInfo")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(string(body))
	certPool := x509.NewCertPool()           // 初始化证书池
	caCrt, err := ioutil.ReadFile("pub.crt") // 读取ca.crt证书内容
	if err != nil {
		log.Fatalln("ca.crt read error:", err)
	}
	certPool.AppendCertsFromPEM(caCrt) // 解析ca证书，并添加到证书池
	// 解析颁发的客户端证书，LoadX509KeyPair由于没有提供解析时传入密码，所以客户端私钥不能有密码，
	// 否则会报错: tls: failed to parse private key
	clientCrt, err := tls.LoadX509KeyPair("pub.crt", "key.pem")
	if err != nil {
		log.Fatalln("client.crt LoadX509KeyPair error:", err)
	}
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:            certPool,                     // 设置双向认证ca证书
			Certificates:       []tls.Certificate{clientCrt}, // 客户端证书，需要传递给服务端
			InsecureSkipVerify: true,                         // 控制客户端是否验证服务器端的证书链和主机名，如果为true则客户端不验证服务器端的证书链和主机名
		},
	}
	// 如果InsecureSkipVerify为false，则可能会存在一个报错：x509: certificate is not valid for any names, but wanted to match www.api_domain.com
	client := &http.Client{Transport: transport}
	resp, err := client.Get("https://10.25.10.120:1443/state/getRegisterInfo")
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
