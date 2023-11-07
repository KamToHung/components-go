package dispatcher

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
