package topics

import (
	"time"

	"github.com/FlameInTheDark/ndms-test/internal/queue"
)

type Topics struct {
	topics map[string]*queue.Queue
}

func NewTopics() *Topics {
	return &Topics{
		topics: make(map[string]*queue.Queue),
	}
}

// warmup creates topic if not exists
func (t *Topics) warmup(topic string) {
	if _, ok := t.topics[topic]; !ok {
		t.topics[topic] = queue.NewQueue()
	}
}

func (t *Topics) Produce(topic string, data []byte) {
	t.warmup(topic)

	q := t.topics[topic]
	q.Push(data)
}

func (t *Topics) Consume(topic string, timeout *time.Duration) []byte {
	t.warmup(topic)

	q := t.topics[topic]
	if timeout != nil {
		return q.PopWait(*timeout)
	}
	return q.Pop()
}
