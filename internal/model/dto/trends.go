package dto

import "time"

type TrendPathCond struct {
	Path string `uri:"path" binding:"required"`
}

type TrendCond struct {
	Title     string    `form:"title"`
	StartDate time.Time `form:"start_date" time_format:"2006-01-02 15:04:05"`
	EndDate   time.Time `form:"end_date" time_format:"2006-01-02 15:04:05"`
	PagerReq
}
