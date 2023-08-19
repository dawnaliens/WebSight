package flag

import (
	"flag"
	"fmt"
	"os"
)

// Configuration stores command line arguments
type Configuration struct {
	Help      bool
	Target    string
	StartPort int
	EndPort   int
	Output    string
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
	fmt.Println("Usage: go run main.go [options]\n")
	fmt.Println("Options:")

	const flagWidth = 10
	const valueWidth = 20
	flag.VisitAll(func(f *flag.Flag) {
		// Modify the display format here, avoiding data type
		fmt.Fprintf(os.Stderr, "  -%-"+fmt.Sprint(flagWidth)+"s %-"+fmt.Sprint(valueWidth)+"s %s\n", f.Name, f.DefValue, f.Usage)
	})
}
