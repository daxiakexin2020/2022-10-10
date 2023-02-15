package main

import "fmt"

/*
*
一个 ATM 机器，存有 5 种面值的钞票：20 ，50 ，100 ，200 和 500 美元。初始时，ATM 机是空的。用户可以用它存或者取任意数目的钱。

取款时，机器会优先取 较大 数额的钱。

比方说，你想取 $300 ，并且机器里有 2 张 $50 的钞票，1 张 $100 的钞票和1 张 $200 的钞票，那么机器会取出 $100 和 $200 的钞票。
但是，如果你想取 $600 ，机器里有 3 张 $200 的钞票和1 张 $500 的钞票，那么取款请求会被拒绝，因为机器会先取出 $500 的钞票，然后无法取出剩余的 $100 。注意，因为有 $500 钞票的存在，机器 不能 取 $200 的钞票。
请你实现 ATM 类：

ATM() 初始化 ATM 对象。
void deposit(int[] banknotesCount) 分别存入 $20 ，$50，$100，$200 和 $500 钞票的数目。
int[] withdraw(int amount) 返回一个长度为 5 的数组，分别表示 $20 ，$50，$100 ，$200 和 $500 钞票的数目，并且更新 ATM 机里取款后钞票的剩余数量。如果无法取出指定数额的钱，请返回 [-1] （这种情况下 不 取出任何钞票）。
*/

func main() {
	atm := Constructor()
	data := []int{1, 2, 3, 4, 5}
	atm.Deposit(data)

	res := atm.Withdraw(1000)
	fmt.Println("res", res)
}

type ATM struct {
	store []int
}

var dict = []int{20, 50, 100, 200, 500}

func Constructor() ATM {
	var store = make([]int, 5)
	return ATM{store}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i, v := range banknotesCount {
		this.store[i] += v
	}
}

func (this *ATM) Withdraw(amount int) []int {
	ret := []int{0, 0, 0, 0, 0}
	for i := 4; i >= 0; i-- {
		//amout/dictp[i]需要几张    this.store[i] 真实剩余几张
		//curmax 当前需要使用的数量
		curmax := min(int(amount/dict[i]), this.store[i])
		if amount < curmax*dict[i] {
			continue
		}
		amount -= curmax * dict[i]
		ret[i] = curmax
	}
	if 0 < amount {
		return []int{-1}
	}
	for i := 0; i < 5; i++ {
		this.store[i] -= ret[i]
	}
	return ret
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
