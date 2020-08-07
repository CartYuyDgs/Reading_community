package models

type Category struct {
	Id     int
	Pid    int
	Title  string `orm:"size(30);unique"`
	Intro  string
	Icon   string
	Cnt    int
	Sort   int
	Status bool //状态 true 显示
}

func (m *Category) TableName() string {
	return TNCategory()
}
