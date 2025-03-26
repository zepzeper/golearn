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
