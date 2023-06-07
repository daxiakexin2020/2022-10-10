package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 1}
	number := singleNumber(nums)
	fmt.Println("number:", number)
}

func singleNumber(nums []int) int {
	flag := make(map[int]int)
	for _, num := range nums {
		flag[num]++
	}
	for num, count := range flag {
		if count == 1 {
			return num
		}
	}
	return 0
}

/**

收款信息模块
	1。收款职员、职员账号有无联动关系，使用登陆者的身份证号分别查询基础数据，得到2个列表，分别选择即可？
			职员  todo 待定
	2。收款客商、客商账号有无联动关系，使用登陆者的身份证号分别查询基础数据，得到2个列表，分别选择即可？
			下拉选择
   	3。收款户名，需要由职员账号或者是客商账号带出，具体是使用职员账号或者是客商账号的哪个字段？ name？
       批量支付  户名  账号
		一次性：手动
		其他不传，传空都可以
*/
