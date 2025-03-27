package crawler

import (
	"log"
	"net/url"
)

type Node struct {
    URL string
    Visited bool
    StatusCode int
    Links []string
    IsDead bool
}

type Graph struct {
    Nodes map[string]*Node
    BaseURL *url.URL
    Algorithm string
    MaxDepth int
    NodesCount int
}

func NewGraph (baseURL string, algorithm string, maxDepth int) (*Graph, error) {
    parsed, err := url.Parse(baseURL);
    
    if err != nil {
        log.Fatalf("Failed to parse base url: %s", baseURL)
        return nil, err
    }

    return &Graph{
        Nodes: make(map[string]*Node),
        BaseURL: parsed,
        Algorithm: algorithm,
        MaxDepth: maxDepth,
        NodesCount: 0,
    }, nil
}

func (g *Graph) AddNode(url string) *Node {
    if node, exists := g.Nodes[url]; exists {
        return node
    }

    node := &Node {
        URL: url,
        Visited: false,
        Links: []string{},
    }

    g.Nodes[url] = node
    g.NodesCount++

    return node
}

func (g *Graph) GetNode(url string) (*Node, bool) {
    node, exists := g.Nodes[url]
    return node, exists
}

func (g *Graph) HasNode(url string) bool {
    _, exists := g.Nodes[url]
    return exists
}

func (g *Graph) SetNodeVisited(url string, statusCode int, isDead bool) {
    node, exists := g.Nodes[url]

    if !exists {
        node = g.AddNode(url)
    }

    node.Visited = true
    node.StatusCode = statusCode
    node.IsDead = isDead
}

func (g *Graph) SetNodeLinks(url string, links []string) {
    node, exists := g.Nodes[url]

    if !exists {
        node = g.AddNode(url)
    }

    node.Links = links
}

func (g *Graph) IsSameDomain(link string) bool {
    parsed, err := url.Parse(link)
    if err != nil {
        return false
    }

    return parsed.Host == g.BaseURL.Host
}

func (g *Graph) PrintStats() {
    log.Printf("\n\n -------------- Results -------------- ")
    log.Printf("Crawling completed using %s algorithm", g.Algorithm)
    log.Printf("Total URLs discovered: %d", g.NodesCount)

    deadCount := 0
    for _, node := range g.Nodes {
        if node.IsDead {
            deadCount++
        }
    }

    log.Printf("Dead links found: %d", deadCount)
    log.Printf("Valid links found: %d", g.NodesCount-deadCount)
}
