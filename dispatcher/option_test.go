package dispatcher

import (
	"reflect"
	"testing"
)

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
				t.Errorf("OptConsumerBufferSize() = %bufferSize, want %bufferSize", got, tt.want)
			}
		})
	}
}

func TestOptConsumerConcurrency(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OptConsumerConcurrency(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OptConsumerConcurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptConsumerConfig(t *testing.T) {
	type args struct {
		option ConsumerOption
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OptConsumerConfig(tt.args.option); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OptConsumerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptProducerBufferSize(t *testing.T) {
	type args struct {
		s int
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OptProducerBufferSize(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OptProducerBufferSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptProducerConcurrency(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OptProducerConcurrency(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OptProducerConcurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptProducerConfig(t *testing.T) {
	type args struct {
		option ProducerOption
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OptProducerConfig(tt.args.option); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OptProducerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
