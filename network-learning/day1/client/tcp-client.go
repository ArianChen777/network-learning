package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	// 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8089")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer conn.Close()

	fmt.Println("已连接到服务器 localhost:8089")
	fmt.Println("输入消息发送给服务器 (输入 'quit' 退出):")

	// 使用WaitGroup等待接收goroutine完成
	var wg sync.WaitGroup
	wg.Add(1)

	// 启动接收goroutine
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Printf("> %s\n", scanner.Text())
		}
		fmt.Println("接收goroutine结束")
	}()

	// 从标准输入读取消息
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()

		// 发送消息
		fmt.Fprintf(conn, "%s\n", message)

		// 如果输入quit就开始优雅关闭
		if strings.ToLower(message) == "quit" {
			fmt.Println("正在关闭连接...")

			// 等待一小段时间确保服务器收到quit消息并回复
			time.Sleep(50 * time.Millisecond)

			// 半关闭连接：关闭写端，但保持读端开放
			// 这样客户端先发送FIN，服务器可以发送剩余数据后再关闭
			if tcpConn, ok := conn.(*net.TCPConn); ok {
				tcpConn.CloseWrite()
				fmt.Println("已关闭写端（发送FIN），等待服务器关闭连接...")
			}

			break
		}
	}

	// 等待接收goroutine完成
	wg.Wait()
	fmt.Println("连接已完全关闭")
}
