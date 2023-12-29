package dispatcher

import (
	"context"
	"hash/fnv"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
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
	// channels
	consumerChannels := make([]chan Message, d.consumerConfig.bufferSize)
	pWaitGroup := sync.WaitGroup{}
	cWaitGroup := sync.WaitGroup{}
	consumerProcess(&cWaitGroup, &d.consumerConfig, consumerChannels)
	producerProcess(&cWaitGroup, &d.producerConfig, consumerChannels, ctx, configs)
	pWaitGroup.Wait()
	closeConsumerChannel(consumerChannels)
	cWaitGroup.Wait()
}

// consumerProcess 消费者处理
// @param consumerConfig 消费者配置
func consumerProcess(waitGroup *sync.WaitGroup, consumerConfig *ConsumerConfig, channels []chan Message) {
	if consumerConfig.consumer == nil {
		panic("consumerConfig is not set")
	}
	// consumerConfig config
	concurrency := consumerConfig.concurrency
	bufferSize := consumerConfig.bufferSize

	//var waitGroup sync.WaitGroup
	waitGroup.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		// create
		ch := make(chan Message, bufferSize)
		channels = append(channels, ch)
		// consumer
		consumer := consumerConfig.consumer()
		go func() {
			defer waitGroup.Done()
			for message := range ch {
				consumer.Consume(message)
				// count message
				atomic.AddUint64(&consumerConfig.messageCount, 1)
			}
		}()
	}
	//waitGroup.Wait()
}

// producerProcess  生产者处理
// @param producerConfig 生产者配置
// @param ctx 上下文
// @param consumerChannels 消费者通道
// @param configs 配置信息
func producerProcess(waitGroup *sync.WaitGroup, producerConfig *ProducerConfig, consumerChannels []chan Message, ctx context.Context, configs []interface{}) {
	if producerConfig.producer == nil {
		panic("producer is not set")
	}
	// producer config
	concurrency := producerConfig.concurrency
	bufferSize := producerConfig.bufferSize

	// producer channels
	configChannels := make(chan interface{}, concurrency)
	go func() {
		for _, config := range configs {
			configChannels <- config
		}
		close(configChannels)
	}()

	//var waitGroup sync.WaitGroup
	waitGroup.Add(concurrency)
	rand.Seed(time.Now().Unix())
	for i := 0; i < concurrency; i++ {
		go func() {
			defer waitGroup.Done()
			for channel := range configChannels {
				ch := make(chan Message, bufferSize)
				go func() {
					producerConfig.producer().Start(ctx, channel, ch)
					close(ch)
				}()

				for message := range ch {
					open, key := message.Route()
					index := 0
					if open {
						// 根据key发送到同一个consumer
						index = calculateIndex(key, concurrency)
					} else {
						// 随机发送到一个consumer
						index = rand.Intn(concurrency)
					}
					// producer count
					atomic.AddUint64(&producerConfig.messageCount, 1)
					// send message
					consumerChannels[index] <- message
				}
			}
		}()
	}
	//waitGroup.Wait()
}

func closeConsumerChannel(consumerChannels []chan Message) {
	// 关闭所有通道
	for idx := range consumerChannels {
		close(consumerChannels[idx])
	}
}

// calculateIndex 计算索引
// @param key 键值
// @param size 大小
// @return int 索引
func calculateIndex(key string, size int) int {
	hash := fnv.New32a()
	_, err := hash.Write([]byte(key))
	if err != nil {
		panic(err)
	}
	hashValue := hash.Sum32()
	return int(hashValue) % size
}
