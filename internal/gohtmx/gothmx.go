package gothmx

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	homePage := []byte(`
		<!DOCTYPE html>
		<html>
		<body>
			<h1>gohtmx</h1>
			<p>The home page for my Go with HTMX play site</p>
		</body>
		</html>
	`)

	_,_ = w.Write(homePage)

}
