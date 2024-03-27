package Adidas

import (
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Entities"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Products/IShirt"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Products/IShoe"
)

type Adidas struct {
}

func (a *Adidas) MakeShoe(Entry Entities.Shoe) IShoe.IShoe {
	return &Entry
}
func (a *Adidas) MakeShirt(Entry Entities.Shirt) IShirt.IShirt {
	return &Entry
}
