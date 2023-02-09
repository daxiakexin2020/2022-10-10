package main

import (
	"fmt"
	"sort"
)

/*
*
设计一个支持下述操作的食物评分系统：

修改 系统中列出的某种食物的评分。
返回系统中某一类烹饪方式下评分最高的食物。
实现 FoodRatings 类：

FoodRatings(String[] foods, String[] cuisines, int[] ratings) 初始化系统。食物由 foods、cuisines 和 ratings 描述，长度均为 n 。
foods[i] 是第 i 种食物的名字。
cuisines[i] 是第 i 种食物的烹饪方式。
ratings[i] 是第 i 种食物的最初评分。
void changeRating(String food, int newRating) 修改名字为 food 的食物的评分。
String highestRated(String cuisine) 返回指定烹饪方式 cuisine 下评分最高的食物的名字。如果存在并列，返回 字典序较小 的名字。
注意，字符串 x 的字典序比字符串 y 更小的前提是：x 在字典中出现的位置在 y 之前，也就是说，要么 x 是 y 的前缀，或者在满足 x[i] != y[i] 的第一个位置 i 处，x[i] 在字母表中出现的位置在 y[i] 之前。
*/
func main() {
	foods := []string{"a", "b", "c", "芒果", "橘子"}
	cuisines := []string{"A", "A", "A", "B", "C"}
	ratings := []int{2, 2, 2, 3, 4}
	c := Constructor(foods, cuisines, ratings)
	res := c.HighestRated("A")
	fmt.Println("res", res, c)
}

// 初始化、修改时需要维护有序，否则会超时
type FoodRatings struct {
	Foods map[string]int
	Types map[string][]string
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {

	fd := make(map[string]int, len(foods))
	types := make(map[string][]string)
	for k, food := range foods {
		fd[food] = ratings[k]
		t := cuisines[k]
		sub := append(types[t], food)
		types[t] = sub
	}
	return FoodRatings{
		Foods: fd,
		Types: types,
	}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	this.Foods[food] = newRating
}

func (this *FoodRatings) HighestRated(cuisine string) string {

	foods := this.Types[cuisine]
	var highScore int
	var f []string
	s := make(map[int][]string)
	var ks []int

	for _, food := range foods {
		if this.Foods[food] >= highScore {
			if this.Foods[food] > highScore {
				delete(s, highScore)
				if len(ks) >= 1 {
					ks = ks[:len(ks)-1]
				}
			}
			highScore = this.Foods[food]
			f = append(f, food)
			fs := s[highScore]
			fs = append(fs, food)
			s[highScore] = fs
			ks = append(ks, highScore)
		}
	}
	sort.Ints(ks)
	fff := s[ks[len(ks)-1]]
	sort.Strings(fff)
	return fff[0]
}
