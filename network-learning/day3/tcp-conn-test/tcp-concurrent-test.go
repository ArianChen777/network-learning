package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// TCP多客户端并发连接测试器（10个客户端）
func main() {
	serverAddr := "localhost:8089"
	numClients := 10       // 并发客户端数量
	messagesPerClient := 3 // 每个客户端发送的消息数

	fmt.Printf("启动TCP并发测试: %d个客户端，每个发送%d条消息\n", numClients, messagesPerClient)
	fmt.Printf("目标服务器: %s\n", serverAddr)
	fmt.Println("==================================================")

	var wg sync.WaitGroup
	startTime := time.Now()

	// 启动多个并发客户端
	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()
			testTCPClient(clientID, serverAddr, messagesPerClient)
		}(i)

		// 稍微错开连接时间，观察连接建立过程
		time.Sleep(100 * time.Millisecond)
	}

	// 等待所有客户端完成
	wg.Wait()

	duration := time.Since(startTime)
	fmt.Println("==================================================")
	fmt.Printf("所有客户端测试完成，总耗时: %v\n", duration)
	fmt.Println("现在可以用以下命令查看网络状态:")
	fmt.Println("  netstat -an | grep 8089")
	fmt.Println("  ss -tuln | grep 8089")
	fmt.Println("  lsof -i :8089")
}

func testTCPClient(clientID int, serverAddr string, messageCount int) {
	// 连接到服务器
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Printf("[客户端%d] 连接失败: %v\n", clientID, err)
		return
	}
	defer conn.Close()

	fmt.Printf("[客户端%d] 已连接到服务器 %s\n", clientID, conn.RemoteAddr())

	// 发送多条测试消息
	for i := 0; i < messageCount; i++ {
		message := fmt.Sprintf("Client%d-Message%d", clientID, i+1)

		// 发送消息
		_, err := fmt.Fprintf(conn, "%s\n", message)
		if err != nil {
			fmt.Printf("[客户端%d] 发送消息失败: %v\n", clientID, err)
			return
		}

		// 读取服务器回复
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("[客户端%d] 读取回复失败: %v\n", clientID, err)
			return
		}

		response := string(buffer[:n])
		fmt.Printf("[客户端%d] 收到回复: %s", clientID, response)

		// 在消息之间稍作延迟，模拟真实场景
		time.Sleep(200 * time.Millisecond)
	}

	// 发送quit消息并优雅关闭连接
	fmt.Fprintf(conn, "quit\n")

	// 读取quit的回复
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err == nil {
		response := string(buffer[:n])
		fmt.Printf("[客户端%d] 收到quit回复: %s", clientID, response)
	}

	// 半关闭连接，触发四次挥手
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		tcpConn.CloseWrite()
		fmt.Printf("[客户端%d] 已发送FIN，等待服务器关闭连接\n", clientID)

		// 等待一段时间让四次挥手完成
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("[客户端%d] 连接已关闭\n", clientID)
}
