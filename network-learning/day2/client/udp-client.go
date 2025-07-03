package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// 连接到UDP服务器
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8090")
	if err != nil {
		fmt.Println("服务器地址解析失败:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer conn.Close()

	fmt.Println("已连接到UDP服务器 localhost:8090")
	fmt.Println("输入消息发送给服务器 (输入 'quit' 退出):")

	// 从标准输入读取消息
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()

		// 发送消息
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("发送消息失败:", err)
			continue
		}

		// 如果输入quit就退出
		if strings.ToLower(message) == "quit" {
			fmt.Println("正在关闭连接...")
			break
		}

		// 接收服务器回复
		buffer := make([]byte, 1024)
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("读取回复失败:", err)
			continue
		}

		response := string(buffer[:n])
		fmt.Printf("> %s\n", response)
	}

	fmt.Println("连接已关闭")
}
