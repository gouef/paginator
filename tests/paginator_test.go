package tests

import (
	"github.com/gouef/paginator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPaginator_DefaultValues(t *testing.T) {
	p := paginator.NewPaginator()

	assert.Equal(t, 1, p.GetPage())
	assert.Equal(t, 1, p.GetBase())
	assert.Equal(t, 1, p.GetItemsPerPage())
	assert.Nil(t, p.GetItemCount())
}

func TestPaginator_SetPage(t *testing.T) {
	p := paginator.NewPaginator()

	p.SetPage(5)
	assert.Equal(t, 5, p.GetPage())
}

func TestPaginator_SetBase(t *testing.T) {
	p := paginator.NewPaginator()

	p.SetBase(0)
	assert.Equal(t, 0, p.GetBase())
}

func TestPaginator_SetItemsPerPage(t *testing.T) {
	p := paginator.NewPaginator()

	p.SetItemsPerPage(10)
	assert.Equal(t, 10, p.GetItemsPerPage())

	p.SetItemsPerPage(0)
	assert.Equal(t, 1, p.GetItemsPerPage(), "itemsPerPage nesmí být menší než 1")
}

func TestPaginator_SetItemCount(t *testing.T) {
	p := paginator.NewPaginator()

	p.SetItemCount(100)
	assert.Equal(t, 100, *p.GetItemCount())

	p.SetItemCount(-10)
	assert.Equal(t, 0, *p.GetItemCount(), "itemCount nesmí být záporné")
}

func TestPaginator_GetPageCount(t *testing.T) {
	p := paginator.NewPaginator()

	p.SetItemsPerPage(10).SetItemCount(95)
	pageCount := p.GetPageCount()
	require.NotNil(t, pageCount)
	assert.Equal(t, 10, *pageCount) // 95 / 10 = 10 (zaokrouhleno nahoru)

	p.SetItemsPerPage(20).SetItemCount(100)
	assert.Equal(t, 5, *p.GetPageCount())
}

func TestPaginator_GetPageCount_ItemCountNil(t *testing.T) {
	p := paginator.NewPaginator()

	assert.Nil(t, p.GetPageCount(), "PageCount should be nil when itemCount is nil")
}

func TestPaginator_GetFirstPage(t *testing.T) {
	p := paginator.NewPaginator()

	assert.Equal(t, 1, p.GetFirstPage(), "FirstPage should return the default base value 1")

	p.SetBase(5)
	assert.Equal(t, 5, p.GetFirstPage(), "FirstPage should return the updated base value")
}

func TestPaginator_GetLastPage(t *testing.T) {
	p := paginator.NewPaginator()

	p.SetItemsPerPage(10).SetItemCount(95)
	lastPage := p.GetLastPage()
	require.NotNil(t, lastPage)
	assert.Equal(t, 10, *lastPage)

	p.SetItemsPerPage(20).SetItemCount(100)
	assert.Equal(t, 5, *p.GetLastPage())
}

func TestPaginator_GetLastPage_ItemCountNil(t *testing.T) {
	p := paginator.NewPaginator()

	assert.Nil(t, p.GetLastPage(), "LastPage should return nil when itemCount is nil")
}

func TestPaginator_IsFirst(t *testing.T) {
	p := paginator.NewPaginator()

	p.SetPage(1)
	assert.True(t, p.IsFirst())

	p.SetPage(2)
	assert.False(t, p.IsFirst())
}

func TestPaginator_IsLast(t *testing.T) {
	p := paginator.NewPaginator()

	p.SetItemsPerPage(10).SetItemCount(30)
	p.SetPage(3)
	assert.True(t, p.IsLast())

	p.SetPage(2)
	assert.False(t, p.IsLast())
}

func TestPaginator_IsLast_WhenItemCountIsNil(t *testing.T) {
	p := paginator.NewPaginator()

	isLast := p.IsLast()

	assert.False(t, isLast, "Expected IsLast() to return false when itemCount is nil")
}

