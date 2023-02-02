package main

import "fmt"

func main() {
	students := []int{1, 1, 1, 0, 0, 1}
	sandwiches := []int{1, 0, 0, 0, 1, 1}
	res := countStudents(students, sandwiches)
	fmt.Printf("res=%d", res)
}

func countStudents(students []int, sandwiches []int) int {
	studentsTypeCount := make(map[int]int)
	for _, v := range students {
		studentsTypeCount[v]++
	}

	var stackIndex int
	for i := 0; i < len(students); {
		dest := students[i]
		if dest == sandwiches[stackIndex] {
			students = students[i+1:]
			studentsTypeCount[dest]--
			stackIndex++
		} else {
			if studentsTypeCount[0] == len(students) || studentsTypeCount[1] == len(students) {
				break
			}
			students = append(students[i+1:], dest)
		}
	}
	return len(students)
}
