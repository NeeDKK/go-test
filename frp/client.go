package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

var (
	// 本地需要暴露的服务端口
	localServerAddr = "127.0.0.1:8080"

	remoteIP = "8.210.42.150"
	// 远端的服务控制通道，用来传递控制信息，如出现新连接和心跳
	remoteControlAddr = remoteIP + ":8009"
	// 远端服务端口，用来建立隧道
	remoteServerAddr = remoteIP + ":8008"
)

func main() {
	tcpConn, err := CreateTCPConn(remoteControlAddr)
	if err != nil {
		log.Println("[连接失败]" + remoteControlAddr + err.Error())
		return
	}
	log.Println("[已连接]" + remoteControlAddr)

	reader := bufio.NewReader(tcpConn)
	for {
		s, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}

		// 当有新连接信号出现时，新建一个tcp连接
		if s == NewConnection+"\n" {
			go connectLocalAndRemote()
		}
	}

	log.Println("[已断开]" + remoteControlAddr)
}

func connectLocalAndRemote() {
	local := connectLocal()
	remote := connectRemote()

	if local != nil && remote != nil {
		Join2Conn(local, remote)
	} else {
		if local != nil {
			_ = local.Close()
		}
		if remote != nil {
			_ = remote.Close()
		}
	}
}

func connectLocal() *net.TCPConn {
	conn, err := CreateTCPConn(localServerAddr)
	if err != nil {
		log.Println("[连接本地服务失败]" + err.Error())
	}
	return conn
}

func connectRemote() *net.TCPConn {
	conn, err := CreateTCPConn(remoteServerAddr)
	if err != nil {
		log.Println("[连接远端服务失败]" + err.Error())
	}
	return conn
}