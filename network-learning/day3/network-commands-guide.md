# 网络状态查看命令实战指南

## 命令概览

### 1. netstat - 网络连接状态查看
```bash
# 查看所有TCP连接
netstat -an | grep tcp

# 查看特定端口的连接
netstat -an | grep 8089

# 查看连接状态统计
netstat -s

# 查看监听端口
netstat -ln
```

### 2. ss - 现代替代工具（更快）
```bash
# 查看所有TCP连接
ss -tuln

# 查看特定端口
ss -tuln | grep 8089

# 查看连接详细信息
ss -i

# 查看进程信息
ss -p
```

### 3. lsof - 查看进程打开的文件和端口
```bash
# 查看特定端口被哪个进程占用
lsof -i :8089

# 查看某个进程打开的所有网络连接
lsof -p <进程ID>

# 查看所有网络连接
lsof -i
```

## 实战练习

### 练习1：观察TCP服务器状态
```bash
# 终端1：启动TCP服务器
cd day1/server
go run tcp-server.go

# 终端2：查看监听状态
netstat -an | grep 8089
# 应该看到: LISTEN状态

lsof -i :8089
# 应该看到: go进程监听8089端口
```

### 练习2：观察TCP并发连接
```bash
# 终端1：保持TCP服务器运行

# 终端2：启动并发测试
cd day2
go run tcp-concurrent-test.go

# 终端3：实时监控连接状态
watch "netstat -an | grep 8089"
# 观察连接状态变化: LISTEN -> ESTABLISHED -> TIME_WAIT
```

### 练习3：观察UDP连接特点
```bash
# 终端1：启动UDP服务器
cd day2/server
go run udp-server.go

# 终端2：启动UDP并发测试
cd day2
go run udp-concurrent-test.go

# 终端3：查看UDP连接
netstat -an | grep 8090
ss -u | grep 8090
# 注意UDP的连接状态特点
```

## 关键状态解析

### TCP连接状态
- **LISTEN**: 服务器正在监听端口
- **ESTABLISHED**: 连接已建立，正在通信
- **TIME_WAIT**: 连接关闭后的等待状态（通常持续2MSL）
- **CLOSE_WAIT**: 等待本地应用关闭连接
- **FIN_WAIT_1/FIN_WAIT_2**: 主动关闭方的等待状态

### UDP特点
- UDP没有连接状态概念
- 只显示监听端口，不显示具体连接
- 无需连接建立和关闭过程

## 高并发场景观察要点

### 1. 连接数限制
```bash
# 查看系统文件描述符限制
ulimit -n

# 查看当前连接数
netstat -an | grep ESTABLISHED | wc -l
```

### 2. TIME_WAIT状态
```bash
# 统计TIME_WAIT连接数
netstat -an | grep TIME_WAIT | wc -l

# 观察TIME_WAIT持续时间
ss -o state time-wait
```

### 3. 端口占用情况
```bash
# 查看端口使用范围
cat /proc/sys/net/ipv4/ip_local_port_range

# 查看可用端口数量
netstat -an | grep tcp | wc -l
```

## 性能分析命令

### 网络统计信息
```bash
# TCP统计
netstat -s | grep -i tcp

# UDP统计  
netstat -s | grep -i udp

# 查看重传、丢包等信息
ss -i
```

### 实时监控
```bash
# 实时显示连接变化
watch "netstat -an | grep 8089 | head -20"

# 实时显示统计信息
watch "netstat -s | grep -E 'segments|packets'"
```

## 故障排查场景

### 1. 端口被占用
```bash
# 查找占用端口的进程
lsof -i :8089
kill -9 <进程ID>
```

### 2. 连接数过多
```bash
# 查看连接数最多的IP
netstat -an | grep ESTABLISHED | awk '{print $5}' | cut -d: -f1 | sort | uniq -c | sort -nr
```

### 3. TIME_WAIT过多
```bash
# 优化TIME_WAIT设置（仅供学习，生产环境需谨慎）
echo 1 > /proc/sys/net/ipv4/tcp_tw_reuse
echo 1 > /proc/sys/net/ipv4/tcp_tw_recycle
```

## 测试场景设计

### 场景1：正常负载测试
- 启动10个TCP客户端
- 每个发送3条消息
- 观察连接建立、数据传输、连接关闭全过程

### 场景2：高并发压力测试
- 启动50-100个并发连接
- 观察系统资源使用情况
- 监控TIME_WAIT状态堆积

### 场景3：UDP vs TCP对比
- 同时测试TCP和UDP
- 对比连接开销和处理能力
- 分析不同协议的适用场景

通过这些命令和场景，你将深入理解网络连接的生命周期和系统资源管理机制。