package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_InsertHead(t *testing.T) {
	list := New[int]()

	list.InsertHead(1)
	list.InsertHead(2)
	list.InsertHead(3)

	assert.NotNil(t, list.head, "head shouldn't be nil")
	assert.NotNil(t, list.head.next, "head next node shouldn't be nil")
	assert.Nil(t, list.head.prev, "head prev node should be nil")
	assert.Equal(t, list.head.data, 3, "incorrect head data")
}

func TestList_InsertTail(t *testing.T) {

}

func TestList_InsertAt(t *testing.T) {

}
