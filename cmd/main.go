package main

import (
	"WebSight/config"
	"WebSight/internal/flag"
	"WebSight/pkg/Scan"
	"fmt"
)

func main() {
	config.Init()
	config := flag.ParseFlags()
	if config.Help {
		flag.DisplayHelp()
		return
	}

	fmt.Println("Scanning target:", config)

	Scan.ExecuteScanner(config)

}
