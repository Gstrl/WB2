package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFromArgs(t *testing.T) {
	os.Args = []string{"cmd", "http://example.com"}
	var wget Wget
	url, depth, err := wget.FromArgs()
	if err != nil {
		t.Fatalf("FromArgs() error: %v", err)
	}
	expectedURL := "http://example.com"
	if url != expectedURL {
		t.Fatalf("Expected URL %s, got %s", expectedURL, url)
	}
	if depth != 0 {
		t.Fatalf("Expected depth 0, got %d", depth)
	}
}

func TestCreateHTML(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<html><body>Hello, World!</body></html>")
	}))
	defer server.Close()

	var wget Wget
	resp, err := wget.getBody(server.URL)
	if err != nil {
		t.Fatalf("getBody() error: %v", err)
	}

	err = wget.createHTML(server.URL, resp)
	if err != nil {
		t.Fatalf("createHTML() error: %v", err)
	}

	filename := fileName(server.URL) + ".html"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("File %s was not created", filename)
	}

	os.Remove(filename)
}

func TestExtractLinks(t *testing.T) {
	testTable := []struct {
		url      string
		expected []string
	}{
		{
			url: "https://go.dev/tour/concurrency/2",
			expected: []string{"https://go.dev/",
				"https://go.dev/tour/list",
			},
		},
		{
			url: "https",
		},
	}

	for i, v := range testTable {
		arrLinks, err := extractLinks(v.url)
		if i == 0 {
			for i, link := range arrLinks {
				if link != v.expected[i] {
					fmt.Println(link, v.expected[i])
					t.Error()
				}
			}
		}
		if i == 1 {
			if err == nil {
				t.Error()
			}
		}
	}
}

func TestFileName(t *testing.T) {
	url := "http://example.com/page1?query=test"
	expected := "examplecom_page1"
	fn := fileName(url)
	fmt.Println(fn)
	if fn != expected {
		t.Fatalf("Expected filename %s, got %s", expected, fn)
	}
}

func TestDedupeStrings(t *testing.T) {
	input := []string{"a", "b", "a", "c", "b"}
	expected := []string{"a", "b", "c"}
	result := dedupeStrings(input)
	if len(result) != len(expected) {
		t.Fatalf("Expected %d unique elements, got %d", len(expected), len(result))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Fatalf("Expected %s at index %d, got %s", expected[i], i, v)
		}
	}
}
