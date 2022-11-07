package pointers_test

import (
	"testing"

	"github.com/MohammadBnei/go-exercice/pointers"
)

func TestSizeAndPush(t *testing.T) {
	list := pointers.NewLinkedList()

	list.Push("hey")
	list.Push(123)

	size, _ := list.Size()

	if size != 2 {
		t.Logf("Error on size. Should have 2 elements, have %v", size)
		t.Fail()
	}
}

func TestPushAndGet(t *testing.T) {
	list := pointers.NewLinkedList()

	list.Push("hey")
	list.Push(123)

	second, _ := list.Get(0)
	hey, _ := list.Get(1)
	errNode, err := list.Get(33)

	if hey.Data != "hey" {
		t.Logf("Error on get. Should have \"hey\", have %v", hey.Data)
		t.Fail()
	}
	if second.Data != 123 {
		t.Logf("Error on get. Should have 123, have %v", second.Data)
		t.Fail()
	}

	if err == nil {
		t.Logf("Error on get. Should have an out of bounds error, have %v", errNode)
		t.Fail()
	}
}

func TestInsert(t *testing.T) {
	list := pointers.NewLinkedList()

	list.Push("hey")
	list.Push(123)

	list.Insert("mamamia", 1)

	mia, _ := list.Get(1)
	hey, _ := list.Get(2)
	numbers, _ := list.Get(0)

	if mia.Data != "mamamia" {
		t.Logf("Error on get. Should have \"mamamia\", have %v", mia.Data)
		t.Fail()
	}
	if hey.Data != "hey" {
		t.Logf("Error on get. Should have \"hey\", have %v", hey.Data)
		t.Fail()
	}
	if numbers.Data != 123 {
		t.Logf("Error on get. Should have 123, have %v", numbers.Data)
		t.Fail()
	}

	err := list.Insert("luigi", 4)
	if err == nil {
		t.Log("Error on insert. Should return an error on index 4 out of bound")
		t.Fail()
	}
	list.Insert("luigi", 3)
	if mario, _ := list.Get(3); mario.Data != "luigi" {
		t.Logf("Error on insert. Should have 123, have %v", mario.Data)
		t.Fail()
	}

	list.Insert("henry", 0)
	if arsenal, _ := list.Get(0); arsenal.Data != "henry" {
		t.Logf("Error on insert. Should have henry, have %v", arsenal.Data)
		t.Fail()
	}

}

func TestSearch(t *testing.T) {
	list := pointers.NewLinkedList()

	list.Push("hey")
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push("peter")
	found, err := list.Search(2)
	if err != nil {
		t.Log("Error on search. Should have found node with data 2")
		t.Fail()
	}
	if found.Data != 2 {
		t.Logf("Error on search. Should have 2, have %v", found.Data)
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	list := pointers.NewLinkedList()

	list.Push("hey")
	list.Push(123)
	list.Push("mamamia")

	removed, _ := list.Remove(0)

	hey, _ := list.Get(1)
	numbers, _ := list.Get(0)

	if found, err := list.Search(removed.Data); err == nil {
		t.Logf("Error on removal. Should not have a node with data %s", found.Data)
		t.Fail()
	}

	if hey.Data != "hey" {
		t.Logf("Error on get. Should have \"hey\", have %v", hey.Data)
		t.Fail()
	}
	if numbers.Data != 123 {
		t.Logf("Error on get. Should have 123, have %v", numbers.Data)
		t.Fail()
	}

	list.Push("mamamia")
	removed, _ = list.Remove(2)

	if found, err := list.Search(removed.Data); err == nil {
		t.Logf("Error on removal. Should not have a node with data %s", found.Data)
		t.Fail()
	}

	numbers, _ = list.Get(1)
	mamamia, _ := list.Get(0)
	if mamamia.Data != "mamamia" {
		t.Logf("Error on get. Should have \"mamamia\", have %v", mamamia.Data)
		t.Fail()
	}
	if numbers.Data != 123 {
		t.Logf("Error on get. Should have 123, have %v", numbers.Data)
		t.Fail()
	}
}
func TestRemoveEmptyList(t *testing.T) {
	list := pointers.NewLinkedList()

	_, err := list.Remove(2)

	if err == nil {
		t.Logf("Error on remove. Should have an error on empty list removal")
		t.Fail()
	}
}
