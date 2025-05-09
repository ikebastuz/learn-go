package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

const INITIAL_URL = "https://scrape-me.dreamsofcode.io"
const TIMEOUT = 2 * time.Second

type Status string

const (
	StatusOK      Status = "OK"
	StatusPending Status = "PENDING"
	StatusError   Status = "ERROR"
)

type Details struct {
	Status Status
	Code   int
}

type Store struct {
	urls map[string]Details
	mu   sync.RWMutex
}

func (s *Store) add(url string, details Details) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[url] = details
}

func (s *Store) isVisited(url string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.urls[url]
	return ok
}

func (s *Store) getBrokenUrls() map[string]Details {
	s.mu.RLock()
	defer s.mu.RUnlock()
	brokenUrls := make(map[string]Details)
	for url, details := range s.urls {
		if details.Status == StatusError {
			brokenUrls[url] = details
		}
	}
	return brokenUrls
}

func main() {
	fmt.Println("Starting...")

	store := &Store{
		urls: make(map[string]Details),
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go scrapeUrl(INITIAL_URL, store, true, &wg)

	wg.Wait()

	brokenUrls := store.getBrokenUrls()

	fmt.Println("Broken URLs:")
	for url, details := range brokenUrls {
		fmt.Printf("  %s -> %d\n", url, details.Code)
	}
}

func scrapeUrl(url string, store *Store, followLinks bool, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Scraping:", url)
	client := &http.Client{
		Timeout: TIMEOUT,
	}
	body, err := client.Get(url)
	if err != nil {
		fmt.Println("Error fetching:", err)
		store.add(url, Details{
			Status: StatusError,
			Code:   -1,
		})
		return
	}

	if body.StatusCode >= 300 {
		store.add(url, Details{
			Status: StatusError,
			Code:   body.StatusCode,
		})
		return
	}
	store.add(url, Details{
		Status: StatusOK,
		Code:   body.StatusCode,
	})
	defer body.Body.Close()

	if !followLinks {
		return
	}

	htmlContent, err := io.ReadAll(body.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	z := html.NewTokenizer(strings.NewReader(string(htmlContent)))
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break // End of the document
		}
		t := z.Token()
		if t.Type == html.StartTagToken && t.Data == "a" {
			for _, attr := range t.Attr {
				if attr.Key == "href" {
					link := attr.Val
					isExternalLink := strings.HasPrefix(link, "http")
					nextLink := link
					if !isExternalLink {
						nextLink = fmt.Sprintf("%s%s", INITIAL_URL, link)
					}
					sanitizedLink := strings.Split(nextLink, "?")[0]
					if store.isVisited(sanitizedLink) {
						continue
					}
					wg.Add(1)
					store.add(sanitizedLink, Details{
						Status: StatusPending,
						Code:   -1,
					})
					go scrapeUrl(sanitizedLink, store, !isExternalLink, wg)
				}
			}
		}
	}
}
