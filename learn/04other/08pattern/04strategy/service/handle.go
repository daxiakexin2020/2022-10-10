package service

type Strategyer interface {
	handle(money float64) float64
}

type S1 struct{}

type S2 struct{}

type Ow struct {
	h Strategyer
}

func (s1 *S1) handle(money float64) float64 {
	return money * 0.9
}

func (s2 *S2) handle(money float64) float64 {
	if money >= 200 {
		return money - 10
	}
	return money
}

/*
TODO 一般策略模式可以结合工厂模式一起使用，此处的Strategy可以由上游决定使用哪个策略，

	也可以结合简单工厂，由工厂吐出具体的策略类，简单工厂模式与策略模式的区别在于，是否给上游吐出具体的业务类
	简单工厂模式是会向上游吐出具体的业务类，策略模式不会吐出，而是会吐出一个类似的代理或者门面的一个类
*/
func NewOW(strategyer Strategyer) *Ow {
	return &Ow{h: strategyer}
}

func (ow *Ow) Handle(money float64) float64 {
	//TODO 此处也可以做一些公共的通用的业务处理，比如类似代理的作用...
	return ow.h.handle(money)
}
