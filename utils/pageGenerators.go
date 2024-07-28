package utils

import (
	"fmt"
	"net/http"
)

func PrintDefaultPage(w http.ResponseWriter) error {
	dpiProtectionStatus := GetDpiProtectionStatus()
	loadAverage := GetLoadAverage()
	memoryStats := GetMemoryStats()
	memoryUsedPercentage := float64(memoryStats.MemTotal-memoryStats.MemAvailable) / float64(memoryStats.MemTotal) * 100
	_, err := fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<body>
<iframe name="dummyframe" id="dummyframe" style="display: none;"></iframe>
<div class="device-stats">
<div class="load-average"> Load Average: %.2f %.2f %.2f. Last created PID: %d</div>
<div class="memory-stats"> Memory Stats: %d kB available of %d (used %.2f %%)</div>
</div>
<div class="container">
	<div class="vertical-center">
		<form action="/switch" method="get" target="dummyframe">
		  <button type="submit">Switch DPI protection (current status: %s)</button>
		</form>
	</div>
</div>
</body>
<style>
* {
	font-size: 32px;
}

.container {
	width: 100vw;
	height: 100vh;
}

.vertical-center {
	margin: 0 auto;
}
button {
	height: 200px;
	width: 90vw;
	margin: 0 auto;
    font-size: 32px;
}
</style>
</html>
`, loadAverage.Load1, loadAverage.Load5, loadAverage.Load15, loadAverage.LastCreatedPid,
		memoryStats.MemAvailable, memoryStats.MemTotal, memoryUsedPercentage,
		dpiProtectionStatus)

	return err
}
