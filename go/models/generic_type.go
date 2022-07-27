package models

import "log"

type Gongstruct interface {
	Hello | Country
	GetValueName() string
}

func PrintNameOfValue[Type Gongstruct](instance *Type) {
	log.Println("Instance ", (*instance).GetValueName())
}
