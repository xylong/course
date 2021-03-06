package model

import (
	"course/src/domain/model/price"
	"course/src/infrastructure/GormDao"
	"course/src/test/internal/lib"
	"testing"
)

func addCoursePrice(courseID int, marketPrice float64, salePrice float64) error {
	cp := price.New(
		price.SetCourseID(courseID),
		price.SetMarketPrice(marketPrice),
		price.SetSalePrice(salePrice),
	)
	return lib.DB.Create(cp).Error
}

// TestAddCoursePrice 测试添加课程价格
func TestAddCoursePrice(t *testing.T) {
	data := []struct {
		id          int
		marketPrice float64
		salePrice   float64
	}{
		{1, 3.1415926, 2.5},
		{2, 0, 20},
		{3, 10, 0},
		{4, 108, 250},
	}

	for i := range data {
		id, mp, sp := data[i].id, data[i].marketPrice, data[i].salePrice
		if err := addCoursePrice(id, mp, sp); err != nil {
			t.Errorf("create %d error: %s\n", id, err.Error())
		}
	}
}

// TestCoursePriceLoad 测试课程价格查询
func TestCoursePriceLoad(t *testing.T) {
	repo := GormDao.NewCoursePriceRepo(lib.DB)
	price := price.New(
		price.SetCourseID(1),
		price.SetRepo(repo),
	)
	if err := price.Load(); err != nil {
		t.Error(err)
	} else {
		t.Logf("id=%d market_price=%s sale_price=%s", price.ID, price.Price.MarketPrice, price.Price.SalePrice)
	}
}
