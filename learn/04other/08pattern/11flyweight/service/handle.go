package service

func Handle(name string, age int) *People {

	//初始化享元属性
	makeColorManage()

	//初始化具体
	p := NewPeople(name, age)

	return p
}
