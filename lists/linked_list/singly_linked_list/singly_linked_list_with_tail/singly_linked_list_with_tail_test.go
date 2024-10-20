package linkedlistwithtail

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	utils "github.com/gyuudon3187/go-data-structures-and-algorithms/test_utils"
	"os"
	"testing"
)

var items = []interface{}{1, "string", 0.4}

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

func TestMain(m *testing.M) {
	if len(items) < 2 {
		fmt.Printf("The global variable 'items' must contain at least 2 elements and preferably 3, but its current length is %d.", len(items))
		os.Exit(1)
	}

	m.Run()
}

func TestPrepend(t *testing.T) {
	t.Run("Prepends items", testCase(func(t *testing.T, c *testContext) {
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

	t.Run("Tail points to the prepended item", testCase(func(t *testing.T, c *testContext) {
		got := c.linkedList.tail.item
		want := items[0]
		utils.ValidateResult(t, got, want)
	}))
}

func TestAppend(t *testing.T) {
	t.Run("Appends items", testCase(func(t *testing.T, c *testContext) {
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

	t.Run("Tail points to the appended item", testCase(func(t *testing.T, c *testContext) {
		randomFloat := 0.5
		c.linkedList.Append(randomFloat)

		got := c.linkedList.tail.item
		want := randomFloat
		utils.ValidateResult(t, got, want)
	}))
}

func TestRemoveHead(t *testing.T) {
	t.Run("Returns the head", testCase(func(t *testing.T, c *testContext) {
		got := c.linkedList.RemoveHead()
		want := items[c.itemsLastIndex]
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Sets the head to its prev pointer", testCase(func(t *testing.T, c *testContext) {
		want := c.linkedList.head.prev
		c.linkedList.RemoveHead()
		got := c.linkedList.head
		utils.ValidateResult(t, got, want)
	}))
}

func TestRemoveTail(t *testing.T) {
	t.Run("Returns the tail", testCase(func(t *testing.T, c *testContext) {
		got := c.linkedList.RemoveTail()
		want := items[0]
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Sets the prev 'pointer' of the item next to tail to nil", testCase(func(t *testing.T, c *testContext) {
		nextAfterTail := c.linkedList.head

		for nextAfterTail.prev != c.linkedList.tail {
			nextAfterTail = nextAfterTail.prev
		}

		c.linkedList.RemoveTail()

		got := nextAfterTail.prev
		var want *linkedListItem = nil
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Sets the tail to the item next to old tail", testCase(func(t *testing.T, c *testContext) {
		c.linkedList.RemoveTail()
		got := c.linkedList.tail.item
		want := items[1]
		utils.ValidateResult(t, got, want)
	}))
}

func TestFind(t *testing.T) {
	t.Run("Returns the sought node", testCase(func(t *testing.T, c *testContext) {
		soughtNode := c.linkedList.Find(items[1])
		var got interface{}

		if soughtNode != nil {
			got = soughtNode.item
		}

		want := items[1]
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Is idempotent", testCase(func(t *testing.T, c *testContext) {
		c.linkedList.Find(items[1])

		var itemsAfterFind []interface{} = make([]interface{}, 0, len(items))
		current := c.linkedList.head

		for current != nil {
			itemsAfterFind = append([]interface{}{current.item}, itemsAfterFind...)
			current = current.prev
		}

		if !cmp.Equal(items, itemsAfterFind) {
			t.Errorf("Expected 'items' and 'itemsAfterFind' to be equal but they were not. items: %v, itemsAfterFind: %v", items, itemsAfterFind)
		}
	}))
}
