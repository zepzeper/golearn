package algorithms

import (
	"log"
	"scrape/internal/crawler"
)

type DFSCrawler struct {}


func (d *DFSCrawler) Crawl(graph *crawler.Graph, crawler *crawler.PageCrawler, startURL string) error {
    var stack []string; 
    stack = append(stack, startURL);

    for len(stack) > 0 {
        currentURL := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1] // LIFO

        if node, exists := graph.GetNode(currentURL); exists && node.Visited { continue; }

        graph.AddNode(currentURL)
        log.Printf("DFS Processing: %s", currentURL)

        links, err := crawler.Crawl(currentURL)
        if err != nil {
            log.Fatalf("Failed to parse base URL: %v", err)
        }


        if node, exists := graph.GetNode(currentURL); exists && node.Visited { continue; }


        if err != nil {
            log.Printf("Error crawling %s: %v", currentURL, err)

            graph.SetNodeVisited(currentURL, 0, true)
            continue
        }

        status := crawler.CheckLink(currentURL);
        isDead := status < 200 || status > 300

        graph.SetNodeVisited(currentURL, status, isDead)
        graph.SetNodeLinks(currentURL, links)

        if isDead {
            log.Printf("DEAD LINK: %s (Status: %d)", currentURL, status)
        } else {
            log.Printf("VALID LINK: %s (Status: %d)", currentURL, status)
        }

        for _, link := range links {

            if !graph.HasNode(link) && graph.IsSameDomain(link) {
                graph.AddNode(link)

                stack = append(stack, link)
            }
        }
    }

    return nil
}
