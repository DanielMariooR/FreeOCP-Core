package pagination

import (
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Meta struct {
	Limit     int    `json:"item_per_page"`
	Page      int    `json:"page"`
	Count     uint64 `json:"total_item"`
	TotalPage int    `json:"total_page"`
}

type PaginationResponse struct {
	Meta  *Meta       `json:"meta"`
	Items interface{} `json:"items"`
}

func (f *Meta) FromContext(c echo.Context) *Meta {
	f.Page, _ = strconv.Atoi(c.QueryParam("page"))
	f.Limit, _ = strconv.Atoi(c.QueryParam("item_per_page"))
	return f
}

func (f *PaginationResponse) SetTotalPage() {
	f.Meta.TotalPage = int(math.Ceil(float64(f.Meta.Count) / float64(f.Meta.Limit)))
}
