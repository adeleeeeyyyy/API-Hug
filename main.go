package main

import (
	"api-stress-tester/reporter"
	"api-stress-tester/tester"
	"flag"
	"fmt"
	"log"
)

func main() {
	var localUrl, serverUrl string
	var requestCount, concurrency int
	var exportHTML, exportJSON bool

	flag.StringVar(&localUrl, "local", "http://localhost:8080/api", "local environtment URL")
	flag.StringVar(&serverUrl, "server", "https://api.example.com/api", "Server environtment URL")
	flag.IntVar(&requestCount, "n", 100, "Number of request per environtment")
	flag.IntVar(&concurrency, "c", 10, "Concurrency level")
	flag.BoolVar(&exportHTML, "html", false, "Export result to html file")
	flag.BoolVar(&exportJSON, "json", false, "Export result to JSON file")
	flag.Parse()

	fmt.Printf("Starting the load test... [%d reqs | %d workers]\n", requestCount, concurrency)

	fmt.Print("Testing local...")
	localStats := tester.RunLoadTest("Local", localUrl, requestCount, concurrency)
	fmt.Println("Done!")

	fmt.Print("Testing Server")
	serverStats := tester.RunLoadTest("Server",serverUrl, requestCount, concurrency)
	fmt.Println("Done!")

	comp := tester.Compare(localStats, serverStats)

	reporter.PrintConsole(comp)

	if exportHTML {
		if err := reporter.ExportHTML(comp, "report.html"); err != nil {
			log.Printf("Failed to export HTML: %v\n", err)
		}
	}

	if exportJSON {
		if err := reporter.ExportJSON(comp, "report.json"); err != nil {
			log.Printf("Failed to export JSON: %v\n", err)
		}
	}
}