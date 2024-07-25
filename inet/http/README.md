## 使用go 1.22 构建api服务器


### 主要实现功能
1. 在handers中获取路径参数
2. 中间件
3. 声明http method
4. 子路由


### net/http

#### client
##### Transport
要管理代理、TLS配置、keep-alive等其他设置
并发安全
应该一次建立、尽量重用



#### server

##### 自定义server