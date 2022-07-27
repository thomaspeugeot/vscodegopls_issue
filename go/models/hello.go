package models

type Hello struct {
	Name string
}

func (hello *Hello) GetName() string {
	return hello.Name
}
