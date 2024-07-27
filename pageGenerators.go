package main

import (
	"fmt"
	"net/http"
)

func printDefaultPage(w http.ResponseWriter) error {
	_, err := fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
<body>
<iframe name="dummyframe" id="dummyframe" style="display: none;"></iframe>
<div class="container">
	<div class="vertical-center">
		<form action="/switch" method="get" target="dummyframe">
		  <button type="submit">Switch DPI protection (current status: `+getDpiProtectionStatus()+`)</button>
		</form>
	</div>
</div>
</body>
<style>
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
`)
	return err
}
