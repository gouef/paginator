package paginator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func test1(t *testing.T) {
	paginator := NewPaginator()
	paginator.setItemCount(7)
	paginator.itemsPerPage = 6
	paginator.base = 0
	paginator.page = 3
	assert.Same(t, 1, paginator.page)
	assert.Same(t, 2, paginator.pageCount)
	assert.Same(t, 0, paginator.firstPage)
	assert.Same(t, 1, paginator.lastPage)
	assert.Same(t, 7, paginator.firstItemOnPage)
	assert.Same(t, 7, paginator.lastItemOnPage)
	assert.Same(t, 6, paginator.offset)
	assert.Same(t, 0, paginator.countdownOffset)
	assert.Same(t, 1, paginator.length)
}

func test2(t *testing.T) {
	paginator := NewPaginator()
	paginator.setItemCount(7)
	paginator.itemsPerPage = 6
	paginator.base = 0
	paginator.page = -1

	assert.Same(t, 0, paginator.page)
	assert.Same(t, 1, paginator.firstItemOnPage)
	assert.Same(t, 6, paginator.lastItemOnPage)
	assert.Same(t, 0, paginator.offset)
	assert.Same(t, 1, paginator.countdownOffset)
	assert.Same(t, 6, paginator.length)
}

func test3(t *testing.T) {
	paginator := NewPaginator()
	paginator.setItemCount(7)
	paginator.itemsPerPage = 7
	paginator.base = 0
	paginator.page = -1

	assert.Same(t, 0, paginator.page)
	assert.Same(t, 1, paginator.pageCount)
	assert.Same(t, 0, paginator.firstPage)
	assert.Same(t, 0, paginator.lastPage)
	assert.Same(t, 1, paginator.firstItemOnPage)
	assert.Same(t, 7, paginator.lastItemOnPage)
	assert.Same(t, 0, paginator.offset)
	assert.Same(t, 0, paginator.countdownOffset)
	assert.Same(t, 7, paginator.length)
}

func test4(t *testing.T) {
	paginator := NewPaginator()
	paginator.setItemCount(-1)
	paginator.itemsPerPage = 7
	paginator.base = 0
	paginator.page = -1

	assert.Same(t, 0, paginator.page)
	assert.Same(t, 0, paginator.pageCount)
	assert.Same(t, 0, paginator.firstPage)
	assert.Same(t, 0, paginator.lastPage)
	assert.Same(t, 0, paginator.firstItemOnPage)
	assert.Same(t, 0, paginator.lastItemOnPage)
	assert.Same(t, 0, paginator.offset)
	assert.Same(t, 0, paginator.countdownOffset)
	assert.Same(t, 0, paginator.length)
}

func test5(t *testing.T) {
	paginator := NewPaginator()
	paginator.setItemCount(7)
	paginator.itemsPerPage = 6
	paginator.base = 1
	paginator.page = 3

	assert.Same(t, 2, paginator.page)
	assert.Same(t, 2, paginator.pageCount)
	assert.Same(t, 1, paginator.firstPage)
	assert.Same(t, 2, paginator.lastPage)
	assert.Same(t, 7, paginator.firstItemOnPage)
	assert.Same(t, 7, paginator.lastItemOnPage)
	assert.Same(t, 6, paginator.offset)
	assert.Same(t, 0, paginator.countdownOffset)
	assert.Same(t, 1, paginator.length)
}

func test6(t *testing.T) {
	paginator := NewPaginator()

	// ItemCount: 0
	paginator.setItemCount(0)
	assert.True(t, paginator.isFirst())
	assert.True(t, paginator.isLast())
	assert.Same(t, 0, paginator.firstItemOnPage)
	assert.Same(t, 0, paginator.lastItemOnPage)

	// ItemCount: 1
	paginator.setItemCount(1)
	assert.True(t, paginator.isFirst())
	assert.True(t, paginator.isLast())
	assert.Same(t, 1, paginator.firstItemOnPage)
	assert.Same(t, 1, paginator.lastItemOnPage)

	// ItemCount: 2
	paginator.setItemCount(2)
	assert.True(t, paginator.isFirst())
	assert.False(t, paginator.isLast())
	assert.Same(t, 1, paginator.firstItemOnPage)
	assert.Same(t, 1, paginator.lastItemOnPage)

	// Page 2
	paginator.setPage(2)
	assert.False(t, paginator.isFirst())
	assert.True(t, paginator.isLast())
	assert.Same(t, 2, paginator.firstItemOnPage)
	assert.Same(t, 2, paginator.lastItemOnPage)
}

func test7(t *testing.T) {
	paginator := NewPaginator()
	paginator.itemsPerPage = 6
	paginator.base = 0
	paginator.page = 3

	assert.Same(t, 3, paginator.page)
	assert.Nil(t, paginator.pageCount)
	assert.Same(t, 0, paginator.firstPage)
	assert.Nil(t, paginator.lastPage)
	assert.Same(t, 19, paginator.firstItemOnPage)
	assert.Same(t, 24, paginator.lastItemOnPage)
	assert.Same(t, 18, paginator.offset)
	assert.Nil(t, paginator.countdownOffset)
	assert.Same(t, 6, paginator.length)
}
