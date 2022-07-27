package models

import "log"

type Gongstruct interface {
	Hello | Country
	GetValueName() string
}

func PrintNameOfValue[Type Gongstruct](instance *Type) {
	log.Println("Instance ", (*instance).GetValueName())
}

type PointerToGongstruct interface {
	*Hello | *Country
	GetValuePointer() string
}

func PrintNameOfPointer[Type PointerToGongstruct](instance Type) {
	log.Println("Instance ", instance.GetValuePointer())
}

func Test() {
	hello := (&Hello{Name: "Bonjour"})

	log.Println("Hello world : ")
	PrintNameOfValue(hello)

	PrintNameOfPointer(hello)
}
