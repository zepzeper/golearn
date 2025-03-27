package algorithms

import (
    "scrape/internal/crawler"
)

type CrawlAlgorithm interface {
    Crawl(graph *crawler.Graph, crawler *crawler.PageCrawler, startURL string) error
}

