package queue

import (
	utils "github.com/gyuudon3187/go-data-structures-and-algorithms/test_utils"
	"testing"
)

var items = []interface{}{1, "string"}

type TestContext interface {
	beforeEach()
}

type testContext struct {
	queue          *queue
	itemsLastIndex int
}

func (c *testContext) beforeEach() {
	q := new(queue)

	for _, item := range items {
		q.Enqueue(item)
	}

	c.queue = q

	c.itemsLastIndex = len(items) - 1
}

func testCase(test func(*testing.T, *testContext)) func(*testing.T) {
	return func(t *testing.T) {
		context := &testContext{}
		context.beforeEach()
		test(t, context)
	}
}

func TestEnqueue(t *testing.T) {
	t.Run("Enqueue adds items in FIFO order", testCase(func(t *testing.T, c *testContext) {
		nthQueueItem := c.queue.first
		var got, want interface{}

		for i := 0; i < len(items); i++ {
			got = nthQueueItem.item
			want = items[i]
			utils.ValidateResult(t, got, want)
			nthQueueItem = nthQueueItem.prev
		}

		if nthQueueItem != nil {
			t.Errorf("Expected nthQueueItem to be nil but got %v", nthQueueItem)
		}
	}))
}

func TestDequeue(t *testing.T) {
	t.Run("Dequeue returns items in FIFO order", testCase(func(t *testing.T, c *testContext) {
		var got, want interface{}

		for i := 0; i < c.itemsLastIndex; i++ {
			got = c.queue.Dequeue()
			want = items[i]
			utils.ValidateResult(t, got, want)
		}
	}))
}
