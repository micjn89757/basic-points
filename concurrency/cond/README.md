## sync.Cond

Cond实现了一个广播和通知的功能，主要还是用于有多个goroutine的情况下，通过一个变量来对两种不同类型的时间进行控制，当然尽量使用channel来进行协程通信


## 方法
### func NewCond(L Locker) *Cound 

传入一把锁，返回一个*Cond对象

锁对象放在其L属性中
```go 
func main() {
    m := &sync.Mutex{}
    c := sync.NewCond(m)
    c.L.Lock()
    defer c.L.Lock()
}
```

### func (*Cond) Broadcast()
通过广播的方式唤醒所有goroutine


### func (*Cond) Signal()
遍历阻塞的goroutine队列，唤醒第一个goroutine

### func (*Cond) Wait()
首先必须获得Cond中锁的持有权，然后自动释放c.L，并将当前goroutine陷入阻塞，并添加到阻塞队列中，所以goroutine会阻塞在Wait方法调用的地方，如果其他goroutine调用了Signal或Broadcast唤醒了该协程，那么Wait方法在结束阻塞时，会重新给c.L加锁(重新持有锁)，并且继续执行Wait后面的代码


