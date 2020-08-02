package main

import (
	"fmt"

	"github.com/go-series/struct/accessModifier"
	nonPtrReceive "github.com/go-series/struct/behavior/nonPtrReceiver"
	"github.com/go-series/struct/behavior/ptrReceiver"
)

func main() {
	fmt.Println("...Struct Initialization...")
	StructInitialize()
	fmt.Println()

	fmt.Println("..Access modifier..")
	AccessModifier()
	fmt.Println()

	fmt.Println("...Behavior of Object...")
	Behavior()
	fmt.Println()

}

//========================================================================
//======================  Struct Initialization ==========================

type Cat struct{}

func (c Cat) Run() {
	fmt.Println("cat is runing")
}

func StructInitialize() {

	var cat Cat
	cat1 := Cat{}
	cat2 := new(Cat)
	cat.Run()
	cat1.Run()
	cat2.Run()
}

//========================================================================
//=======================  Access Modifier ===============================

func AccessModifier() {
	public()
	private()
}

func public() {
	box := accessModifier.Box{
		Width: 5,
		Hight: 5,
	}
	fmt.Println(box)
}

func private() {
	//พอลองเอา comment ออกตัว linter จะบอกทัันทีว่าไม่พบ x y ใน Box
	//เพราะไม่สามารถ access ได้นั้นเอง

	// box := accessModifier.Box{
	// 	x: 5,
	// 	y: 5,
	// }
	// fmt.Println(box)
}

//========================================================================
//=====================  Behavior of Object  =============================

func NonPtrReceiver() {
	cat := nonPtrReceive.Cat{
		Name: "nano",
	}
	cat.Run()
	cat.Run()
}

func PtrReceiver() {
	cat := ptrReceiver.Cat{
		Name: "nano",
	}
	cat.Run()
	cat.Run()
}

func Behavior() {
	fmt.Println()
	fmt.Println("non pointer receiver...")
	NonPtrReceiver()
	fmt.Println()
	fmt.Println("pointer receiver...")
	PtrReceiver()
}
