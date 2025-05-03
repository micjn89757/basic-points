## grpc中间件


- bin 二进制文件
- client grpc客户端
- server grpc服务端
- service 服务实现


MakeFile 记录protoc生成go代码的模板

grpc中的上下文信息都是使用context承载

### 中间件(拦截器)
#### 一元中间件(只添加一个中间件)
grpc.UnaryInterceptor(func)

#### 链式调用(多个一元中间件)
grpc.ChainUnaryInterceptor(func1, func2, func3...)
根据传入顺序，依次进行处理

#### 官方中间件
go-grpc-middleware提供了很多现成的中间件可以使用



### metadata
> 相当于http报文的header, 存放很多key-value数据，key是string类型，value通常是[]string或二进制数据

#### 接受数据——incoming
从RPC请求的上下文中获取元数据:metadata.FromIncomingContext(ctx)

#### 发送数据——outgoing
使用NewOutgoingContext方法会覆盖原有的metadata
使用AppendOutgoingContext是一种追加metadata模式, 保留上有服务的Metadata,也要比NewOutgoingContext快