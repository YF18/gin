package main

import "./router"

func main() {
	router := router.Init()
	router.Run(":88")
}
