package paginator

import "math"

type Paginator struct {
	page            int
	firstPage       int
	lastPage        *int
	firstItemOnPage int
	lastItemOnPage  int
	base            int
	first           bool
	last            bool
	pageCount       int
	itemsPerPage    int
	itemCount       *int
	offset          int
	countdownOffset int
	length          int
}

func NewPaginator() *Paginator {
	return &Paginator{
		base:         1,
		itemsPerPage: 1,
		page:         1,
		itemCount:    nil,
	}
}

func (p *Paginator) SetPage(page int) *Paginator {
	p.page = page
	return p
}

func (p *Paginator) GetPage() int {
	return p.page
}

func (p *Paginator) GetFirstPage() int {
	return p.base
}

func (p *Paginator) GetLastPage() *int {
	if p.GetItemCount() == nil {
		return nil
	}

	lastPage := p.base + GetLastPage(p) - 1
	return &lastPage
}

func (p *Paginator) GetFirstItemOnPage() int {
	if *p.GetItemCount() != 0 {
		i := p.offset + 1
		return i
	}

	return 0
}

func (p *Paginator) GetLastItemOnPage() int {
	return p.offset + p.length
}

func (p *Paginator) SetBase(base int) *Paginator {
	p.base = base
	return p
}

func (p *Paginator) GetBase() int {
	return p.base
}

func (p *Paginator) GetPageIndex() int {
	index := math.Max(0, float64(p.page-p.base))

	if p.GetItemCount() == nil {
		return int(index)
	}

	return int(math.Min(index, math.Max(0, float64(*p.GetPageCount()-1))))
}

func (p *Paginator) IsFirst() bool {
	return p.GetPageIndex() == 0
}

func (p *Paginator) IsLast() bool {
	if p.GetItemCount() == nil {
		return false
	}

	return p.GetPageIndex() >= *p.GetPageCount()-1
}

func (p *Paginator) GetPageCount() *int {
	if p.GetItemCount() == nil {
		return nil
	}

	count := GetLastPage(p)

	return &count
}

func (p *Paginator) SetItemsPerPage(itemsPerPage int) *Paginator {
	p.itemsPerPage = max(1, itemsPerPage)
	return p
}

func (p *Paginator) GetItemsPerPage() int {
	return p.itemsPerPage
}

func (p *Paginator) SetItemCount(itemCount int) *Paginator {
	if itemCount < 0 {
		itemCount = 0
	}
	p.itemCount = &itemCount

	return p
}

func (p *Paginator) GetItemCount() *int {
	return p.itemCount
}

func (p *Paginator) GetOffset() int {
	return p.GetPageIndex() * p.GetItemsPerPage()
}

func (p *Paginator) GetCountdownOffset() *int {
	if p.GetItemCount() == nil {
		return nil
	}
	val := max(0, *p.GetItemCount()-(p.GetPageIndex()-1)*p.GetItemsPerPage())
	return &val
}
func (p *Paginator) GetLength() int {

	if p.GetItemCount() == nil {
		return p.GetItemsPerPage()
	}

	val := min(p.GetItemsPerPage(), *p.GetItemCount()-p.GetPageIndex()*p.GetItemsPerPage())

	return val
}

func GetLastPage(paginator *Paginator) int {
	return int(math.Ceil(float64(*paginator.GetItemCount()) / float64(paginator.itemsPerPage)))
}
