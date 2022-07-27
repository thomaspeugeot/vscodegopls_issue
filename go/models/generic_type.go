package models

import "log"

type PointerToGongstruct interface {
	*Hello | *Country
	GetName() string
}

func PrintName[Type PointerToGongstruct](instance Type) {
	log.Println("Instance ", instance.GetName())
}

func Test() {
	hello := (&Hello{Name: "Bonjour"})

	log.Println("Hello world : ")
	PrintName(hello)
}
