package reporter

// Template UI dari report stress test menggunakan HTML5 & CSS modern
const htmlTemplateString = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>API Stress Test Results</title>
    <style>
        :root { --primary: #34495e; --success: #1abc9c; --danger: #e74c3c; --bg: #f9f9f9; --card: #ffffff; }
        body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: var(--bg); color: #333; padding: 20px; text-align: center;}
        .container { max-width: 900px; margin: 0 auto; background: var(--card); padding: 30px; border-radius: 12px; box-shadow: 0 10px 15px -3px rgba(0,0,0,0.1); }
        h1 { color: var(--primary); margin-bottom: 30px; }
        .grid { display: grid; grid-template-columns: 1fr 1fr; gap: 20px; margin-bottom: 25px; }
        .card { background: #ecf0f1; padding: 20px; border-radius: 8px; text-align: left;}
        .card h3 { margin-top: 0; color: var(--primary); border-bottom: 2px solid #bdc3c7; padding-bottom: 10px;}
        .conclusion { font-size: 1.25rem; background: #e8f8f5; border-left: 6px solid var(--success); padding: 15px 20px; margin: 25px 0; font-weight: bold; text-align: left;}
        table { width: 100%; border-collapse: collapse; margin-top: 20px; box-shadow: 0 4px 6px rgba(0,0,0,0.05); }
        th, td { padding: 15px; text-align: left; border-bottom: 1px solid #ddd; }
        th { background-color: var(--primary); color: white; text-transform: uppercase; font-size: 0.9em; letter-spacing: 0.05em; }
        tr:nth-child(even) { background-color: #fcfcfc; }
        tr:hover { background-color: #f1f2f6; }
        .metric { font-weight: 500; color: var(--primary); }
    </style>
</head>
<body>
    <div class="container">
        <h1>📊 API Performance Report</h1>
        
        <div class="grid">
            <div class="card">
                <h3>🏠 Local Environment</h3>
                <p><strong>Config:</strong> {{ .LocalStats.URL }}</p>
            </div>
            <div class="card">
                <h3>☁️ Server Environment</h3>
                <p><strong>Config:</strong> {{ .ServerStats.URL }}</p>
            </div>
        </div>
        
        <div class="conclusion">
            ✨ Kesimpulan: {{ .FasterEnv }} Environment lebih CEPAT rata-rata {{ printf "%.2f" .DifferenceMs }} ms.
        </div>

        <table>
            <tr>
                <th>Metrik Pengujian</th>
                <th>Local</th>
                <th>Server</th>
            </tr>
            <tr>
                <td class="metric">Total Requests Dikirim</td>
                <td>{{ .LocalStats.TotalRequests }}</td>
                <td>{{ .ServerStats.TotalRequests }}</td>
            </tr>
            <tr>
                <td class="metric">Requests Berhasil (HTTP < 400)</td>
                <td style="color: var(--success); font-weight: bold;">{{ .LocalStats.SuccessCount }}</td>
                <td style="color: var(--success); font-weight: bold;">{{ .ServerStats.SuccessCount }}</td>
            </tr>
            <tr>
                <td class="metric">Requests Gagal (Error)</td>
                <td style="color: {{ if gt .LocalStats.ErrorCount 0 }}var(--danger){{else}}inherit{{end}};">{{ .LocalStats.ErrorCount }}</td>
                <td style="color: {{ if gt .ServerStats.ErrorCount 0 }}var(--danger){{else}}inherit{{end}};">{{ .ServerStats.ErrorCount }}</td>
            </tr>
            <tr>
                <td class="metric">Waktu Tercepat (Min)</td>
                <td>{{ .LocalStats.MinTime }}</td>
                <td>{{ .ServerStats.MinTime }}</td>
            </tr>
            <tr>
                <td class="metric">Waktu Terlambat (Max)</td>
                <td>{{ .LocalStats.MaxTime }}</td>
                <td>{{ .ServerStats.MaxTime }}</td>
            </tr>
            <tr>
                <td class="metric">Rata-Rata Respons (Avg)</td>
                <td><strong>{{ .LocalStats.AvgTime }}</strong></td>
                <td><strong>{{ .ServerStats.AvgTime }}</strong></td>
            </tr>
        </table>
    </div>
</body>
</html>
`
