package linkedlist

import (
	utils "github.com/gyuudon3187/go-data-structures-and-algorithms/test_utils"
	"testing"
)

var items = []interface{}{1, "string"}

type TestContext interface {
	beforeEach()
}

type testContext struct {
	linkedlist     *linkedList
	itemsLastIndex int
}

func (c *testContext) beforeEach() {
	l := new(linkedList)

	for _, item := range items {
		l.Prepend(item)
	}

	c.linkedlist = l

	c.itemsLastIndex = len(items) - 1
}

func testCase(test func(*testing.T, *testContext)) func(*testing.T) {
	return func(t *testing.T) {
		context := &testContext{}
		context.beforeEach()
		test(t, context)
	}
}

func TestPrepend(t *testing.T) {
	t.Run("Prepend adds items first-in", testCase(func(t *testing.T, c *testContext) {
		nthLinkedListItem := c.linkedlist.head
		var got, want interface{}

		for i := 0; i < len(items); i++ {
			got = nthLinkedListItem.item
			want = items[c.itemsLastIndex-i]
			utils.ValidateResult(t, got, want)
			nthLinkedListItem = nthLinkedListItem.prev
		}

		if nthLinkedListItem != nil {
			t.Errorf("Expected nthLinkedListItem to be nil but got %v", nthLinkedListItem)
		}
	}))
}