func TestPaginator_GetOffset(t *testing.T) {
	p := paginator.NewPaginator()

	p.SetPage(3).SetItemsPerPage(10)
	assert.Equal(t, 20, p.GetOffset())
}

func TestPaginator_GetCountdownOffset(t *testing.T) {
	p := paginator.NewPaginator()

	p.SetItemsPerPage(10).SetItemCount(30).SetPage(2)
	countdown := p.GetCountdownOffset()
	require.NotNil(t, countdown)
	assert.Equal(t, 20, *countdown)
}

func TestPaginator_GetCountdownOffset_ItemCountNil(t *testing.T) {
	p := paginator.NewPaginator()

	assert.Nil(t, p.GetCountdownOffset(), "CountdownOffset should be nil when itemCount is nil")
}

func TestPaginator_GetLength(t *testing.T) {
	p := paginator.NewPaginator()

	p.SetItemsPerPage(10).SetItemCount(25).SetPage(3)
	assert.Equal(t, 5, p.GetLength()) // 25 - (2 * 10) = 5

	p.SetPage(2)
	assert.Equal(t, 10, p.GetLength())
}

func TestPaginator_GetLength_ItemCountNil(t *testing.T) {
	p := paginator.NewPaginator()
	p.SetItemsPerPage(10)

	assert.Equal(t, 10, p.GetLength(), "Length should return itemsPerPage when itemCount is nil")
}

func TestPaginator_GetFirstItemOnPage(t *testing.T) {
	p := paginator.NewPaginator()

	p.SetItemsPerPage(10).SetItemCount(100).SetPage(3)
	assert.Equal(t, 21, p.GetFirstItemOnPage()) // offset + 1 = 20 + 1

	p.SetPage(1)
	assert.Equal(t, 1, p.GetFirstItemOnPage())
}

func TestPaginator_GetFirstItemOnPage_ItemCountNil(t *testing.T) {
	p := paginator.NewPaginator()
	p.SetPage(2)

	assert.Equal(t, 0, p.GetFirstItemOnPage(), "FirstItemOnPage should return 0 when itemCount is nil")
}

func TestPaginator_GetLastItemOnPage(t *testing.T) {
	p := paginator.NewPaginator()

	p.SetItemsPerPage(10).SetItemCount(100).SetPage(3)
	assert.Equal(t, 30, p.GetLastItemOnPage()) // offset + length = 20 + 10

	p.SetPage(1)
	assert.Equal(t, 10, p.GetLastItemOnPage())
}

func TestPaginator_GetLastItemOnPage_ItemCountNil(t *testing.T) {
	p := paginator.NewPaginator()
	p.SetPage(2)

	assert.Equal(t, 0, p.GetLastItemOnPage(), "LastItemOnPage should return 0 when itemCount is nil")
}

func test1(t *testing.T) {
	p := paginator.NewPaginator()
	p.SetItemCount(7)
	p.SetItemsPerPage(6)
	p.SetBase(0)
	p.SetPage(3)
	assert.Same(t, 1, p.GetPage())
	assert.Same(t, 2, p.GetPageCount())
	assert.Same(t, 0, p.GetFirstPage())
	assert.Same(t, 1, p.GetLastPage())
	assert.Same(t, 7, p.GetFirstItemOnPage())
	assert.Same(t, 7, p.GetLastItemOnPage())
	assert.Same(t, 6, p.GetOffset())
	assert.Same(t, 0, p.GetCountdownOffset())
	assert.Same(t, 1, p.GetLength())
}

func test2(t *testing.T) {
	p := paginator.NewPaginator()
	p.SetItemCount(7)
	p.SetItemsPerPage(6)
	p.SetBase(0)
	p.SetPage(-1)

	assert.Same(t, 0, p.GetPage())
	assert.Same(t, 1, p.GetFirstItemOnPage())
	assert.Same(t, 6, p.GetLastItemOnPage())
	assert.Same(t, 0, p.GetOffset())
	assert.Same(t, 1, p.GetCountdownOffset())
	assert.Same(t, 6, p.GetLength())
}

