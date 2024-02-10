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
			<h1>Go Htmx Home</h1>
			<p>The home page for my Go with HTMX play site</p>
			<ul>	
				<li><a href="/hi">Hi</a></li>
			</ul>
		</body>
		</html>
	`)

	_, _ = w.Write(homePage)

}

func HiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	homePage := []byte(`
		<!DOCTYPE html>
		<html>
		<head>
			<script src="https://unpkg.com/htmx.org@1.9.4/dist/htmx.min.js"></script>
		</head>

		<body>
			<h1>Hi</h1>
			<button hx-get="/greeting" hx-swap="#data-container">Show Greeting</button>
			<div id="data-container"></div>
		</body>
		</html>
	`)

	_, _ = w.Write(homePage)

}

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	homePage := []byte("Hello from the gohtmx project!")

	_, _ = w.Write(homePage)

}
