/*
	使用net/rpc并使用tcp协议实现rpc
*/

package main

// 参数
type Args struct {
	X, Y int
}


// Service
type Service struct{}


// Add 为Service类型增加一个可导出的Add方法
func (s *Service) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}


