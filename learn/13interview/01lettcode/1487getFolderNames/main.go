package main

import (
	"fmt"
	"strconv"
)

func main() {
	/**
	给你一个长度为 n 的字符串数组 names 。你将会在文件系统中创建 n 个文件夹：在第 i 分钟，新建名为 names[i] 的文件夹。

	由于两个文件 不能 共享相同的文件名，因此如果新建文件夹使用的文件名已经被占用，系统会以 (k) 的形式为新文件夹的文件名添加后缀，其中 k 是能保证文件名唯一的 最小正整数 。

	返回长度为 n 的字符串数组，其中 ans[i] 是创建第 i 个文件夹时系统分配给该文件夹的实际名称。

	来源：力扣（LeetCode）
	链接：https://leetcode.cn/problems/making-file-names-unique
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
	*/

	names := []string{"pes", "pes", "pes", "pes(2019)", "pes(2019)", "pes(2019)(4)", "pes(2019)"}
	//names := []string{"gta", "gta(1)", "gta", "avalon"}
	//names := []string{"kaido", "kaido(1)", "kaido", "kaido(1)"}
	folderNames := getFolderNames(names)
	fmt.Println("res:", folderNames)
}

func getFolderNames(names []string) []string {
	ans := make([]string, len(names))
	index := map[string]int{}
	for p, name := range names {
		i := index[name]
		//不在map中，直接保存跳过
		if i == 0 {
			index[name] = 1
			ans[p] = name
			continue
		}
		//从map中寻找最小值，带括号的
		for index[name+"("+strconv.Itoa(i)+")"] > 0 {
			i++
		}
		newFile := name + "(" + strconv.Itoa(i) + ")"
		index[name] = i + 1
		index[newFile] = 1
		ans[p] = newFile
	}
	return ans
}
