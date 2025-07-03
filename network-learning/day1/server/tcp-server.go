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

		// 如果收到"quit"就断开连接
		if strings.ToLower(message) == "quit" {
			fmt.Printf("客户端 %s 断开连接\n", conn.RemoteAddr())
			break
		}
	}
}
