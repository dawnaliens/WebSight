package flag

import (
	"flag"
	"fmt"
)

// Define all the flags
// Configuration stores command line arguments
type Configuration struct {
	Help      bool
	Target    string
	StartPort int
	EndPort   int
}

// ParseFlags parses command line arguments
func ParseFlags() Configuration {
	config := Configuration{}

	flag.BoolVar(&config.Help, "h", false, "Displays help information")
	flag.StringVar(&config.Target, "target", "example.com", "Target domain to scan")
	flag.IntVar(&config.StartPort, "start", 80, "Starting port for scanning")
	flag.IntVar(&config.EndPort, "end", 100, "Ending port for scanning")

	flag.Parse()

	return config
}

func DisplayHelp() {
	fmt.Println("Usage of Subdomain and Port Scanner:")
	flag.PrintDefaults()
}
