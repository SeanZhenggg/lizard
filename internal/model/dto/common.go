package dto

type PagerReq struct {
	PageIndex int `form:"page_index,default=1"`
	PageSize  int `form:"page_size,default=10"`
}

type PagerResp struct {
	Index int `json:"pageIndex"` // 頁碼
	Size  int `json:"pageSize"`  // 比數
	Pages int `json:"pages"`     // 總頁數
	Total int `json:"total"`     // 總筆數
}

type ListRes struct {
	List  interface{}
	Pager PagerResp
}
