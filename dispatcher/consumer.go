package dispatcher

type Consumer interface {
	Consume(message Message)
}
