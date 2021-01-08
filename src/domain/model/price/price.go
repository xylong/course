package price

import "github.com/shopspring/decimal"

// Price 价格
type Price struct {
	MarketPrice decimal.Decimal `json:"market_price" gorm:"type:decimal(10,2);comment:'市场价'"`
	SalePrice   decimal.Decimal `json:"sale_price" gorm:"type:decimal(10,2);comment:'销售价'"`
	Discount    decimal.Decimal `json:"discount" gorm:"-"` // 折扣
}

// getDiscount 获取折扣
func (p *Price) getDiscount() {
	zero, one := decimal.NewFromInt(0), decimal.NewFromInt(1)

	if p.MarketPrice.GreaterThan(zero) && p.SalePrice.GreaterThan(zero) {
		p.Discount = p.MarketPrice.Sub(p.SalePrice).DivRound(p.MarketPrice, 2).Sub(one).Abs()
	}
}

func NewPrice() *Price {
	return &Price{}
}

// SetMarketPrice 设置市场价
func SetMarketPrice(price float64) Attr {
	return func(coursePrice *CoursePrice) {
		coursePrice.Price.MarketPrice = decimal.NewFromFloat(price)
		coursePrice.Price.getDiscount()
	}
}

// SetSalePrice 设置销售价
func SetSalePrice(price float64) Attr {
	return func(coursePrice *CoursePrice) {
		coursePrice.Price.SalePrice = decimal.NewFromFloat(price)
		coursePrice.Price.getDiscount()
	}
}
