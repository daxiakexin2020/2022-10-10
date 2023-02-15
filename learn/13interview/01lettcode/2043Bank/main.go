package main

import (
	"fmt"
	"sync"
)

/**
你的任务是为一个很受欢迎的银行设计一款程序，以自动化执行所有传入的交易（转账，存款和取款）。
银行共有 n 个账户，编号从 1 到 n 。每个账号的初始余额存储在一个下标从 0 开始的整数数组 balance 中，其中第 (i + 1) 个账户的初始余额是 balance[i] 。

请你执行所有 有效的 交易。如果满足下面全部条件，则交易 有效 ：

指定的账户数量在 1 和 n 之间，且取款或者转账需要的钱的总数 小于或者等于 账户余额。
实现 Bank 类：

Bank(long[] balance) 使用下标从 0 开始的整数数组 balance 初始化该对象。
boolean transfer(int account1, int account2, long money) 从编号为 account1 的账户向编号为 account2 的账户转帐 money 美元。如果交易成功，返回 true ，否则，返回 false 。
boolean deposit(int account, long money) 向编号为 account 的账户存款 money 美元。如果交易成功，返回 true ；否则，返回 false 。
boolean withdraw(int account, long money) 从编号为 account 的账户取款 money 美元。如果交易成功，返回 true ；否则，返回 false 。
*/

func main() {
	b := []int64{
		100, 200,
	}
	back := Constructor(b)
	res := back.Transfer(1, 2, 2)
	fmt.Printf("res=%t\n", res)
	fmt.Println("back", back.balance[0], back.balance[1])
}

type Bank struct {
	balance []*Entry
}
type Entry struct {
	mu    sync.Mutex
	money int64
}

func Constructor(balance []int64) Bank {
	var b []*Entry
	for _, v := range balance {
		e := &Entry{
			money: v,
		}
		b = append(b, e)
	}
	return Bank{balance: b}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 > len(this.balance) || account2 > len(this.balance) {
		return false
	}
	a1 := this.balance[account1-1]
	a2 := this.balance[account2-1]
	a1.mu.Lock()
	a2.mu.Lock()
	defer a1.mu.Unlock()
	defer a2.mu.Unlock()

	if a1.money < money {
		return false
	}
	a1.money -= money
	a2.money += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if money < 0 {
		return false
	}
	if account > len(this.balance) {
		return false
	}
	a := this.balance[account-1]
	a.mu.Lock()
	defer a.mu.Unlock()
	a.money += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if money < 0 {
		return false
	}
	if account > len(this.balance) {
		return false
	}
	a := this.balance[account-1]
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.money < money {
		return false
	}
	a.money -= money
	return true
}
