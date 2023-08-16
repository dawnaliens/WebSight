package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

func scanPort(host string, port int, results chan<- int) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err == nil {
		conn.Close()
		results <- port
	}
}

var commonSubdomains = []string{"www", "mail", "ftp", "webmail"}

func findSubdomains(domain string, found chan<- string) {
	for _, sub := range commonSubdomains {
		d := fmt.Sprintf("%s.%s", sub, domain)
		ips, err := net.LookupIP(d)
		if err == nil && len(ips) > 0 {
			found <- d
		}
	}
	close(found)
}

func getHTTPStatus(subdomain string, statuses chan<- struct {
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

func main() {
	target := "auckland.ac.nz"

	var wg sync.WaitGroup

	// Concurrent port scanning
	results := make(chan int)
	for i := 80; i <= 100; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			scanPort(target, port, results)
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
	go findSubdomains(target, subdomainChannel)

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
			getHTTPStatus(sd, statuses)
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
