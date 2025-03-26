package crawler

import (
    "log"
    "net/http"
    "net/url"

    "golang.org/x/net/html"
    "scrape/internal/parser"
)

type PageCrawler struct {
    client *http.Client 
}

func NewPageCrawler(client *http.Client) *PageCrawler {
    return &PageCrawler{
        client: client,
    } 
}

func (p *PageCrawler) Crawl(pageURL string) ([]string, error)  {
    // parse url 
    parsedURL, err := url.Parse(pageURL)

    if err != nil {
        return nil, err
    }

    log.Printf("Fetching: %s", pageURL)
    resp, nil := p.client.Get(pageURL)

    if err != nil {
        return []string{}, err
    }

    defer resp.Body.Close()

    doc, err := html.Parse(resp.Body)

    if err != nil {
        return []string{}, err
    }

    links := parser.ExtractLinks(doc, parsedURL)

    return links, nil
}

func (p *PageCrawler) CheckLink(link string) int {
	resp, err := p.client.Head(link)
	if err != nil {
		return 0 // Error code
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
