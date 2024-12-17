package tests

import (
	paginator2 "github.com/gouef/paginator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func test1(t *testing.T) {
	paginator := paginator2.NewPaginator()
	paginator.SetItemCount(7)
	paginator.SetItemsPerPage(6)
	paginator.SetBase(0)
	paginator.SetPage(3)
	assert.Same(t, 1, paginator.GetPage())
	assert.Same(t, 2, paginator.GetPageCount())
	assert.Same(t, 0, paginator.GetFirstPage())
	assert.Same(t, 1, paginator.GetLastPage())
	assert.Same(t, 7, paginator.GetFirstItemOnPage())
	assert.Same(t, 7, paginator.GetLastItemOnPage())
	assert.Same(t, 6, paginator.GetOffset())
	assert.Same(t, 0, paginator.GetCountdownOffset())
	assert.Same(t, 1, paginator.GetLength())
}

func test2(t *testing.T) {
	paginator := paginator2.NewPaginator()
	paginator.SetItemCount(7)
	paginator.SetItemsPerPage(6)
	paginator.SetBase(0)
	paginator.SetPage(-1)

	assert.Same(t, 0, paginator.GetPage())
	assert.Same(t, 1, paginator.GetFirstItemOnPage())
	assert.Same(t, 6, paginator.GetLastItemOnPage())
	assert.Same(t, 0, paginator.GetOffset())
	assert.Same(t, 1, paginator.GetCountdownOffset())
	assert.Same(t, 6, paginator.GetLength())
}

func test3(t *testing.T) {
	paginator := paginator2.NewPaginator()
	paginator.SetItemCount(7)
	paginator.SetItemsPerPage(7)
	paginator.SetBase(0)
	paginator.SetPage(-1)

	assert.Same(t, 0, paginator.GetPage())
	assert.Same(t, 1, paginator.GetPageCount())
	assert.Same(t, 0, paginator.GetFirstPage())
	assert.Same(t, 0, paginator.GetLastPage())
	assert.Same(t, 1, paginator.GetFirstItemOnPage())
	assert.Same(t, 7, paginator.GetLastItemOnPage())
	assert.Same(t, 0, paginator.GetOffset())
	assert.Same(t, 0, paginator.GetCountdownOffset())
	assert.Same(t, 7, paginator.GetLength())
}

func test4(t *testing.T) {
	paginator := paginator2.NewPaginator()
	paginator.SetItemCount(-1)
	paginator.SetItemsPerPage(7)
	paginator.SetBase(0)
	paginator.SetPage(-1)

	assert.Same(t, 0, paginator.GetPage())
	assert.Same(t, 0, paginator.GetPageCount())
	assert.Same(t, 0, paginator.GetFirstPage())
	assert.Same(t, 0, paginator.GetLastPage())
	assert.Same(t, 0, paginator.GetFirstItemOnPage())
	assert.Same(t, 0, paginator.GetLastItemOnPage())
	assert.Same(t, 0, paginator.GetOffset())
	assert.Same(t, 0, paginator.GetCountdownOffset())
	assert.Same(t, 0, paginator.GetLength())
}

func test5(t *testing.T) {
	paginator := paginator2.NewPaginator()
	paginator.SetItemCount(7)
	paginator.SetItemsPerPage(6)
	paginator.SetBase(1)
	paginator.SetPage(3)

	assert.Same(t, 2, paginator.GetPage())
	assert.Same(t, 2, paginator.GetPageCount())
	assert.Same(t, 1, paginator.GetFirstPage())
	assert.Same(t, 2, paginator.GetLastPage())
	assert.Same(t, 7, paginator.GetFirstItemOnPage())
	assert.Same(t, 7, paginator.GetLastItemOnPage())
	assert.Same(t, 6, paginator.GetOffset())
	assert.Same(t, 0, paginator.GetCountdownOffset())
	assert.Same(t, 1, paginator.GetLength())
}

func test6(t *testing.T) {
	paginator := paginator2.NewPaginator()

	// ItemCount: 0
	paginator.SetItemCount(0)
	assert.True(t, paginator.IsFirst())
	assert.True(t, paginator.IsLast())
	assert.Same(t, 0, paginator.GetFirstItemOnPage())
	assert.Same(t, 0, paginator.GetLastItemOnPage())

	// ItemCount: 1
	paginator.SetItemCount(1)
	assert.True(t, paginator.IsFirst())
	assert.True(t, paginator.IsLast())
	assert.Same(t, 1, paginator.GetFirstItemOnPage())
	assert.Same(t, 1, paginator.GetLastItemOnPage())

	// ItemCount: 2
	paginator.SetItemCount(2)
	assert.True(t, paginator.IsFirst())
	assert.False(t, paginator.IsLast())
	assert.Same(t, 1, paginator.GetFirstItemOnPage())
	assert.Same(t, 1, paginator.GetLastItemOnPage())

	// Page 2
	paginator.SetPage(2)
	assert.False(t, paginator.IsFirst())
	assert.True(t, paginator.IsLast())
	assert.Same(t, 2, paginator.GetFirstItemOnPage())
	assert.Same(t, 2, paginator.GetLastItemOnPage())
}

func test7(t *testing.T) {
	paginator := paginator2.NewPaginator()
	paginator.SetItemsPerPage(6)
	paginator.SetBase(0)
	paginator.SetPage(3)

	assert.Same(t, 3, paginator.GetPage())
	assert.Nil(t, paginator.GetPageCount())
	assert.Same(t, 0, paginator.GetFirstPage())
	assert.Nil(t, paginator.GetLastPage())
	assert.Same(t, 19, paginator.GetFirstItemOnPage())
	assert.Same(t, 24, paginator.GetLastItemOnPage())
	assert.Same(t, 18, paginator.GetOffset())
	assert.Nil(t, paginator.GetCountdownOffset())
	assert.Same(t, 6, paginator.GetLength())
}
