package IShoe

import "github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Constant"

type IShoe interface {
	SetLogo(Logo Constant.Brand)
	SetSize(Size int)
	GetLogo() Constant.Brand
	GetSize() int
	SetPrice(Price string)
	GetPrice() string
	SetIsModern(iSModern bool)
	GetIsModern() bool
}
type Shoe struct {
	Logo     Constant.Brand
	Size     int
	Price    string
	IsModern bool
}

func (s *Shoe) SetLogo(Logo Constant.Brand) {
	s.Logo = Logo
}
func (s *Shoe) SetSize(Size int) {
	s.Size = Size
}
func (s *Shoe) GetLogo() Constant.Brand {

	return s.Logo
}
func (s *Shoe) GetSize() int {
	return s.Size
}
func (s *Shoe) SetPrice(Price string) {
	s.Price = Price
}
func (s *Shoe) SetIsModern(iSModern bool) {
	s.IsModern = iSModern
}
func (s *Shoe) GetIsModern() bool {
	return s.IsModern

}
func (s *Shoe) GetPrice() string {
	return s.Price
}
