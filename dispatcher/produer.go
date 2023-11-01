package dispatcher

import "context"

type Producer interface {
	// Start
	// 启动producer
	// @param ctx 上下文
	// @param config 配置信息
	// @param ch chan
	Start(ctx context.Context, config interface{}, ch chan<- Message)
}
