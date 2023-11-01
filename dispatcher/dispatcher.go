package dispatcher

import (
	"context"
	"sync/atomic"
)

type Dispatcher struct {
	producerConfig ProducerConfig
	consumerConfig ConsumerConfig
}

type ProducerConfig struct {
	concurrency  int            // 并发数
	bufferSize   int            // 缓冲区大小
	producer     ProducerOption // 生产者
	messageCount uint64         // 消息数量统计
}

type ConsumerConfig struct {
	concurrency  int            // 并发数
	bufferSize   int            // 缓冲区大小
	consumer     ConsumerOption // 生产者
	messageCount uint64         // 消息数量统计
}

// DefaultDispatcher 调度器默认配置
var defaultDispatcher = Dispatcher{
	producerConfig: ProducerConfig{
		concurrency: 1,
		bufferSize:  1,
	},
	consumerConfig: ConsumerConfig{
		concurrency: 1,
		bufferSize:  1,
	},
}

// New 创建调度器
// @param opts 配置选项
// @return *Dispatcher 调度器
func New(opts ...Option) *Dispatcher {
	d := defaultDispatcher
	for _, opt := range opts {
		opt(&d)
	}
	return &d
}

// GetSendMessageCount 获取发送消息数量
// @receiver d 调度器
// @return uint64 消息数量
func (d *Dispatcher) GetSendMessageCount() uint64 {
	return atomic.LoadUint64(&d.producerConfig.messageCount)
}

// GetConsumeMessageCount 获取消费消息数量
// @receiver d 调度器
// @return uint64 消息数量
func (d *Dispatcher) GetConsumeMessageCount() uint64 {
	return atomic.LoadUint64(&d.consumerConfig.messageCount)
}

// Start  启动调度器
// @receiver d 调度器
// @param ctx 上下文
// @param configs 配置信息
func (d *Dispatcher) Start(ctx context.Context, configs ...interface{}) {
	// TODO
}
