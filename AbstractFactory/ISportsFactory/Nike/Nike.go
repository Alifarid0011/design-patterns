package Nike

import (
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Entities"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Products/IShirt"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Products/IShoe"
)

type Nike struct {
}

func (n *Nike) MakeShoe(Entry Entities.Shoe) IShoe.IShoe {
	return &Entry
}
func (n *Nike) MakeShirt(Entry Entities.Shirt) IShirt.IShirt {
	return &Entry
}
