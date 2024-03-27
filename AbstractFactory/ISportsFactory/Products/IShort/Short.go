package IShort

import "github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Constant"

type IShort interface {
	SetLogo(Logo Constant.Brand)
	SetSize(Size string)
	SetGender(Gender string)
	SetModel(Model string)
	GetGender() string
	GetModel() string
	GetSize() string
}

type Short struct{}

func (Short) SetLogo(Logo Constant.Brand) {

}

func (Short) SetSize(Size string) {
	//TODO implement me
	panic("implement me")
}

func (Short) SetGender(Gender string) {
	//TODO implement me
	panic("implement me")
}

func (Short) SetModel(Model string) {
	//TODO implement me
	panic("implement me")
}

func (Short) GetGender() string {
	//TODO implement me
	panic("implement me")
}

func (Short) GetModel() string {
	//TODO implement me
	panic("implement me")
}

func (Short) GetSize() string {
	//TODO implement me
	panic("implement me")
}
