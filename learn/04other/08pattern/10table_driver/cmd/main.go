package main

import "10table_driver/service"

func main() {
	handle()
}

func handle() {

	var i = 1
	ti1 := service.Tmap[i]
	ti1.SetName("t1 name 1")
	ti1.Pname()

	ti2 := service.Tmap[i]
	//ti2.SetName("t1 name 2")
	ti2.Pname()

}
