package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"scrape/internal/algorithms"
	"scrape/internal/crawler"
	"time"
)

func main() {
    log.SetOutput(os.Stdout)
    log.SetPrefix("[web-crawler] ")

    baseURL := flag.String("url", "https://scrape-me.dreamsofcode.io/", "The base URL to start crawling from")
    algorithm := flag.String("algo", "bfs", "Crawling algorithm to use: 'bfs' or 'dfs'")
    maxDepth := flag.Int("depth", -1, "Maximum crawling depth")
    timeout := flag.Int("timeout", 0, "HTTP request timeout in seconds")
    flag.Parse()

    log.Printf("Starting web crawler with %s algorithm on %s", *algorithm, *baseURL)

    client := &http.Client{
        Timeout: time.Duration(*timeout) * time.Second,
    }

    pageCrawler := crawler.NewPageCrawler(client)

    graph, err := crawler.NewGraph(*baseURL, *algorithm, *maxDepth)
    if err != nil {
        log.Fatalf("Failed to initialize web graph: %v", err)
    }

    // Base node
    graph.AddNode(*baseURL)

    algo := GetAlgo(*algorithm)

    startTime := time.Now()
    err = algo.Crawl(graph, pageCrawler, *baseURL)
    if err != nil {
        log.Fatalf("Crawling failed: %v", err)
    }

    // Print statistics
    duration := time.Since(startTime)
    log.Printf("Crawling took %s", duration)
    graph.PrintStats()
}


func GetAlgo(name string) algorithms.CrawlAlgorithm {
	switch name {
      case "dfs":
        return &algorithms.DFSCrawler{}
      default:
        // Default to BFS
        return &algorithms.DFSCrawler{}
	}
}

