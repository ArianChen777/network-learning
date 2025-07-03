# 第1周详细学习计划：网络基础实战 + 抓包技能

## 学习目标
- 从实践中深入理解TCP/UDP协议
- 掌握Wireshark和tcpdump抓包分析技能
- 学会Go语言网络编程基础
- 具备基础的网络问题排查能力

## 每日详细计划

### Day 1 (Monday): 环境搭建 + TCP基础实战

#### 上午 (2-3小时)
**任务1: 环境准备**
- 安装Wireshark
- 安装Go开发环境 (如果还没有)
- 准备Linux虚拟机或WSL环境

**任务2: TCP客户端-服务器实现**
```bash
# 创建项目目录
mkdir -p ~/network-learning/day1
cd ~/network-learning/day1
```

实现简单的TCP echo服务器和客户端：
- `server.go`: TCP服务器，监听8080端口，回显收到的消息
- `client.go`: TCP客户端，连接服务器并发送消息

#### 下午 (2-3小时)
**任务3: Wireshark抓包分析**
- 启动Wireshark，选择网络接口
- 运行TCP服务器和客户端
- 抓取数据包，观察TCP三次握手过程
- 分析数据包结构：IP头、TCP头、数据部分

**学习重点:**
- TCP头部字段含义 (seq, ack, flags等)
- 三次握手的具体过程
- Wireshark基本过滤器使用

#### 晚上 (1小时)
**复习总结:**
- 记录今天遇到的问题和解决方案
- 整理TCP协议关键知识点
- 准备明天的学习材料

---

### Day 2 (Tuesday): UDP编程 + 四次挥手分析

#### 上午 (2-3小时)
**任务1: UDP实现**
- 实现UDP echo服务器和客户端
- 对比TCP和UDP的编程差异
- 观察UDP数据包结构

**任务2: TCP连接关闭分析**
- 修改TCP客户端，观察四次挥手过程
- 分析FIN、ACK包的交换
- 理解TIME_WAIT状态

#### 下午 (2-3小时)  
**任务3: 网络状态查看**
- 学习netstat/ss命令
- 查看LISTEN、ESTABLISHED、TIME_WAIT等状态
- 使用lsof查看进程占用的端口

**任务4: 抓包进阶**
- 学习Wireshark高级过滤器
- 使用"tcp.port == 8080"等过滤条件
- 学会跟踪TCP流 (Follow TCP Stream)

#### 晚上 (1小时)
**实践练习:**
- 同时运行多个客户端连接服务器
- 观察并发连接的数据包交互
- 分析不同连接的序列号

---

### Day 3 (Wednesday): tcpdump + 网络异常模拟

#### 上午 (2-3小时)
**任务1: tcpdump命令行抓包**
- 学习tcpdump基本语法和参数
- 抓取特定端口、IP的数据包
- 将抓包结果保存为pcap文件，在Wireshark中分析

常用命令练习：
```bash
# 抓取8080端口的包
sudo tcpdump -i any port 8080 -w capture.pcap

# 实时查看HTTP流量
sudo tcpdump -i any port 80 -A

# 抓取特定主机的包
sudo tcpdump -i any host 192.168.1.100
```

#### 下午 (2-3小时)
**任务2: 网络异常模拟**
- 使用tc命令模拟网络延迟
- 模拟丢包情况
- 观察TCP如何处理这些异常

```bash
# 添加100ms延迟
sudo tc qdisc add dev eth0 root handle 1: prio
sudo tc qdisc add dev eth0 parent 1:3 handle 30: netem delay 100ms

# 模拟1%丢包率
sudo tc qdisc add dev eth0 root netem loss 1%
```

**任务3: 异常情况分析**
- 抓包分析重传机制
- 观察超时和窗口调整
- 理解TCP的可靠性保证机制

#### 晚上 (1小时)
**深度学习:**
- 阅读Go net包源码中的关键部分
- 理解Go是如何封装系统调用的

---

### Day 4 (Thursday): Go网络编程进阶

