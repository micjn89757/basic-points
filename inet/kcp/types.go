package kcp


//==================================
// KCP BASIC
//==================================

const (
	IKCP_RTO_NDL = 30		// 


	// 报文类型
	IKCP_CMD_PUSH int8 = 81
	IKCP_CMD_ACK  int8 = 82
	IKCP_CMD_WASK int8 = 83
	IKCP_CMD_WINS int8 = 84
)


type KcpPacketHeader struct {
	conv 	uint32			// 连接标识
	cmd  	uint8    		// 报文类型
	frg  	uint8			// 数据分段标识
	wnd  	uint16			// 接收端通告窗口大小

	ts		uint32			// 时间戳
	sn 		uint32 			// 报文序列号

	una 	uint32 			// 接收滑动窗口起始序号标识，所有小于sn的报文都已有序到达
	len 	uint32			// 数据部分长度
}



