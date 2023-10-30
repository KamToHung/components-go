package dispatcher

type Dispatcher struct {
	producerConfig ProducerConfig
	consumerConfig ConsumerConfig
}

type ProducerConfig struct {
	concurrency  int      // 并发数
	bufferSize   int      // 缓冲区大小
	producer     Producer // 生产者
	messageCount uint64   // 消息数量统计
}

type ConsumerConfig struct {
	concurrency  int      // 并发数
	bufferSize   int      // 缓冲区大小
	consumer     Consumer // 生产者
	messageCount uint64   // 消息数量统计
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
