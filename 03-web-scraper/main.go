package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

const INITIAL_URL = "https://scrape-me.dreamsofcode.io"
const TIMEOUT = 5 * time.Second
const CodeError = -1

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

type LinkInfo struct {
	URL        string
	IsExternal bool
}

type ctxKey string

const ctxKeyInitialURL ctxKey = "initialURL"

func (s *Store) add(url string, details Details) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[url] = details
}

func (s *Store) isVisited(url string) bool {
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
	client := &http.Client{
		Timeout: TIMEOUT,
	}
	var wg sync.WaitGroup
	wg.Add(1)

	ctx := context.WithValue(context.Background(), ctxKeyInitialURL, INITIAL_URL)
	go scrapeUrl(ctx, client, INITIAL_URL, store, true, &wg)

	wg.Wait()

	brokenUrls := store.getBrokenUrls()

	fmt.Println("Broken URLs:")
	for url, details := range brokenUrls {
		fmt.Printf("  %s -> %d\n", url, details.Code)
	}
}

func scrapeUrl(ctx context.Context, client *http.Client, url string, store *Store, followLinks bool, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Scraping:", url)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error fetching:", err)
		store.add(url, Details{
			Status: StatusError,
			Code:   CodeError,
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		store.add(url, Details{
			Status: StatusError,
			Code:   resp.StatusCode,
		})
		return
	}
	store.add(url, Details{
		Status: StatusOK,
		Code:   resp.StatusCode,
	})

	if !followLinks {
		return
	}

	htmlContent, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}
	baseUrl := ctx.Value(ctxKeyInitialURL).(string)
	links := bytesToLinks(htmlContent, baseUrl)

	for _, link := range links {
		if store.isVisited(link.URL) {
			continue
		}
		wg.Add(1)
		store.add(link.URL, Details{
			Status: StatusPending,
			Code:   CodeError,
		})
		go scrapeUrl(ctx, client, link.URL, store, !link.IsExternal, wg)
	}
}

func bytesToLinks(htmlContent []byte, baseUrl string) []LinkInfo {
	var result []LinkInfo
	z := html.NewTokenizer(strings.NewReader(string(htmlContent)))
	base, err := url.Parse(baseUrl)
	if err != nil {
		return result // fallback: return empty if base URL is invalid
	}
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
					var nextLink string
					if isExternalLink {
						nextLink = link
					} else {
						href, err := url.Parse(link)
						if err != nil {
							continue // skip malformed links
						}
						nextLink = base.ResolveReference(href).String()
					}
					// Strip query parameters robustly
					parsed, err := url.Parse(nextLink)
					if err == nil {
						parsed.RawQuery = ""
						nextLink = parsed.String()
					}
					result = append(result, LinkInfo{
						URL:        nextLink,
						IsExternal: isExternalLink,
					})
				}
			}
		}
	}
	return result
}
