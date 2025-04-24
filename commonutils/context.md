## context

在异步场景中用于实现并发协调以及对goroutine的生命周期控制，并且有一定的数据存储能力
简化对于处理单个请求的多个goroutine之间与请求域的数据、取消信号、截止时间等操作

对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。它们之间的函数调用链必须传递上下文，或者可以使用WithCancel、WithDeadline、WithTimeout或WithValue创建的派生上下文。当一个上下文被取消时，它派生的所有上下文也被取消。

### context.Context
是一个接口，定义了四个需要实现的方法
- Deadline方法返回当前Context被取消的时间————ddl
- Done方法返回一个channel， 这个channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done会返回同一个channel
- Err方法会返回当前Context结束的原因，只会在Done返回的channel被关闭时才会返回非空的值
  - 如果当前Context被取消就会返回Canceled错误
  - 如果当前Context超时就会返回DeadlineExceeded错误
- Value方法会从Context中返回键对应的值，对于同一个context来说，多次调用Value并传入相同的Key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据

### Background, TODO
这两个函数分别返回一个实现了Context接口的background和todo，代码中最开始都是以这两个内置的上下文对象作为最顶层的partent context，衍生出更多上下文对象

Background: 主要用于main函数、初始化以及测试代码中作为根Context
TODO: 如果不知道具体的使用场景，可以使用这个

上面两个方法返回的background和todo(emptyCtx类型)不可取消，没有设置截止时间，没有携带任何value的Context


### With函数

#### WithCancel
```golang
func WithCancel(parent Context) {ctx Context, cancel CancelFunc}
```
返回带有新Done通道的父节点的副本。当调用返回的cancel函数或者当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论什么情况。

取消此上下文将释放与其关联的资源

代码应该在此上下文中运行的操作完成后立即调用cancel

#### WithDeadline
```golang
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
```
返回父上下文的副本，并将deadline调整为不迟于d(第二个实参)。如果父上下文的deadline已经早于d，则WithDeadline(parent, d)在语义上等同于父上下文。

截止日过期时，当调用返回的cancel函数时，或者将父上下文通道关闭时，返回上下文的Done通道将被关闭

取消上下文会释放与其关联的资源

#### WithTimeout
```golang
func WithTimeout(parent Context, time.Duration) (Context, CancelFunc)
```
WithTimeout返回WithDeadline(parent, time.Now().Add(timeout))

通常用于数据库或者网络连接的超时控制

#### WithValue
```golang
func WithValue(parent Context, key, val any) Context
```
返回父节点的副本，其中与key关联的值为value

仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数

所提供的key必须是可比较的，不能是string类型或者任何其他内置类型，以避免使用上下文在包之间发生冲突。WithValue应该为key定义自己的类型，避免在分配给interface{}时进行分配，上下文键通常具有具体类型struct{}。或者导出的上下文关键变量的静态类型应该是指针或接口


### 注意事项
> 推荐以参数的方式显示传递Context
> 以Context作为参数的函数方法，应该把Context作为第一个参数
> 给一个函数方法传递Context的时候，不要传递nil, 如果不知道传递什么，就用context.TODO()
> Context的Value相关方法应该传递请求域的必要数据，不应该用于传递可选参数
> Context是线程安全的，可以放心的在多个goroutine中传递