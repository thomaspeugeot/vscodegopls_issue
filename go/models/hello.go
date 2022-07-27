package models

type Hello struct {
	Name string
}

func (hello *Hello) GetName() string {
	return hello.Name
}

func (hello Hello) GetValueName() string {
	return hello.Name
}

func (hello *Hello) GetValuePointer() string {
	return hello.Name
}