func test3(t *testing.T) {
	p := paginator.NewPaginator()
	p.SetItemCount(7)
	p.SetItemsPerPage(7)
	p.SetBase(0)
	p.SetPage(-1)

	assert.Same(t, 0, p.GetPage())
	assert.Same(t, 1, p.GetPageCount())
	assert.Same(t, 0, p.GetFirstPage())
	assert.Same(t, 0, p.GetLastPage())
	assert.Same(t, 1, p.GetFirstItemOnPage())
	assert.Same(t, 7, p.GetLastItemOnPage())
	assert.Same(t, 0, p.GetOffset())
	assert.Same(t, 0, p.GetCountdownOffset())
	assert.Same(t, 7, p.GetLength())
}

func test4(t *testing.T) {
	p := paginator.NewPaginator()
	p.SetItemCount(-1)
	p.SetItemsPerPage(7)
	p.SetBase(0)
	p.SetPage(-1)

	assert.Same(t, 0, p.GetPage())
	assert.Same(t, 0, p.GetPageCount())
	assert.Same(t, 0, p.GetFirstPage())
	assert.Same(t, 0, p.GetLastPage())
	assert.Same(t, 0, p.GetFirstItemOnPage())
	assert.Same(t, 0, p.GetLastItemOnPage())
	assert.Same(t, 0, p.GetOffset())
	assert.Same(t, 0, p.GetCountdownOffset())
	assert.Same(t, 0, p.GetLength())
}

func test5(t *testing.T) {
	p := paginator.NewPaginator()
	p.SetItemCount(7)
	p.SetItemsPerPage(6)
	p.SetBase(1)
	p.SetPage(3)

	assert.Same(t, 2, p.GetPage())
	assert.Same(t, 2, p.GetPageCount())
	assert.Same(t, 1, p.GetFirstPage())
	assert.Same(t, 2, p.GetLastPage())
	assert.Same(t, 7, p.GetFirstItemOnPage())
	assert.Same(t, 7, p.GetLastItemOnPage())
	assert.Same(t, 6, p.GetOffset())
	assert.Same(t, 0, p.GetCountdownOffset())
	assert.Same(t, 1, p.GetLength())
}

func test6(t *testing.T) {
	p := paginator.NewPaginator()

	// ItemCount: 0
	p.SetItemCount(0)
	assert.True(t, p.IsFirst())
	assert.True(t, p.IsLast())
	assert.Same(t, 0, p.GetFirstItemOnPage())
	assert.Same(t, 0, p.GetLastItemOnPage())

	// ItemCount: 1
	p.SetItemCount(1)
	assert.True(t, p.IsFirst())
	assert.True(t, p.IsLast())
	assert.Same(t, 1, p.GetFirstItemOnPage())
	assert.Same(t, 1, p.GetLastItemOnPage())

	// ItemCount: 2
	p.SetItemCount(2)
	assert.True(t, p.IsFirst())
	assert.False(t, p.IsLast())
	assert.Same(t, 1, p.GetFirstItemOnPage())
	assert.Same(t, 1, p.GetLastItemOnPage())

	// Page 2
	p.SetPage(2)
	assert.False(t, p.IsFirst())
	assert.True(t, p.IsLast())
	assert.Same(t, 2, p.GetFirstItemOnPage())
	assert.Same(t, 2, p.GetLastItemOnPage())
}

func test7(t *testing.T) {
	p := paginator.NewPaginator()
	p.SetItemsPerPage(6)
	p.SetBase(0)
	p.SetPage(3)

	assert.Same(t, 3, p.GetPage())
	assert.Nil(t, p.GetPageCount())
	assert.Same(t, 0, p.GetFirstPage())
	assert.Nil(t, p.GetLastPage())
	assert.Same(t, 19, p.GetFirstItemOnPage())
	assert.Same(t, 24, p.GetLastItemOnPage())
	assert.Same(t, 18, p.GetOffset())
	assert.Nil(t, p.GetCountdownOffset())
	assert.Same(t, 6, p.GetLength())
}
