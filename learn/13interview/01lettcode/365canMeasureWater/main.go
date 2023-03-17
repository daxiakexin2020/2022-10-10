package main

func main() {
	/**
	有两个水壶，容量分别为 jug1Capacity 和 jug2Capacity 升。水的供应是无限的。确定是否有可能使用这两个壶准确得到 targetCapacity 升。
	如果可以得到 targetCapacity 升水，最后请用以上水壶中的一或两个来盛放取得的 targetCapacity 升水。

	你可以：
	装满任意一个水壶
	清空任意一个水壶
	从一个水壶向另外一个水壶倒水，直到装满或者倒空

	输入: jug1Capacity = 3, jug2Capacity = 5, targetCapacity = 4
	输出: true

	输入: jug1Capacity = 2, jug2Capacity = 6, targetCapacity = 5
	输出: false

	*/
}

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if jug1Capacity == targetCapacity || jug2Capacity == targetCapacity {
		return true
	}
	return true
}
