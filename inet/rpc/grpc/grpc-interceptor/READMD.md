## grpc中间件


- bin 二进制文件
- client grpc客户端
- server grpc服务端
- service 服务实现


MakeFile 记录protoc生成go代码的模板



### 中间件
#### 一元中间件(只添加一个中间件)
grpc.UnaryInterceptor(func)

#### 链式调用(多个一元中间件)
grpc.ChainUnaryInterceptor(func1, func2, func3...)
根据传入顺序，依次进行处理

#### 官方中间件
go-grpc-middleware提供了很多现成的中间件可以使用