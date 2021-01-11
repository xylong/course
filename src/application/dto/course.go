package dto

import "github.com/shopspring/decimal"

// 输入
type (
	SimpleCourseReq struct {
		ID int `uri:"id" binding:"required,min=1"`
	}
)

// 输出
type (
	SimpleCourseRep struct {
		ID          int `json:"id"`
		Name        string
		MarketPrice decimal.Decimal `json:"market_price"`
		SalePrice   decimal.Decimal `json:"sale_price"`
	}
)
