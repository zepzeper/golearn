package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

const baseURL = "https://scrape-me.dreamsofcode.io/";

func main() {
    resp, _ := http.Get(baseURL)

    defer resp.Body.Close()

    // Parse the HTML document
    doc, _ := html.Parse(resp.Body)

    var externalLinks []string;

    findAnchorTags(doc, &externalLinks)

    for _, link := range externalLinks {
        fmt.Println(link)
    }
}

// Helper function to extract the title from HTML
func findAnchorTags(n *html.Node, externalLinks *[]string) string {
    
    if n.Type == html.ElementNode && n.Data == "a" {
        // Found anchor tag
        for _, attr := range n.Attr {
            if attr.Key == "href" {
                if isExternal(attr.Val) {
                    *externalLinks = append(*externalLinks, attr.Val);
                }
            }
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        findAnchorTags(c, externalLinks)
    }

    return ""
}

func isExternal(href string) bool {
    hrefParsed, err := url.Parse(href);
    if err != nil {
        log.Fatalf("Failed to parse href: %v", err)
        return false
    }
    
    // Empty means same url
    return hrefParsed.Hostname() == "";
}
