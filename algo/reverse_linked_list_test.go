package algo

import "testing"

func TestNewListHeadIsNil(t *testing.T) {
	list := LinkedList{}
	if list.head != nil {
		t.Fail()
	}
}

func TestAddValues(t *testing.T) {
	list := LinkedList{}
	list.Add(1)

	if list.head == nil {
		t.Error("head is nil")
	}

	if list.head.value != 1 {
		t.Error("value is not correct")
	}

	list.Add(2)

	if list.head.next == nil {
		t.Error("second node is nil")
	}

	if list.head.next.value != 2 {
		t.Error("Value is not correct")
	}
	list.Add(3)

	if list.head.next.next == nil {
		t.Error("second node is nil")
	}

	if list.head.next.next.value != 3 {
		t.Error("Value is not correct")
	}
}

func TestReverse(t *testing.T) {
	list := LinkedList{}
	list.Reverse()
	if list.head != nil {
		t.Error("head is nil")
	}

	list.Add(1)
	list.Add(2)

	list.Reverse()

	if list.head.value != 2 {
		t.Errorf("1,2.Reverse() = %v,%v, wants 2,1",list.head.value, list.head.next.value)
	}
}
