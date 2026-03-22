# API Stress Tester & Comparator

A lightweight, Go-based command-line (CLI) tool for stress testing API endpoints and comparing performance between two environments (e.g., Local vs. Server/Production). 

> *We all know the local server is definitely faster, but I just want to make it more fun! and of course you would like to hug your API with 50000 requests per second* 🥰

## ✨ Features
- **Concurrent Load Testing:** Uses goroutines to send HTTP requests simultaneously.
- **Direct Comparison:** Automatically tests two URLs (Local & Server) and concludes which environment is faster.
- **Comprehensive Metrics:** Records Total Requests, Success/Failure Rates, Minimum, Maximum, and Average Response Times.
- **JSON Export:** Supports exporting test results to a `.json` file for automation or further parsing.
- **Styled HTML Export:** Generates a clean, reader-friendly `.html` dashboard report.

---

## 🚀 Prerequisites
Make sure you have Go installed on your system. 
If you haven't, download and install it from [golang.org](https://golang.org/).

---

## 🔧 Usage

You can run this tool directly using the `go run` command.

### Basic Command Structure
```bash
go run main.go [flags]
```

### Available Flags / Arguments
The tool accepts several configuration arguments:

| Flag       | Type   | Default                           | Description                                                  |
|------------|--------|-----------------------------------|--------------------------------------------------------------|
| `-local`   | string | `http://localhost:8080/api`       | URL of the API running in the local environment              |
| `-server`  | string | `https://api.example.com/api`     | URL of the API running in the server/production environment  |
| `-n`       | int    | `100`                             | Total number of requests to send per environment             |
| `-c`       | int    | `10`                              | Concurrency level (number of parallel requests running at once)|
| `-json`    | bool   | `false`                           | Add this flag to export the results to JSON                  |
| `-html`    | bool   | `false`                           | Add this flag to export the results to HTML                  |

---

## ⚡ Examples

**1. Simple Execution (Default)**
Sends 100 requests to localhost and 100 requests to example.com with a concurrency level of 10.
```bash
go run main.go
```

**2. Custom URL and Load Execution**
Send a total of 500 requests, with 50 goroutines running in parallel.
```bash
go run main.go -local="http://localhost:3000/api/users" -server="https://production.com/api/users" -n=500 -c=50
```

**3. Execution with HTML & JSON Report Generation**
Add `-html=true` and `-json=true` to automatically generate `report.html` and `report.json` output files in the same directory.
```bash
go run main.go -n=200 -c=20 -html=true -json=true
```

## 📂 Terminal Output
The terminal output will look like this:
```text
Starting the load test... [100 reqs | 10 workers]
Testing local...Done!
Testing Server...Done!

=========================================================
                   TEST RESULTS COMPARISON               
=========================================================
Metric               | Local           | Server         
---------------------------------------------------------
Total Requests       | 100             | 100            
Successful Requests  | 100             | 100            
Failed Requests      | 0               | 0              
Min Response Time    | 1.2ms           | 25.1ms         
Max Response Time    | 12.5ms          | 85.3ms         
Avg Response Time    | 3.1ms           | 32.5ms         
=========================================================
Conclusion: Local environment is FASTER by 29.40 ms on average.
```

If the `-html=true` flag is enabled, you can simply open the `report.html` file in any modern web browser to view its dashboard-style user interface.

---
*Built for Internal Load Testing & API Optimization.*
