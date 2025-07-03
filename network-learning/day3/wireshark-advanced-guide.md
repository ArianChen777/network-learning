# Wireshark高级分析指南：并发连接与协议深度解析

## 高级过滤器语法

### 1. 基础过滤器
```wireshark
# 单个端口
tcp.port == 8089

# 多个端口
tcp.port == 8089 or udp.port == 8090

# IP地址过滤
ip.addr == 127.0.0.1

# 组合条件
tcp.port == 8089 and ip.addr == 127.0.0.1
```

### 2. TCP状态过滤
```wireshark
# 三次握手包
tcp.flags.syn == 1

# 四次挥手包
tcp.flags.fin == 1

# 数据包（有负载）
tcp.len > 0

# ACK包
tcp.flags.ack == 1
```

### 3. 高级TCP分析
```wireshark
# 重传包
tcp.analysis.retransmission

# 乱序包
tcp.analysis.out_of_order

# 零窗口
tcp.window_size == 0

# 重复ACK
tcp.analysis.duplicate_ack
```

## 并发连接分析步骤

### 步骤1：启动抓包环境
```bash
# 终端1：启动Wireshark并设置过滤器
wireshark &
# 过滤器设置为: tcp.port == 8089 or udp.port == 8090

# 终端2：启动TCP服务器
cd day1/server
go run tcp-server.go

# 终端3：启动UDP服务器  
cd day2/server
go run udp-server.go
```

### 步骤2：执行并发测试
```bash
# 终端4：执行TCP并发测试
cd day2
go run tcp-concurrent-test.go

# 稍等片刻，再执行UDP测试
go run udp-concurrent-test.go
```

### 步骤3：分析捕获的数据包

## TCP并发连接分析要点

### 1. 连接建立模式
- **观察重点**: 多个客户端的SYN包时序
- **分析方法**: 
  ```wireshark
  tcp.flags.syn == 1 and tcp.flags.ack == 0
  ```
- **关注指标**:
  - 源端口分配规律
  - 序列号初始值
  - 连接建立时间间隔

### 2. 数据传输分析
- **过滤器**: `tcp.port == 8089 and tcp.len > 0`
- **观察重点**:
  - 不同连接的序列号管理
  - PSH标志的使用
  - 窗口大小变化

### 3. 连接关闭分析
- **过滤器**: `tcp.flags.fin == 1`
- **观察重点**:
  - 四次挥手的完整过程
  - TIME_WAIT状态持续时间
  - 客户端主动关闭 vs 服务器主动关闭

## UDP并发通信分析

### 1. 无连接特性观察
- **过滤器**: `udp.port == 8090`
- **对比要点**:
  - 无握手过程，直接数据传输
  - 每个数据包的独立性
  - 源端口的变化模式

### 2. 数据包结构对比
```wireshark
# UDP头部分析
udp

# TCP头部分析  
tcp
```

## Follow TCP Stream功能

### 使用方法
1. 右键点击任意TCP数据包
2. 选择"Follow" → "TCP Stream"
3. 查看完整的对话流程

### 分析价值
- 查看完整的应用层数据交换
- 理解请求-响应模式
- 发现协议解析问题

## 统计分析功能

### 1. 连接统计
- **菜单**: Statistics → Conversations
- **TCP标签页**: 查看每个连接的统计信息
- **分析指标**:
  - 数据传输量
  - 连接持续时间
  - 数据包数量

### 2. 协议层次统计
- **菜单**: Statistics → Protocol Hierarchy
- **分析用途**:
  - TCP vs UDP流量占比
  - 应用层协议分布
  - 数据传输效率

### 3. I/O图形分析
- **菜单**: Statistics → I/O Graphs
- **设置过滤器**: 
  - Filter 1: `tcp.port == 8089`
  - Filter 2: `udp.port == 8090`
- **观察指标**: 并发连接的时间分布

## 高级分析技巧

### 1. 序列号分析
```wireshark
# 显示相对序列号
tcp.seq

# 显示下一个期望序列号
tcp.nxtseq

# 分析序列号跳跃
tcp.analysis.lost_segment
```

### 2. 时间分析
```wireshark
# RTT分析
tcp.analysis.ack_rtt

# 连接建立时间
tcp.time_relative

# 数据包间隔时间
frame.time_delta
```

### 3. 窗口分析
```wireshark
# 接收窗口大小
tcp.window_size

# 窗口缩放因子
tcp.options.wscale

# 零窗口探测
tcp.analysis.zero_window_probe
```

## 实际抓包场景

### 场景1：正常并发连接
```bash
# 预期观察结果
1. 10个TCP连接的三次握手
2. 有序的数据传输
3. 优雅的四次挥手
4. TIME_WAIT状态的持续
```

### 场景2：UDP高并发测试
```bash
# 预期观察结果
1. 无连接建立过程
2. 大量并发数据包
3. 可能的丢包现象
4. 无连接关闭过程
```

### 场景3：资源耗尽场景
```bash
# 增加并发数测试
go run tcp-concurrent-test.go  # 修改为100个客户端

# 观察异常情况
1. 连接被拒绝
2. 超时重传
3. RST包的出现
4. 端口耗尽现象
```

## 导出和保存分析结果

### 1. 保存数据包
- **格式**: .pcap文件
- **用途**: 后续详细分析
- **操作**: File → Save As

### 2. 导出统计数据
- **连接统计**: 复制到Excel分析
- **图形数据**: 导出为CSV
- **流数据**: 保存为文本文件

### 3. 截图记录
- **关键时刻**: 三次握手、四次挥手
- **异常情况**: 重传、RST、超时
- **统计图表**: I/O图形、协议分布

## 学习检查点

### 基础掌握
- [ ] 能设置复杂的过滤器表达式
- [ ] 理解TCP连接的完整生命周期
- [ ] 掌握Follow TCP Stream功能
- [ ] 会使用统计分析功能

### 进阶技能
- [ ] 能分析并发连接的性能特征
- [ ] 识别网络异常和性能问题
- [ ] 对比不同协议的网络行为
- [ ] 理解系统资源限制的影响

通过这些高级分析技巧，你将能够深入理解网络协议的工作机制和性能特征。