package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const CONTENT = `
		<html>
			<body>
				<a href="/a">About</a>
				<a href="/a?test=1">About with query params</a>
				<a href="https://example.com">External Link</a>
				<a href="http://10.255.255.1">Not Found</a>
			</body>
		</html>
	`

func TestBytesToLinks(t *testing.T) {
	htmlContent := []byte(CONTENT)
	links := bytesToLinks(htmlContent, INITIAL_URL)
	assertEqual(t, 4, len(links))
	assertEqual(t, []LinkInfo{
		{URL: fmt.Sprintf("%s/a", INITIAL_URL), IsExternal: false},
		{URL: fmt.Sprintf("%s/a", INITIAL_URL), IsExternal: false},
		{URL: "https://example.com", IsExternal: true},
		{URL: "http://10.255.255.1", IsExternal: true},
	}, links)
}

func assertEqual(t *testing.T, expected, actual any) {
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Errorf("Mismatch: (-expected, +actual)\n%s", diff)
	}
}

func TestScrapeUrl(t *testing.T) {
	fmt.Println("Starting test")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(CONTENT))
		case "/a":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(CONTENT))
		case "/a?test=1":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(CONTENT))
		case "https://example.com":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`External link`))
		default:
			time.Sleep(6 * time.Second)
		}
	}))
	store := &Store{
		urls: make(map[string]Details),
	}
	client := server.Client()
	client.Timeout = 5 * time.Second
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ctx := context.WithValue(context.Background(), ctxKeyInitialURL, server.URL)
	go scrapeUrl(ctx, client, server.URL, store, true, wg)
	wg.Wait()

	assertEqual(t, 4, len(store.urls))
	assertEqual(t, store.urls, map[string]Details{
		server.URL:            {Status: StatusOK, Code: http.StatusOK},
		server.URL + "/a":     {Status: StatusOK, Code: http.StatusOK},
		"https://example.com": {Status: StatusOK, Code: http.StatusOK},
		"http://10.255.255.1": {Status: StatusError, Code: -1},
	})
}
