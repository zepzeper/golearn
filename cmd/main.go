package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"scrape/internal/checker"
	"scrape/internal/crawler"
	"scrape/internal/utils"
	"time"
)

func main() {
    log.SetOutput(os.Stdout)
    log.SetPrefix("[web-crawler] ")

    urlManager := utils.NewURLManager()

    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    pageCrawler := crawler.NewPageCrawler(client)

    links, err := pageCrawler.Crawl(urlManager.BaseURL);

    if err != nil {
        log.Fatalf("Failed to parse base URL: %v", err)
    }


    for _, link := range links {

        if urlManager.IsVisited(link) {
            continue;
        }

        urlManager.MarkVisited(link)

        _, _ = url.Parse(link);
        
        status := pageCrawler.CheckLink(link);

        isDead := status < 200 || status > 400
        _ = checker.NewLinkChecker(link, status, isDead, urlManager.BaseURL)

        if isDead {
            log.Printf("DEAD LINK: %s (Status: %d)", link, status)
        } else {
            log.Printf("VALID LINK: %s (Status: %d)", link, status)
        }
    }

	log.Printf("Crawling completed. Visited %d URLs.", len(urlManager.GetVisited()))
}
