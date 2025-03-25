package parser

import (
	"net/url"
	"golang.org/x/net/html"
)

func ExtractLinks(n *html.Node, baseURL *url.URL) []string {
    var links []string
    ExtractLinksRecursive(n, baseURL, &links);
    return links
}

func ExtractLinksRecursive(n *html.Node, baseURL *url.URL, links *[]string) {

    if n.Type == html.ElementNode && n.Data == "a" {
        // Anchor tag found
        for _, attr := range n.Attr {
            link, err := url.Parse(attr.Val);

            if err != nil {
                continue;
            }

            resolvedURL := baseURL.ResolveReference(link);
            *links = append(*links, resolvedURL.String());
        }
    }

    for c:= n.FirstChild; c != nil; c = c.NextSibling {
        ExtractLinksRecursive(c, baseURL, links)
    }
}
