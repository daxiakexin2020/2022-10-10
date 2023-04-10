package main

func main() {
	/**
	题目中第i 天适合打劫需满足：第i 天前连续time天警卫数目都是非递增与第i 天后连续time 天警卫数目都是非递减。只需要预先计算出第
	i 天前警卫数目连续非递增的天数以及第i 天后警卫数目连续非递减的天数即可判断第i 天是否适合打劫。设第i 天前警卫数目连续非递增的天数为
	lefti，第i 天后警卫数目连续非递减的天数为righti，当第i 天同时满足
	left≥time,right≥time lefti≥time,righti≥time 时，即可认定第
	i 天适合打劫。计算连续非递增和非递减的天数的方法如下：
	如果第i 天的警卫数目小于等于第−1天的警卫数目，假设已知第i−1 天前有
	j 天连续非递增，则此时满足
	security−1≤
	security−2⋯≤security−1
	securityi−1≤securityi−2 ⋯≤securityi−j−1，已知security≤security−1
	security
	i
	​
	 ≤security
	i−1
	​
	 ，可推出
	security
	�
	≤
	security
	�
	−
	1
	⋯
	≤
	security
	�
	−
	�
	−
	1
	security
	i
	​
	 ≤security
	i−1
	​
	 ⋯≤security
	i−j−1
	​
	 ，则此时
	left
	�
	=
	�
	+
	1
	=
	left
	�
	−
	1
	+
	1
	left
	i
	​
	 =j+1=left
	i−1
	​
	 +1；如果第
	�
	i 天的警卫数目大于第
	�
	−
	1
	i−1 天的警卫数目，则此时
	left
	�
	=
	0
	left
	i
	​
	 =0。

	如果第
	�
	i 天的警卫数目小于等于第
	�
	+
	1
	i+1 天的警卫数目，假设已知第
	�
	+
	1
	i+1 天后有
	�
	j 天连续非递减，则此时满足
	security
	�
	+
	1
	≤
	security
	�
	+
	2
	⋯
	≤
	security
	�
	+
	�
	+
	1
	security
	i+1
	​
	 ≤security
	i+2
	​
	 ⋯≤security
	i+j+1
	​
	 ，已知
	security
	�
	≤
	security
	�
	+
	1
	security
	i
	​
	 ≤security
	i+1
	​
	 ，可推出
	security
	�
	≤
	security
	�
	+
	1
	⋯
	≤
	security
	�
	+
	�
	+
	1
	security
	i
	​
	 ≤security
	i+1
	​
	 ⋯≤security
	i+j+1
	​
	 ，则此时
	right
	�
	=
	�
	+
	1
	=
	right
	�
	+
	1
	+
	1
	right
	i
	​
	 =j+1=right
	i+1
	​
	 +1；如果第
	�
	i 天的警卫数目大于第
	�
	+
	1
	i+1 天的警卫数目，则此时
	right
	�
	=
	0
	right
	i
	​
	 =0。
	*/
}

func goodDaysToRobBank(security []int, time int) (ans []int) {
	n := len(security)
	left := make([]int, n)
	right := make([]int, n)
	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			left[i] = left[i-1] + 1
		}
		if security[n-i-1] <= security[n-i] {
			right[n-i-1] = right[n-i] + 1
		}
	}

	for i := time; i < n-time; i++ {
		if left[i] >= time && right[i] >= time {
			ans = append(ans, i)
		}
	}
	return
}
