package model

import (
	"course/src/domain/model/price"
	"course/src/test/lib"
	"testing"
)

func addCoursePrice(courseID int,marketPrice float64,salePrice float64) error {
	cp:=price.New(
		price.SetCourseID(courseID),
		price.SetMarketPrice(marketPrice),
		price.SetSalePrice(salePrice),
		)
	return lib.DB.Create(cp).Error
}

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

	for i:=range data{
		id,mp,sp:=data[i].id,data[i].marketPrice,data[i].salePrice
		if err:=addCoursePrice(id,mp,sp);err!=nil {
			t.Errorf("create %d error: %s\n", id, err.Error())
		}
	}
}
