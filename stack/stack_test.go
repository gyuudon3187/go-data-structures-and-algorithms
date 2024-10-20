package stack

import (
	utils "github.com/gyuudon3187/go-data-structures-and-algorithms/test_utils"
	"testing"
)

var items = []interface{}{1, "string"}

type testContext struct {
	stack          *stack
	itemsLastIndex int
}

func (c *testContext) beforeEach() {
	s := new(stack)

	for _, item := range items {
		s.Push(item)
	}

	c.stack = s

	c.itemsLastIndex = len(items) - 1
}

func testCase(test func(*testing.T, *testContext)) func(*testing.T) {
	return func(t *testing.T) {
		context := &testContext{}
		context.beforeEach()
		test(t, context)
	}
}

func TestPush(t *testing.T) {
	t.Run("Push adds items in LIFO order", testCase(func(t *testing.T, c *testContext) {
		nthStackItem := c.stack.sp
		var got, want interface{}

		for i := 0; i < c.itemsLastIndex; i++ {
			got = nthStackItem.item
			want = items[c.itemsLastIndex-i]
			utils.ValidateResult(t, got, want)
			nthStackItem = nthStackItem.next
		}
	}))
}

func TestPop(t *testing.T) {
	t.Run("Pop returns items in LIFO order", testCase(func(t *testing.T, c *testContext) {
		var got, want interface{}

		for i := 0; i < c.itemsLastIndex; i++ {
			got = c.stack.Pop()
			want = items[c.itemsLastIndex-i]
			utils.ValidateResult(t, got, want)
		}
	}))
}

func TestPeek(t *testing.T) {
	t.Run("Peek only returns the latest items", testCase(func(t *testing.T, c *testContext) {
		var got interface{}

		for i := 0; i < 2; i++ {
			got = c.stack.Peek()
		}

		want := items[c.itemsLastIndex]
		utils.ValidateResult(t, got, want)
	}))
}
