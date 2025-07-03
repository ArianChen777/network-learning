package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听UDP端口8090
	addr, err := net.ResolveUDPAddr("udp", ":8090")
	if err != nil {
		fmt.Println("地址解析失败:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("监听失败:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP服务器启动，监听端口 :8090")
	fmt.Println("等待客户端消息...")

	buffer := make([]byte, 1024)

	for {
		// 接收数据
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("读取数据失败:", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Printf("收到来自 %s 的消息: %s\n", clientAddr, message)

		// 回显消息
		response := fmt.Sprintf("服务器回复: %s", message)
		_, err = conn.WriteToUDP([]byte(response), clientAddr)
		if err != nil {
			fmt.Println("发送回复失败:", err)
			continue
		}

		// 如果收到"quit"，记录日志但继续服务
		if message == "quit" {
			fmt.Printf("收到quit消息，来自客户端: %s\n", clientAddr)
		}
	}
}