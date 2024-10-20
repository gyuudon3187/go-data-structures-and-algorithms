package linkedlistwithtail

import (
	utils "github.com/gyuudon3187/go-data-structures-and-algorithms/test_utils"
	"testing"
)

var items = []interface{}{1, "string"}

type TestContext interface {
	beforeEach()
}

type testContext struct {
	linkedList     *linkedList
	itemsLastIndex int
}

func (c *testContext) beforeEach() {
	l := new(linkedList)

	for _, item := range items {
		l.Prepend(item)
	}

	c.linkedList = l

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
		nthLinkedListItem := c.linkedList.head
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

	t.Run("Tail points to the item added first by Prepend", testCase(func(t *testing.T, c *testContext) {
		got := c.linkedList.tail.item
		want := items[0]
		utils.ValidateResult(t, got, want)
	}))
}

func TestAppend(t *testing.T) {
	t.Run("Append adds items last-in", testCase(func(t *testing.T, c *testContext) {
		randomFloat := 0.5
		c.linkedList.Append(randomFloat)
		current := c.linkedList.head
		var lastItem *linkedListItem

		for current != nil {
			lastItem = current
			current = current.prev
		}

		got := lastItem.item
		want := randomFloat
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Tail points to the item added last by Append", testCase(func(t *testing.T, c *testContext) {
		randomFloat := 0.5
		c.linkedList.Append(randomFloat)

		got := c.linkedList.tail.item
		want := randomFloat
		utils.ValidateResult(t, got, want)
	}))
}