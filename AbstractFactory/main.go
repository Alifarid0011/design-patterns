package main

import (
	"fmt"

	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Constant"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Products"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Products/IShirt"
	"github.com/Alifarid0011/design-patterns/AbstractFactory/ISportsFactory/Products/IShoe"
)

func FactoryFunc(brand Constant.Brand) ISportsFactory.ISportFactory {
	return ISportsFactory.GetSportsFactory(brand)
}

var Methods = map[Constant.EntityType]func(brand Constant.Brand) ISportsFactory.ISportFactory{
	Constant.Shoe:  FactoryFunc,
	Constant.Shirt: FactoryFunc,
	Constant.Short: FactoryFunc,
}

func main() {
	Shoes := Products.CreatShoe()
	Shirts := Products.CreatShirt()
	for _, shirt := range Shirts {
		fmt.Println("--------Shirt-------")
		factoryShirt := Methods[Constant.Shirt](shirt.GetLogo())
		printShirtDetails(factoryShirt.MakeShirt(shirt))
	}
	for _, Shoe := range Shoes {
		fmt.Println("--------Shoe-------")
		factoryShirt := Methods[Constant.Shirt](Shoe.GetLogo())
		printShoeDetails(factoryShirt.MakeShoe(Shoe))
	}
}

func printShoeDetails(s IShoe.IShoe) {
	fmt.Printf("Logo: %s \n", s.GetLogo())
	fmt.Printf("Size: %d \n", s.GetSize())
	fmt.Printf("Price: %s \n", s.GetPrice())
	fmt.Printf("IsModern : %v \n", s.GetIsModern())
}

func printShirtDetails(s IShirt.IShirt) {
	fmt.Printf("Logo: %s \n", s.GetLogo())
	fmt.Printf("Size: %s \n", s.GetSize())
}
