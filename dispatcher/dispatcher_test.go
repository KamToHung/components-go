package dispatcher

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestConfig struct {
	name string
}

type TestMessage struct {
	id   int
	data string
}

func (t *TestMessage) Route() (open bool, key string) {
	if t.id > 10 {
		return true, "{\"id\":123,\"name\": \"Ailey\"}"
	}
	return false, "{\"id\":123,\"name\": \"Terry\"}"
}

type TestRunner struct {
}

func (t TestRunner) Start(ctx context.Context, config interface{}, ch chan<- Message) {
	testConfig, ok := config.(*TestConfig)
	if !ok {
		panic("invalid config")
	}
	for i := 0; i < 100; i++ {
		ok = true
		select {
		case <-ctx.Done():
			ok = false
		default:
		}
		if !ok {
			break
		}
		ch <- &TestMessage{id: i, data: testConfig.name}
	}
}

func (t TestRunner) Consume(message Message) {
	msg, ok := message.(*TestMessage)
	if !ok {
		panic("invalid message")
	}
	fmt.Printf("message id: %d msg: %s", msg.id, msg.data)
}

func TestDispatcher_GetConsumeMessageCount(t *testing.T) {
	type fields struct {
		producerConfig ProducerConfig
		consumerConfig ConsumerConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				producerConfig: tt.fields.producerConfig,
				consumerConfig: tt.fields.consumerConfig,
			}
			assert.Equalf(t, tt.want, d.GetConsumeMessageCount(), "GetConsumeMessageCount()")
		})
	}
}

func TestDispatcher_GetSendMessageCount(t *testing.T) {
	type fields struct {
		producerConfig ProducerConfig
		consumerConfig ConsumerConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				producerConfig: tt.fields.producerConfig,
				consumerConfig: tt.fields.consumerConfig,
			}
			assert.Equalf(t, tt.want, d.GetSendMessageCount(), "GetSendMessageCount()")
		})
	}
}

func TestDispatcher_Start(t *testing.T) {
	type fields struct {
		producerConfig ProducerConfig
		consumerConfig ConsumerConfig
	}
	type args struct {
		ctx     context.Context
		configs []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				producerConfig: tt.fields.producerConfig,
				consumerConfig: tt.fields.consumerConfig,
			}
			d.Start(tt.args.ctx, tt.args.configs...)
		})
	}
}

func TestNew(t *testing.T) {
	// TODO
}

func Test_calculateIndex(t *testing.T) {
	type args struct {
		key  string
		size int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, calculateIndex(tt.args.key, tt.args.size), "calculateIndex(%v, %v)", tt.args.key, tt.args.size)
		})
	}
}

func Test_consumerProcess(t *testing.T) {
	type args struct {
		consumerConfig *ConsumerConfig
		channels       []chan Message
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			consumerProcess(tt.args.consumerConfig, tt.args.channels)
		})
	}
}

func Test_producerProcess(t *testing.T) {
	type args struct {
		producerConfig   *ProducerConfig
		ctx              context.Context
		consumerChannels []chan Message
		configs          []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			producerProcess(tt.args.producerConfig, tt.args.ctx, tt.args.consumerChannels, tt.args.configs...)
		})
	}
}
