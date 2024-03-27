package Products

import (
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Constant"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Entities"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Products/IShirt"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Products/IShoe"
)

var ListOfShirts = []Entities.Shirt{
	{
		IShirt: IShirt.Shirt{
			Logo: Constant.Adidas,
			Size: "L",
		},
	},
	{
		IShirt: IShirt.Shirt{
			Logo: Constant.Nike,
			Size: "M",
		},
	},
}
var ListOfShoes = []Entities.Shoe{
	{
		IShoe: &IShoe.Shoe{
			Logo:     Constant.Adidas,
			Size:     42,
			Price:    "2,500,000 T",
			IsModern: true,
		},
	},
	{
		IShoe: &IShoe.Shoe{
			Logo:     Constant.Adidas,
			Size:     40,
			Price:    "4,000,000 T",
			IsModern: false,
		},
	},
	{
		IShoe: &IShoe.Shoe{
			Logo:     Constant.Nike,
			Size:     42,
			Price:    "7,000,000 T",
			IsModern: true,
		},
	},
}

func CreatShoe() (Shoe [3]Entities.Shoe) {
	for i, shoe := range ListOfShoes {
		Shoe[i] = shoe
	}
	return Shoe
}
func CreatShirt() (Shirt [2]Entities.Shirt) {
	for i, shirt := range ListOfShirts {
		Shirt[i] = shirt
	}
	return Shirt
}
