package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// UDP并发测试器（15个客户端，含成功率统计）
func main() {
	serverAddr := "localhost:8090"
	numClients := 15       // UDP可以支持更多并发
	messagesPerClient := 5 // 每个客户端发送的消息数

	fmt.Printf("启动UDP并发测试: %d个客户端，每个发送%d条消息\n", numClients, messagesPerClient)
	fmt.Printf("目标服务器: %s\n", serverAddr)
	fmt.Println("==================================================")

	var wg sync.WaitGroup
	startTime := time.Now()

	// 用于统计成功和失败的消息
	var successCount, failCount int64
	var mu sync.Mutex

	// 启动多个并发客户端
	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()
			success, fail := testUDPClient(clientID, serverAddr, messagesPerClient)

			mu.Lock()
			successCount += success
			failCount += fail
			mu.Unlock()
		}(i)

		// UDP无需等待连接建立，可以更快启动
		time.Sleep(50 * time.Millisecond)
	}

	// 等待所有客户端完成
	wg.Wait()

	duration := time.Since(startTime)
	fmt.Println("==================================================")
	fmt.Printf("所有客户端测试完成，总耗时: %v\n", duration)
	fmt.Printf("成功消息: %d, 失败消息: %d\n", successCount, failCount)
	fmt.Printf("总消息数: %d, 成功率: %.2f%%\n",
		successCount+failCount,
		float64(successCount)/float64(successCount+failCount)*100)

	fmt.Println("现在可以用以下命令查看网络状态:")
	fmt.Println("  netstat -an | grep 8090")
	fmt.Println("  ss -u | grep 8090")
	fmt.Println("  lsof -i :8090")
}

func testUDPClient(clientID int, serverAddr string, messageCount int) (success, fail int64) {
	// 解析服务器地址
	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Printf("[UDP客户端%d] 地址解析失败: %v\n", clientID, err)
		return 0, int64(messageCount)
	}

	// 连接到UDP服务器
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Printf("[UDP客户端%d] 连接失败: %v\n", clientID, err)
		return 0, int64(messageCount)
	}
	defer conn.Close()

	fmt.Printf("[UDP客户端%d] 已连接到服务器 %s\n", clientID, conn.RemoteAddr())

	// 发送多条测试消息
	for i := 0; i < messageCount; i++ {
		message := fmt.Sprintf("UDPClient%d-Message%d", clientID, i+1)

		// 发送消息
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("[UDP客户端%d] 发送消息失败: %v\n", clientID, err)
			fail++
			continue
		}

		// 设置读取超时
		conn.SetReadDeadline(time.Now().Add(2 * time.Second))

		// 读取服务器回复
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("[UDP客户端%d] 读取回复超时或失败: %v\n", clientID, err)
			fail++
			continue
		}

		response := string(buffer[:n])
		fmt.Printf("[UDP客户端%d] 收到回复: %s\n", clientID, response)
		success++

		// UDP消息之间的间隔可以更短
		time.Sleep(100 * time.Millisecond)
	}

	// UDP无需优雅关闭，直接退出
	fmt.Printf("[UDP客户端%d] 测试完成，成功: %d, 失败: %d\n", clientID, success, fail)
	return success, fail
}
