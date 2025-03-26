package main

import (
	"net/http"
	"net/url"
	"scrape/internal/parser"
	"scrape/internal/utils"

	"golang.org/x/net/html"
)

func main() {
    URLManager := utils.NewURLManager();



    resp, err := http.Get(URLManager.BaseURL);

    if err != nil {
        return;
    }

    doc, err := html.Parse(resp.Body);
    
    url, err := url.Parse(URLManager.BaseURL)

    URLManager.visited = parser.ExtractLinks(doc, url);
}
