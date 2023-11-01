package dispatcher

type Message interface {
	// Route
	// 是否需要把指定key的消息路由到同一个消费者
	// 这样的话同一个key的消息就不会并发消费
	Route() (open bool, key string)
}
