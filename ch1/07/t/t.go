package t

import "time"

// T 这是一测试.
// 看的到效果吗
// this
/*
`f
fkr
fkd
hello world
 */
func T(i int,c chan<-string,ch chan <-struct{}){

}
type myGenerator struct {
	caller      int64           // 调用器。
	timeoutNs   time.Duration        // 处理超时时间，单位：纳秒。
	lps         uint32               // 每秒载荷量。
	durationNs  time.Duration        // 负载持续时间，单位：纳秒。
	concurrency uint32               // 并发量。
	tickets     int64         // Goroutine票池。
	stopSign    chan byte            // 停止信号的传递通道。
	cancelSign  byte                 // 取消发送后续结果的信号。
	endSign     chan uint64          // 完结信号的传递通道，同时被用于传递调用执行计数。
	callCount   uint64               // 调用执行计数。
	status      int64         // 状态。
	resultCh    chan *int64  // 调用结果通道。
}
