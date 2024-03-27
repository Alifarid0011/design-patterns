package ISportsFactory

import (
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Adidas"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Constant"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Entities"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Nike"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Products/IShirt"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Products/IShoe"
)

type ISportFactory interface {
	MakeShoe(Entry Entities.Shoe) IShoe.IShoe
	MakeShirt(Entry Entities.Shirt) IShirt.IShirt
}

var (
	BrandMapper = map[Constant.Brand]ISportFactory{
		Constant.Adidas: &Adidas.Adidas{},
		Constant.Nike:   &Nike.Nike{},
	}
)

func GetSportsFactory(brand Constant.Brand) ISportFactory {
	value, exist := BrandMapper[brand]
	if exist {
		return value
	}
	return BrandMapper[Constant.Adidas]
}
