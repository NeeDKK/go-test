package main

import (
	"sync"
	"test/socket/utils"
)

var (
	WaitGroup sync.WaitGroup
	sql       = "'2022-08-11 14:24:35','2022-08-11 14:24:35','a0:36:9f:65:32:99','00:00:00:11:4c:02', 4,'128.128.2.134',6168,'128.128.2.4',4369,17,'UDP',57, 000000114C02A0369F65329908004500002B1879000080110000808002868080020418181111001705B301AB0000C6001A00010040000100FF,'scnet2'----bit_type,message_number,bit_field_describe,value----'0xab','198','TT_12405 产品气缓冲罐入口温度','0.000000'----"
	sql1      = "10.25.16.208,28:d0:ea:5f:4c:a4,10.25.17.129,a8:93:4a:04:7d:25,modbus"
	whitelist = "('192.168.0.211','192.168.0.222',80,8080,'TCP','modbus',78,'whitelist reject!','{protocol:modbus,func:3,startaddr:1001,endaddr:1001}','2022-05-16 11:03:28','AAEE','080027ffee23','08002767053c','')"
)

//socketGroup.TrafficSocket = socket.NewUnixSocket("/data/socket/triffic_sock", 1024)
//socketGroup.AllowListSocket = socket.NewUnixSocket("/data/socket/allow_list_sock", 1024)
//socketGroup.ObtainProtoSocket = socket.NewUnixSocket("/data/socket/obtain_proto_sock", 1024)
//socketGroup.EventSocket = socket.NewUnixSocket("/data/socket/event_sock", 1024)
//socketGroup.UnknownDeviceSocket = socket.NewUnixSocket("/data/socket/unknown_device_sock", 1024)
//socketGroup.WhiteListWarningSocket = socket.NewUnixSocket("/data/socket/white_list_warning_sock", 1024)
//socketGroup.SessionSocket = socket.NewUnixSocket("/data/socket/session_statistic_sock", 1024)

func main() {
	//声明unixsocket
	us := utils.NewUnixSocket("/data/socket/obtain_proto_sock")
	//发送数据unixsocket并返回服务端处理结果
	//for z := 0; z < 1; z++ {
	//	WaitGroup.Add(1)
	//	go func() {
	//		for i := 0; i < 1; i++ {
	us.ClientSendContext(sql)
	//time.Sleep(time.Second * 2)
	//		}
	//		defer WaitGroup.Done()
	//	}()
	//}
	//WaitGroup.Wait()
}
