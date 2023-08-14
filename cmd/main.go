package main

import "WebSight/internal/flag"

func main() {
	config := flag.ParseFlags()
	if config.Help {
		flag.DisplayHelp()
		return
	}
}
