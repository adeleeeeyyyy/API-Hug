package reporter

import (
	"api-stress-tester/tester"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/template"
)

func PrintConsole(comp tester.Comparison) {
	fmt.Println("\n=========================================================")
	fmt.Println("                   TEST RESULTS COMPARISON               ")
	fmt.Println("=========================================================")
	fmt.Printf("%-20s | %-15s | %-15s\n", "Metric", "Local", "Server")
	fmt.Println("---------------------------------------------------------")

	fmt.Printf("%-20s | %-15d | %-15d\n", "Total Requests", comp.LocalStats.TotalRequests, comp.ServerStats.TotalRequests)
	fmt.Printf("%-20s | %-15d | %-15d\n", "Successful Requests", comp.LocalStats.SuccessCount, comp.ServerStats.SuccessCount)
	fmt.Printf("%-20s | %-15d | %-15d\n", "Failed Requests", comp.LocalStats.ErrorCount, comp.ServerStats.ErrorCount)
	fmt.Printf("%-20s | %-15s | %-15s\n", "Min Response Time", comp.LocalStats.MinTime, comp.ServerStats.MinTime)
	fmt.Printf("%-20s | %-15s | %-15s\n", "Max Response Time", comp.LocalStats.MaxTime, comp.ServerStats.MaxTime)
	fmt.Printf("%-20s | %-15s | %-15s\n", "Avg Response Time", comp.LocalStats.AvgTime, comp.ServerStats.AvgTime)
	fmt.Println("=========================================================")
	fmt.Printf("Conclusion: %s environment is FASTER by %.2f ms on average.\n", comp.FasterEnv, comp.DifferenceMs)
}

func ExportJSON(comp tester.Comparison, filename string) error {
	data, err := json.MarshalIndent(comp, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err == nil {
		log.Printf("Successfully exported JSON report to %s", filename)
	}

	return err
}

func ExportHTML(comp tester.Comparison, filename string) error {
	t := template.Must(template.New("report").Parse(htmlTemplateString))
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, comp)
	if err == nil {
		log.Printf("Successfully exported html report to %s", filename)
	}

	return err
}