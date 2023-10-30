package dispatcher

// Option 配置选项
type Option func(d *Dispatcher)

// ProducerOption 生产者
type ProducerOption func() Producer

// ConsumerOption 消费者
type ConsumerOption func() Consumer

// OptProducerConfig 设置生产者配置
// @param option 生产者配置
// @return Option 配置选项
func OptProducerConfig(option ProducerOption) Option {
	return func(d *Dispatcher) {
		d.producerConfig.producer = option
	}
}

// OptProducerConcurrency 设置生产者并发数
// @param c 并发数
// @return Option 配置选项
func OptProducerConcurrency(c int) Option {
	if c <= 0 {
		panic("concurrency value is not valid")
	}
	return func(d *Dispatcher) {
		d.producerConfig.concurrency = c
	}
}

// OptProducerBufferSize 设置生产者缓冲区大小
// @param s 缓冲区大小
// @return Option 配置选项
func OptProducerBufferSize(s int) Option {
	if s <= 0 {
		panic("buffer size value is not valid")
	}
	return func(d *Dispatcher) {
		d.producerConfig.bufferSize = s
	}
}

// OptConsumerConfig 设置消费者配置
// @param option 消费者配置
// @return Option 配置选项
func OptConsumerConfig(option ConsumerOption) Option {
	return func(d *Dispatcher) {
		d.consumerConfig.consumer = option
	}
}

// OptConsumerConcurrency 设置消费者并发数
// @param c 并发数
// @return Option 配置选项
func OptConsumerConcurrency(c int) Option {
	if c <= 0 {
		panic("concurrency value is not valid")
	}
	return func(d *Dispatcher) {
		d.consumerConfig.concurrency = c
	}
}

// OptConsumerBufferSize 设置消费者缓冲区大小
// @param s 缓冲区大小
// @return Option 配置选项
func OptConsumerBufferSize(s int) Option {
	if s <= 0 {
		panic("buffer size value is not valid")
	}
	return func(d *Dispatcher) {
		d.consumerConfig.bufferSize = s
	}
}
