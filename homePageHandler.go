package main

import (
	"net/http"
)

type HomePageData struct {
	Text string
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	serveHTML(w, nil, r.URL.Path, baseHTML, homePageHTML)
}
