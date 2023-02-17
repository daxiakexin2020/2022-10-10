package main

import "fmt"

/*
*
你有一个只支持单个标签页的 浏览器，最开始你浏览的网页是homepage，你可以访问其他的网站url，也可以在浏览历史中后退steps步或前进steps步。

请你实现BrowserHistory 类：

BrowserHistory(string homepage)，用homepage初始化浏览器类。
void visit(string url)从当前页跳转访问 url 对应的页面。执行此操作会把浏览历史前进的记录全部删除。
string back(int steps)在浏览历史中后退steps步。如果你只能在浏览历史中后退至多x步且steps > x，那么你只后退x步。请返回后退 至多 steps步以后的url。
string forward(int steps)在浏览历史中前进steps步。如果你只能在浏览历史中前进至多x步且steps > x，那么你只前进 x步。请返回前进至多steps步以后的 url。
*/
func main() {
	bh := Constructor("h")
	bh.Visit("1")
	bh.Visit("2")
	bh.Visit("3")

	u := bh.Back(3)
	fmt.Println("u", u)

	u1 := bh.Forward(31)
	fmt.Println("u", u1)
}

type BrowserHistory struct {
	list  []string
	index int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		list: []string{"homepage"},
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.list = append(this.list, url)
	this.index = len(this.list) - 1
}

func (this *BrowserHistory) Back(steps int) string {
	if this.index < steps {
		return this.list[0]
	}
	bindex := this.index - steps
	this.index = bindex
	return this.list[bindex]
}

func (this *BrowserHistory) Forward(steps int) string {
	if this.index+steps > len(this.list)-1 {
		return this.list[len(this.list)-1]
	}
	bindex := this.index + steps
	this.index = bindex
	return this.list[bindex]
}
