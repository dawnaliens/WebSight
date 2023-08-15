package main

import (
	"WebSight/config"
	"WebSight/internal/flag"
)

func main() {
	config.Init()
	config := flag.ParseFlags()
	if config.Help {
		flag.DisplayHelp()
		return
	}
}
