package flag

import (
	"flag"
	"fmt"
	"os"
	"strings"
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
	flag.StringVar(&config.Output, "output", "output.txt", "Output file name")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of Subdomain and Port Scanner:\n")
		flag.VisitAll(func(f *flag.Flag) {
			// Split by comma and remove data type from the usage string
			parts := strings.Split(f.Usage, ",")
			fmt.Fprintf(os.Stderr, "  -%s=%s\n    \t%s\n", f.Name, f.DefValue, parts[0])
		})
	}

	flag.Parse()

	return config
}

func DisplayHelp() {
	fmt.Println("Usage of Subdomain and Port Scanner:")
	flag.PrintDefaults()
}
