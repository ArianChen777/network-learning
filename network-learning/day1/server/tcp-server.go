package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// 监听8089端口
	listener, err := net.Listen("tcp", ":8089")
	if err != nil {
		fmt.Println("监听失败:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP服务器启动，监听端口 :8089")
	fmt.Println("等待客户端连接...")

	for {
		// 接受连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受连接失败:", err)
			continue
		}

		fmt.Printf("新客户端连接: %s\n", conn.RemoteAddr())

		// 处理连接
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("收到消息: %s\n", message)

		// 回显消息
		response := fmt.Sprintf("服务器回复: %s\n", strings.ToUpper(message))
		conn.Write([]byte(response))

		// 如果收到"quit"，发送回复但不立即断开，等待客户端关闭连接
		if strings.ToLower(message) == "quit" {
			fmt.Printf("收到quit消息，等待客户端关闭连接: %s\n", conn.RemoteAddr())
			// 不break，让连接继续，直到客户端关闭连接时scanner.Scan()返回false
		}
	}
	
	// 当scanner.Scan()返回false时（客户端关闭连接），程序到达这里
	fmt.Printf("客户端连接已关闭: %s\n", conn.RemoteAddr())
}