#### 上午 (2-3小时)
**任务1: 并发服务器实现**
- 实现能处理多客户端连接的TCP服务器
- 使用goroutine处理每个连接
- 学习连接管理和资源清理

**任务2: 超时处理**
- 实现连接超时控制
- 使用context包管理超时
- 学习SetDeadline的使用

#### 下午 (2-3小时)
**任务3: 高级网络编程**
- 实现带缓冲的数据处理
- 学习bufio包的使用
- 实现简单的协议解析

**任务4: 性能测试**
- 编写简单的性能测试客户端
- 测试服务器的并发处理能力
- 使用go tool pprof分析性能

#### 晚上 (1小时)
**项目整合:**
- 整合前几天的代码
- 创建一个小型的网络测试工具
- 添加日志和错误处理

---

### Day 5 (Friday): 综合实战项目

#### 上午 (2-3小时)
**任务1: 网络调试工具开发**
创建一个综合的网络调试工具，包含：
- TCP/UDP连接测试功能
- 端口扫描功能
- 延迟测试功能
- 简单的数据包分析功能

#### 下午 (2-3小时)
**任务2: 真实场景问题排查**
- 模拟常见网络问题（端口不通、服务不响应等）
- 使用本周学到的工具进行排查
- 形成问题排查的标准流程

**任务3: 学习成果验证**
- 完成网络基础知识自测
- 实际操作技能验证
- 准备下周学习内容

#### 晚上 (1-2小时)
**周总结:**
- 整理本周学习笔记
- 总结掌握的技能和工具
- 规划下周学习重点

---

## 学习资源推荐

### 书籍资料
1. **《TCP/IP详解 卷1》** - 经典网络协议书籍，重点看前几章
2. **《Go语言圣经》** - 第8章：Goroutines和Channels
3. **在线资源**: 
   - Go官方网络编程教程
   - Wireshark官方用户手册

### 实用工具清单
```bash
# 必备工具安装
sudo apt-get update
sudo apt-get install wireshark tcpdump netstat-nat lsof iperf3

# Go相关工具
go install golang.org/x/tools/cmd/pprof@latest
```

### 实践项目代码结构
```
~/network-learning/
├── day1/
│   ├── tcp-server.go
│   ├── tcp-client.go
│   └── README.md
├── day2/
│   ├── udp-server.go
│   ├── udp-client.go
│   └── analysis.md
├── day3/
│   ├── tcpdump-examples.sh
│   └── network-simulation.sh
├── day4/
│   ├── concurrent-server.go
│   └── performance-test.go
└── day5/
    ├── network-debugger/
    └── week-summary.md
```

## 评估标准

### 技能掌握检查表
**基础技能 (必须掌握):**
- [ ] 能独立编写TCP/UDP服务器和客户端
- [ ] 熟练使用Wireshark分析数据包
- [ ] 掌握tcpdump基本命令
- [ ] 理解TCP三次握手、四次挥手过程
- [ ] 会使用netstat/ss查看网络状态

**进阶技能 (建议掌握):**
- [ ] 能分析网络性能问题
- [ ] 理解TCP重传和拥塞控制机制
- [ ] 会模拟网络异常进行测试
- [ ] 能编写并发网络程序
- [ ] 掌握Go语言网络编程最佳实践

### 项目成果
1. **完成的代码项目**: 至少5个可运行的网络程序
2. **抓包分析报告**: 包含TCP/UDP协议分析的文档
3. **问题排查案例**: 至少3个网络问题的排查过程记录
4. **网络调试工具**: 一个综合性的命令行网络工具

## 注意事项

1. **时间安排建议**: 每天实际学习时间4-6小时，理论与实践比例3:7
2. **遇到困难时**: 优先查官方文档，其次是Stack Overflow，最后考虑询问同事
3. **代码管理**: 建议使用Git管理所有练习代码，养成良好的版本控制习惯
4. **学习记录**: 每天记录学习日志，包括遇到的问题、解决方案、心得体会

---

**下周预告**: 第2周将深入学习四层协议详细机制，并开始实现真正的负载均衡器，为实际工作做准备。