package dispatcher

import (
	"context"
	"fmt"
	"testing"
)

type TestConfig struct {
	name string
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
	// TODO: Add test cases.
}

func TestDispatcher_GetSendMessageCount(t *testing.T) {
	// TODO: Add test cases.
}

func TestDispatcher_Start(t *testing.T) {
	// TODO: Add test cases.
}

func TestNew(t *testing.T) {
	// TODO
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
