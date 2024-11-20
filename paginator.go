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

func (p *Paginator) setPage(page int) *Paginator {
	p.page = page
	return p
}

func (p *Paginator) getPage() int {
	return p.page
}

func (p *Paginator) getFirstPage() int {
	return p.base
}

func (p *Paginator) getLastPage() *int {
	if p.getItemCount() == nil {
		return nil
	}

	lastPage := p.base + getLastPage(p) - 1
	return &lastPage
}

func (p *Paginator) getFirstItemOnPage() int {
	if *p.getItemCount() != 0 {
		i := p.offset + 1
		return i
	}

	return 0
}

func (p *Paginator) getLastItemOnPage() int {
	return p.offset + p.length
}

func (p *Paginator) setBase(base int) *Paginator {
	p.base = base
	return p
}

func (p *Paginator) getBase() int {
	return p.base
}

func (p *Paginator) getPageIndex() int {
	index := math.Max(0, float64(p.page-p.base))

	if p.getItemCount() == nil {
		return int(index)
	}

	return int(math.Min(index, math.Max(0, float64(*p.getPageCount()-1))))
}

func (p *Paginator) isFirst() bool {
	return p.getPageIndex() == 0
}

func (p *Paginator) isLast() bool {
	if p.getItemCount() == nil {
		return false
	}

	return p.getPageIndex() >= *p.getPageCount()-1
}

func (p *Paginator) getPageCount() *int {
	if p.getItemCount() == nil {
		return nil
	}

	count := getLastPage(p)

	return &count
}

func (p *Paginator) setItemsPerPage(itemsPerPage int) *Paginator {
	p.itemsPerPage = max(1, itemsPerPage)
	return p
}

func (p *Paginator) getItemsPerPage() int {
	return p.itemsPerPage
}

func (p *Paginator) setItemCount(itemCount int) *Paginator {
	if itemCount < 0 {
		itemCount = 0
	}
	p.itemCount = &itemCount

	return p
}

func (p *Paginator) getItemCount() *int {
	return p.itemCount
}

func (p *Paginator) getOffset() int {
	return p.getPageIndex() * p.getItemsPerPage()
}

func (p *Paginator) getCountdownOffset() *int {
	if p.getItemCount() == nil {
		return nil
	}
	val := max(0, *p.getItemCount()-(p.getPageIndex()-1)*p.getItemsPerPage())
	return &val
}
func (p *Paginator) getLength() int {

	if p.getItemCount() == nil {
		return p.getItemsPerPage()
	}

	val := min(p.getItemsPerPage(), *p.getItemCount()-p.getPageIndex()*p.getItemsPerPage())

	return val
}

func getLastPage(paginator *Paginator) int {
	return int(math.Ceil(float64(*paginator.getItemCount()) / float64(paginator.itemsPerPage)))
}
