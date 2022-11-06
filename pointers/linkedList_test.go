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
		t.Errorf("Error on size. Should have 2 elements, have %v", size)
	}
}

func TestPushAndGet(t *testing.T) {
	list := pointers.NewLinkedList()

	list.Push("hey")
	list.Push(123)

	second, _ := list.Get(0)
	hey, _ := list.Get(1)

	if hey.Data != "hey" {
		t.Errorf("Error on get. Should have \"hey\", have %v", hey.Data)
	}
	if second.Data != 123 {
		t.Errorf("Error on get. Should have 123, have %v", second.Data)
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
		t.Errorf("Error on get. Should have \"mamamia\", have %v", mia.Data)
	}
	if hey.Data != "hey" {
		t.Errorf("Error on get. Should have \"hey\", have %v", hey.Data)
	}
	if numbers.Data != 123 {
		t.Errorf("Error on get. Should have 123, have %v", numbers.Data)
	}

	err := list.Insert("luigi", 4)
	if err == nil {
		t.Log("Error on insert. Should return an error on index 4 out of bound")
	}
	list.Insert("luigi", 3)
	if mario, _ := list.Get(3); mario.Data != "luigi" {
		t.Errorf("Error on insert. Should have 123, have %v", mario.Data)
	}

	list.Insert("henry", 0)
	if arsenal, _ := list.Get(0); arsenal.Data != "henry" {
		t.Errorf("Error on insert. Should have henry, have %v", arsenal.Data)
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
	}
	if found.Data != 2 {
		t.Errorf("Error on search. Should have 2, have %v", found.Data)
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
		t.Errorf("Error on removal. Should not have a node with data %s", found.Data)
	}

	if hey.Data != "hey" {
		t.Errorf("Error on get. Should have \"hey\", have %v", hey.Data)
	}
	if numbers.Data != 123 {
		t.Errorf("Error on get. Should have 123, have %v", numbers.Data)
	}

	list.Push("mamamia")
	removed, _ = list.Remove(2)

	if found, err := list.Search(removed.Data); err == nil {
		t.Errorf("Error on removal. Should not have a node with data %s", found.Data)
	}

	numbers, _ = list.Get(1)
	mamamia, _ := list.Get(0)
	if mamamia.Data != "mamamia" {
		t.Errorf("Error on get. Should have \"mamamia\", have %v", mamamia.Data)
	}
	if numbers.Data != 123 {
		t.Errorf("Error on get. Should have 123, have %v", numbers.Data)
	}
}
func TestRemoveEmptyList(t *testing.T) {
	list := pointers.NewLinkedList()

	_, err := list.Remove(2)

	if err == nil {
		t.Errorf("Error on remove. Should have an error on empty list removal")
	}
}
