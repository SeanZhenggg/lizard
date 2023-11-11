package po

const (
	pagerDefaultIndex = 1
	pagerDefaultSize  = 10
)

type Pager struct {
	Index int
	Size  int
}

func (p *Pager) GetOffset() int64 {
	return int64(p.Size * (p.Index - 1))
}

func (p *Pager) GetSize() int64 {
	if p.Size < 1 {
		return pagerDefaultSize
	}

	return int64(p.Size)
}

func (p *Pager) GetIndex() int64 {
	if p.Index < 1 {
		return pagerDefaultIndex
	}

	return int64(p.Index)
}

func NewPagerResult(paging *Pager, total int64) *PagerResult {
	totalPage := total / paging.GetSize()
	if total%paging.GetSize() > 0 {
		totalPage++
	}

	return &PagerResult{
		Index: int(paging.GetIndex()),
		Size:  int(paging.GetSize()),
		Pages: int(totalPage),
		Total: int(total),
	}
}

type PagerResult struct {
	Index int
	Size  int
	Pages int
	Total int
}
