package utils

type URLManager struct {
    visited map[string]bool
    BaseURL string
}

func NewURLManager() *URLManager {
   return &URLManager{
        visited: make(map[string]bool),
        BaseURL: "https://scrape-me.dreamsofcode.io/",
    } 
}

func (u *URLManager) IsVisited(url string) bool {
    _, exists := u.visited[url] 

    return exists
}

func (u *URLManager) MarkVisited(url string) {
    u.visited[url] = true
}

func (u *URLManager) GetVisited() map[string]bool {
    return u.visited
}
