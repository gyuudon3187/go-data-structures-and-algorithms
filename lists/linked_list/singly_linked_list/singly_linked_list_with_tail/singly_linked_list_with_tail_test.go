package linkedlistwithtail

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	utils "github.com/gyuudon3187/go-data-structures-and-algorithms/test_utils"
	"os"
	"testing"
)

var items = []interface{}{1, "string", 0.4, "another string"}

type TestContext interface {
	beforeEach()
}

type testContext struct {
	linkedList     *linkedList
	itemsLastIndex int
}

func (c *testContext) beforeEach() {
	l := New()

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
	if len(items) < 3 {
		fmt.Printf("The global variable 'items' must contain at least 3 elements and preferably 4, but its current length is %d.", len(items))
		os.Exit(1)
	}

	m.Run()
}

func TestPrepend(t *testing.T) {
	t.Run("Prepends items", testCase(func(t *testing.T, tc *testContext) {
		nthLinkedListItem := tc.linkedList.head
		var got, want interface{}

		for i := 0; i < len(items); i++ {
			got = nthLinkedListItem.item
			want = items[tc.itemsLastIndex-i]
			utils.ValidateResult(t, got, want)
			nthLinkedListItem = nthLinkedListItem.next
		}

		if nthLinkedListItem != nil {
			t.Errorf("Expected nthLinkedListItem to be nil but got %v", nthLinkedListItem)
		}
	}))

	t.Run("Tail points to the prepended item", testCase(func(t *testing.T, tc *testContext) {
		got := tc.linkedList.tail.item
		want := items[0]
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Increments length", testCase(func(t *testing.T, tc *testContext) {
		got := tc.linkedList.Length()
		want := len(items)
		utils.ValidateResult(t, got, want)
	}))
}

func TestAppend(t *testing.T) {
	t.Run("Appends items", testCase(func(t *testing.T, tc *testContext) {
		randomFloat := 0.5
		tc.linkedList.Append(randomFloat)
		current := tc.linkedList.head
		var lastItem *node

		for current != nil {
			lastItem = current
			current = current.next
		}

		got := lastItem.item
		want := randomFloat
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Tail points to the appended item", testCase(func(t *testing.T, tc *testContext) {
		randomFloat := 0.5
		tc.linkedList.Append(randomFloat)

		got := tc.linkedList.tail.item
		want := randomFloat
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Increments length", testCase(func(t *testing.T, tc *testContext) {
		tc.linkedList.Append("random string")
		got := tc.linkedList.Length()
		want := len(items) + 1
		utils.ValidateResult(t, got, want)
	}))
}

