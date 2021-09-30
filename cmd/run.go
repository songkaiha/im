package main

import "time"

func main() {
	//时间初始化
	timeLocal := time.FixedZone("CST", 3600*8)
	time.Local = timeLocal

	// Print the current configuration
	models.Conf.Print()
	// Starting Http Server
	if err := routers.App.Run(); err != nil {
		panic(err.Error())
	}
}
