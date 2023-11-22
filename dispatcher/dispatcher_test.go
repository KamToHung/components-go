package dispatcher

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

type TestConfig struct {
	name string
}

func TestDispatcher_GetConsumeMessageCount(t *testing.T) {
	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx)
	// 创建一个通道来接收操作系统信号
	sigs := make(chan os.Signal, 1)
	// 注册感兴趣的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 启动一个goroutine来等待信号
	go func() {
		sig := <-sigs
		fmt.Println("接收到信号:", sig)
		cancel()
		os.Exit(0)
	}()
	d := New(
		OptProducerConfig(func() Producer {
			return &TestRunner{}
		}),
		OptProducerConcurrency(20),
		OptProducerBufferSize(1024),
		OptConsumerConfig(func() Consumer {
			return &TestRunner{}
		}),
		OptProducerConcurrency(10),
		OptProducerBufferSize(1024),
	)
	d.Start(ctx, &TestConfig{name: "test"})
	assert.Equal(t, d.GetConsumeMessageCount(), 0)
}

func TestDispatcher_GetSendMessageCount(t *testing.T) {
	// TODO: Add test cases.
}

func TestDispatcher_Start(t *testing.T) {
	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx)
	// 创建一个通道来接收操作系统信号
	sigs := make(chan os.Signal, 1)
	// 注册感兴趣的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 启动一个goroutine来等待信号
	go func() {
		sig := <-sigs
		fmt.Println("接收到信号:", sig)
		cancel()
		os.Exit(0)
	}()
	d := New(
		OptProducerConfig(func() Producer {
			return &TestRunner{}
		}),
		OptProducerConcurrency(20),
		OptProducerBufferSize(1024),
		OptConsumerConfig(func() Consumer {
			return &TestRunner{}
		}),
		OptProducerConcurrency(10),
		OptProducerBufferSize(1024),
	)
	d.Start(ctx, TestConfig{name: "test"}, TestConfig{name: "test2"})
}

func TestNew(t *testing.T) {

}

func Test_calculateIndex(t *testing.T) {
	// TODO: Add test cases.
}

func Test_consumerProcess(t *testing.T) {
	// TODO: Add test cases.
}

func Test_producerProcess(t *testing.T) {
	// TODO: Add test cases.
}
