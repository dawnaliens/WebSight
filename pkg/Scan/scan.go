package Scan

import (
	"WebSight/internal/flag"
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

// commonSubdomains represents a list of frequently used subdomains
var commonSubdomains []string

func ReadSubdomains(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commonSubdomains = append(commonSubdomains, scanner.Text())
	}
	return commonSubdomains
}

func ScanPort(host string, port int, results chan<- int) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err == nil {
		conn.Close()
		results <- port
	}
}

func FindSubdomains(domain string, found chan<- string) {
	ReadSubdomains("subdomains-top1mil.txt")
	for _, sub := range commonSubdomains {
		d := fmt.Sprintf("%s.%s", sub, domain)
		ips, err := net.LookupIP(d)
		if err == nil && len(ips) > 0 {
			found <- d
		}
	}
	close(found)
}

func GetHTTPStatus(subdomain string, statuses chan<- struct {
	URL    string
	Status int
}) {
	url := fmt.Sprintf("http://%s", subdomain)
	resp, err := http.Get(url)
	if err == nil {
		statuses <- struct {
			URL    string
			Status int
		}{URL: url, Status: resp.StatusCode}
	}
}

func ExecuteScanner(config flag.Configuration) {
	target := config.Target

	var wg sync.WaitGroup

	// Concurrent port scanning
	results := make(chan int)
	for i := config.StartPort; i <= config.EndPort; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			ScanPort(target, port, results)
		}(i)
	}

	// Collect and print open ports
	go func() {
		wg.Wait()
		close(results)
	}()

	for port := range results {
		fmt.Printf("Port %d is open\n", port)
	}

	// Concurrent subdomain collection
	subdomainChannel := make(chan string)
	go FindSubdomains(target, subdomainChannel)

	var subdomains []string
	for subdomain := range subdomainChannel {
		subdomains = append(subdomains, subdomain)
		fmt.Printf("Found subdomain: %s\n", subdomain)
	}

	// Concurrent HTTP status checks
	statuses := make(chan struct {
		URL    string
		Status int
	}, len(subdomains))

	for _, subdomain := range subdomains {
		wg.Add(1)
		go func(sd string) {
			defer wg.Done()
			GetHTTPStatus(sd, statuses)
		}(subdomain)
	}

	// Collect and print HTTP statuses
	go func() {
		wg.Wait()
		close(statuses)
	}()

	for status := range statuses {
		fmt.Printf("Status for %s: %d\n", status.URL, status.Status)
	}
}
