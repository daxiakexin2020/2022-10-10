package main

func main() {
	Handle()
}

func Handle() {
	c := InitializeC("test 依赖反转")
	c.Show()
}
