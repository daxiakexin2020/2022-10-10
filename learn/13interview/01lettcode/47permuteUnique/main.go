package main

import "sort"

func main() {
	/**
	  思路和算法
	  此题是「46. 全排列」的进阶，序列中包含了重复的数字，要求我们返回不重复的全排列，那么我们依然可以选择使用搜索回溯的方法来做。
	  我们将这个问题看作有
	  n 个排列成一行的空格，我们需要从左往右依次填入题目给定的
	  n 个数，每个数只能使用一次。那么很直接的可以想到一种穷举的算法，即从左往右每一个位置都依此尝试填入一个数，看能不能填完这
	  n 个空格，在程序中我们可以用「回溯法」来模拟这个过程。
	  我们定义递归函数
	  backtrack(idx,perm) 表示当前排列为perm，下一个待填入的位置是第idx,idx 个位置（下标从0开始）。那么整个递归函数分为两个情况：
	  如果idx=n，说明我们已经填完了n 个位置，找到了一个可行的解，我们将perm放入答案数组中，递归结束。
	  如果idx<n，我们要考虑第idx 个位置填哪个数。根据题目要求我们肯定不能填已经填过的数，因此很容易想到的一个处理手段是我们定义一个标记数组
	  vis 来标记已经填过的数，那么在填第 idx 个数的时候我们遍历题目给定的n 个数，如果这个数没有被标记过，我们就尝试填入，并将其标记，继续尝试填下一个位置，即调用函数
	  backtrack(idx+1,perm)。搜索回溯的时候要撤销该个位置填的数以及标记，并继续尝试其他没被标记过的数。
	  但题目解到这里并没有满足「全排列不重复」 的要求，在上述的递归函数中我们会生成大量重复的排列，因为对于第
	  idx 的位置，如果存在重复的数字i，我们每次会将重复的数字都重新填上去并继续尝试导致最后答案的重复，因此我们需要处理这个情况。
	  要解决重复问题，我们只要设定一个规则，保证在填第
	  idx 个数的时候重复数字只会被填入一次即可。而在本题解中，我们选择对原数组排序，保证相同的数字都相邻，然后每次填入的数一定是这个数所在重复数集合中「从左往右第一个未被填过的数字」，即如下的判断条件：
	*/
}

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			//重复的数
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}
