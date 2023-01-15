package main

import "fmt"

func main() {
	digits := "234"
	dest := letterCombinations(digits)

	fmt.Println("dest", len(dest), dest)
}

/*
*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

// init : "234" 0, ""
func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		//a d g
		digit := string(digits[index]) //"2", "3" , "4"
		letters := phoneMap[digit]     //abc,def,gih
		lettersCount := len(letters)   //3
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i])) // abc, 1, combination=a
		}
	}
}
