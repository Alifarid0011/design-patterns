package IShirt

import "github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Constant"

type IShirt interface {
	SetLogo(Logo Constant.Brand)
	SetSize(size string)
	GetLogo() Constant.Brand
	GetSize() string
}
type Shirt struct {
	Logo Constant.Brand
	Size string
}

func (s Shirt) SetLogo(Logo Constant.Brand) {
	s.Logo = Logo
}

func (s Shirt) SetSize(Size string) {
	s.Size = Size
}

func (s Shirt) GetLogo() Constant.Brand {
	return s.Logo
}

func (s Shirt) GetSize() string {
	return s.Size
}
