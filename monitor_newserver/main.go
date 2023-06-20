package main

import "monitor_newserver/util"

func main() {
	server := new(util.BlockAnalysis)
	server.Run()
}
