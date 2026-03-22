package tester

import (
	"io"
	"net/http"
	"sync"
	"time"
)

func RunLoadTest(envName, url string, totalRequests, concurrencyLevel int) Stats {
	var wg sync.WaitGroup
	requests := make(chan int, totalRequests)
	results := make(chan Result, totalRequests)

	for i := 0; i < concurrencyLevel; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := &http.Client{Timeout: 10 * time.Second}
			for range requests {
				start := time.Now()
				req, err := http.NewRequest("GET", url, nil)

				var statusCode int
				if err == nil {
					resp, errReq := client.Do(req)
					if errReq == nil {
						statusCode = resp.StatusCode
						_, _ = io.Copy(io.Discard, resp.Body)
						resp.Body.Close()
					} else {
						err = errReq
					}
				}

				results <- Result{
					Duration: time.Since(start),
					StatusCode: statusCode,
					Error: err,
				}
			}
		}()
	}

	startTest := time.Now()

	go func ()  {
		for i := 0; i < totalRequests; i++ {
			requests <- 1
		}
		close(requests)
	}()

	go func ()  {
		wg.Wait()
		close(results)
	}()

	stats := Stats{
		Environtment: envName,
		URL: url,
		MinTime: time.Hour,
	}

	for res := range results {
		stats.TotalRequests++
		stats.TotalTime += res.Duration
		if res.Error != nil || res.StatusCode >= 400 {
			stats.ErrorCount++
		} else {
			stats.SuccessCount++
		}
		if res.Duration < stats.MinTime {stats.MinTime = res.Duration}
		if res.Duration > stats.MaxTime {stats.MaxTime = res.Duration}
	}

	if stats.MinTime == time.Hour { stats.MinTime = 0 }

	if stats.TotalRequests > 0 {
		stats.AvgTime = time.Duration(int64(stats.TotalTime) / int64(stats.TotalRequests))
	}

	stats.TotalTime = time.Since(startTest)
	return stats
}

func Compare(local, server Stats) Comparison {
	diff := float64(server.AvgTime - local.AvgTime) / float64(time.Millisecond)
	faster := "Local"
	if diff < 0 {
		faster = "Server"
		diff = -diff
	}

	return Comparison{
		LocalStats: local,
		ServerStats: server,
		DifferenceMs: diff,
		FasterEnv: faster,
	}
}