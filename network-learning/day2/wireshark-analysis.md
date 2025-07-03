# Wireshark抓包分析：TCP vs UDP

## 抓包前的准备工作

### 1. 启动Wireshark
```bash
# 如果已安装Wireshark，可以直接启动
wireshark
```

### 2. 启动服务器程序
```bash
# 终端1：启动TCP服务器
cd network-learning/day1/server
go run tcp-server.go

# 终端2：启动UDP服务器
cd network-learning/day2
go run udp-server.go
```

### 3. 配置Wireshark过滤器
- TCP分析：`tcp.port == 8089`
- UDP分析：`udp.port == 8090`
- 同时观察：`tcp.port == 8089 or udp.port == 8090`

## TCP协议抓包分析

### 三次握手过程
1. **SYN**: 客户端发送SYN包，序列号为随机值
2. **SYN-ACK**: 服务器回复SYN-ACK包，确认序列号+1
3. **ACK**: 客户端发送ACK包，确认连接建立

### 数据传输过程
- 每个数据包都有序列号和确认号
- 服务器会确认收到的每个数据包
- 数据包大小可变，受MSS限制

### 四次挥手过程
1. **FIN**: 客户端发送FIN包，请求关闭连接
2. **ACK**: 服务器确认FIN包
3. **FIN**: 服务器发送FIN包，准备关闭连接
4. **ACK**: 客户端确认FIN包，连接完全关闭

## UDP协议抓包分析

### 数据传输特点
- **无连接**: 直接发送数据，无需建立连接
- **无状态**: 每个数据包独立处理
- **无确认**: 发送后不等待确认

### 数据包结构
- UDP头部只有8字节（源端口、目标端口、长度、校验和）
- 相比TCP头部（20字节），开销更小

## 关键差异对比

| 特性 | TCP | UDP |
|------|-----|-----|
| 连接建立 | 需要三次握手 | 无需连接建立 |
| 数据传输 | 可靠，有序 | 不可靠，无序 |
| 头部大小 | 20字节 | 8字节 |
| 连接关闭 | 四次挥手 | 直接停止发送 |
| 适用场景 | 文件传输、Web | 实时游戏、视频流 |

## 实际测试步骤

### 测试TCP连接
```bash
# 终端3：启动TCP客户端
cd network-learning/day1/client
go run tcp-client.go

# 发送消息：hello
# 发送消息：world
# 发送消息：quit
```

### 测试UDP连接
```bash
# 终端4：启动UDP客户端
cd network-learning/day2
go run udp-client.go

# 发送消息：hello
# 发送消息：world
# 发送消息：quit
```

## 观察要点

### TCP抓包重点
1. 观察SYN、SYN-ACK、ACK的序列号变化
2. 数据包的PSH标志位
3. FIN、ACK的四次挥手过程
4. TIME_WAIT状态的持续时间

### UDP抓包重点
1. 数据包的简单结构
2. 无连接建立和关闭过程
3. 每个数据包的独立性
4. 比TCP更少的网络开销

## 学习收获

通过抓包分析，你将深入理解：
- TCP的可靠性机制如何实现
- UDP的高效性体现在哪里
- 两种协议的适用场景
- 网络协议的底层工作原理