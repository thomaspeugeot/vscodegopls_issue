package main

import "fmt"

func main() {

	foo := (&Foo{}).Stage()
	foo.Bar = "Hi"

	for _, foo_ := range FooStage {
		fmt.Println(foo_.Bar)
	}
}

type Foo struct {
	Bar string
}

func (foo *Foo) Stage() *Foo {
	FooStage = append(FooStage, foo)
	return foo
}

var FooStage []*Foo