func TestRemoveHead(t *testing.T) {
	t.Run("Returns the head", testCase(func(t *testing.T, tc *testContext) {
		got := tc.linkedList.RemoveHead()
		want := items[tc.itemsLastIndex]
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Sets the head to tail when only one item remains after removal", func(t *testing.T) {
		linkedList := New()
		linkedList.Prepend("foo")
		linkedList.Prepend("bar")
		linkedList.RemoveHead()
		got := linkedList.head
		want := linkedList.tail
		utils.ValidateResult(t, got, want)
	})

	t.Run("Sets the head to its 'next' pointer", testCase(func(t *testing.T, tc *testContext) {
		want := tc.linkedList.head.next
		tc.linkedList.RemoveHead()
		got := tc.linkedList.head
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Decrements length", testCase(func(t *testing.T, tc *testContext) {
		tc.linkedList.RemoveHead()
		got := tc.linkedList.Length()
		want := len(items) - 1
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Does not decrement below 0", func(t *testing.T) {
		linkedList := New()
		linkedList.RemoveHead()

		got := linkedList.Length()
		want := 0
		utils.ValidateResult(t, got, want)
	})
}

func TestRemoveTail(t *testing.T) {
	t.Run("Returns the tail", testCase(func(t *testing.T, tc *testContext) {
		got := tc.linkedList.RemoveTail()
		want := items[0]
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Sets the head to tail when only one item remains after removal", func(t *testing.T) {
		linkedList := New()
		linkedList.Prepend("foo")
		linkedList.Prepend("bar")
		linkedList.RemoveTail()
		got := linkedList.head
		want := linkedList.tail
		utils.ValidateResult(t, got, want)
	})

	t.Run("Sets the 'next' pointer of the item before tail to nil", testCase(func(t *testing.T, tc *testContext) {
		nextAfterTail := tc.linkedList.head

		for nextAfterTail.next != tc.linkedList.tail {
			nextAfterTail = nextAfterTail.next
		}

		tc.linkedList.RemoveTail()

		got := nextAfterTail.next
		var want *node = nil
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Sets the tail to the item before old tail", testCase(func(t *testing.T, tc *testContext) {
		tc.linkedList.RemoveTail()
		got := tc.linkedList.tail.item
		want := items[1]
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Decrements length", testCase(func(t *testing.T, tc *testContext) {
		tc.linkedList.RemoveTail()
		got := tc.linkedList.Length()
		want := len(items) - 1
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Does not decrement below 0", func(t *testing.T) {
		linkedList := New()
		linkedList.RemoveTail()

		got := linkedList.Length()
		want := 0
		utils.ValidateResult(t, got, want)
	})
}

func TestRemoveAt(t *testing.T) {
	t.Run("Returns removed element", testCase(func(t *testing.T, tc *testContext) {
		for i := range items {
			got, err := tc.linkedList.RemoveAt(0)

			if err != nil {
				t.Errorf("Could not remove item: %s", err.Error())
			}

			want := items[tc.itemsLastIndex-i]
			utils.ValidateResult(t, got, want)
		}
	}))

	t.Run("Removes first item", testCase(func(t *testing.T, tc *testContext) {
		tc.linkedList.RemoveAt(0)
		got := tc.linkedList.head.item
		want := items[tc.itemsLastIndex-1]
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Removes intermediate item", testCase(func(t *testing.T, tc *testContext) {
		tc.linkedList.RemoveAt(1)
		got := tc.linkedList.head.next.item
		want := items[tc.itemsLastIndex-2]
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Removes last item", testCase(func(t *testing.T, tc *testContext) {
		tc.linkedList.RemoveAt(tc.itemsLastIndex)
		got := tc.linkedList.tail.item
		want := items[1]
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Throws error when negative bounds", testCase(func(t *testing.T, tc *testContext) {
		_, err := tc.linkedList.RemoveAt(-1)
		if err == nil {
			t.Error("Expected negative bounds to throw error but it didn't")
		}
	}))

	t.Run("Throws error when index exceeds upper bound", testCase(func(t *testing.T, tc *testContext) {
		_, err := tc.linkedList.RemoveAt(tc.itemsLastIndex + 1)
		if err == nil {
			t.Error("Expected index exceeding upper bound to throw error but it didn't")
		}
	}))
}

func TestFind(t *testing.T) {
	t.Run("Returns the sought node", testCase(func(t *testing.T, tc *testContext) {
		soughtNode := tc.linkedList.Find(items[1])
		var got interface{}

		if soughtNode != nil {
			got = soughtNode.item
		}

		want := items[1]
		utils.ValidateResult(t, got, want)
	}))

	t.Run("Is idempotent", testCase(func(t *testing.T, tc *testContext) {
		tc.linkedList.Find(items[1])

		var itemsAfterFind []interface{} = make([]interface{}, 0, len(items))
		current := tc.linkedList.head

		for current != nil {
			itemsAfterFind = append([]interface{}{current.item}, itemsAfterFind...)
			current = current.next
		}

		if !cmp.Equal(items, itemsAfterFind) {
			t.Errorf("Expected 'items' and 'itemsAfterFind' to be equal but they were not. items: %v, itemsAfterFind: %v", items, itemsAfterFind)
		}
	}))
}

func TestIsEmpty(t *testing.T) {
	t.Run("True when empty", func(t *testing.T) {
		linkedList := new(linkedList)
		got := linkedList.IsEmpty()
		want := true
		utils.ValidateResult(t, got, want)
	})

	t.Run("False when not empty", testCase(func(t *testing.T, tc *testContext) {
		got := tc.linkedList.IsEmpty()
		want := false
		utils.ValidateResult(t, got, want)
	}))
}

func TestIterate(t *testing.T) {
	t.Run("Performs given callback: collects each element into slice", testCase(func(t *testing.T, tc *testContext) {
		got := []interface{}{}

		tc.linkedList.Iterate(func(item interface{}) {
			got = append([]interface{}{item}, got...)
		})

		want := items

		if !cmp.Equal(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}))
}
