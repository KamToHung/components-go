package dispatcher

import (
	"context"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type TestRunner struct {
}

func (t TestRunner) Start(ctx context.Context, config interface{}, ch chan<- Message) {
	//TODO implement me
	panic("implement me")
}

func (t TestRunner) Consume(message Message) {
	//TODO implement me
	panic("implement me")
}

func TestOptConsumerBufferSize(t *testing.T) {
	type args struct {
		s int
	}

	bufferSize := 512
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "consumer buffer size",
			args: args{
				s: bufferSize,
			},
			want: bufferSize,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(OptConsumerBufferSize(tt.args.s)).consumerConfig.bufferSize; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OptConsumerBufferSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptConsumerConcurrency(t *testing.T) {
	type args struct {
		c int
	}

	concurrency := 512
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "consumer concurrency",
			args: args{
				c: concurrency,
			},
			want: concurrency,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(OptConsumerConcurrency(tt.args.c)).consumerConfig.concurrency; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OptConsumerConcurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptConsumerConfig(t *testing.T) {
	runner := &TestRunner{}
	dispatcher := New(OptConsumerConfig(func() Consumer {
		return runner
	}))
	assert.Equal(t, dispatcher.consumerConfig.consumer(), runner)
}

func TestOptProducerBufferSize(t *testing.T) {
	size := 1024
	option := New(OptProducerBufferSize(size))
	assert.Equal(t, option.producerConfig.bufferSize, size)
}

func TestOptProducerConcurrency(t *testing.T) {
	concurrency := 10
	option := New(OptProducerConcurrency(concurrency))
	assert.Equal(t, option.producerConfig.concurrency, concurrency)
}

func TestOptProducerConfig(t *testing.T) {
	runner := &TestRunner{}
	dispatcher := New(OptProducerConfig(func() Producer {
		return runner
	}))
	assert.Equal(t, dispatcher.producerConfig.producer(), runner)

}
