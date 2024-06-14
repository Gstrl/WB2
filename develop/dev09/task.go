package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Wget struct {
	processed map[string]struct{}
}

// FromArgs читаем флаги, создаем папку где будем всё хранить, инициалицируем мапу
func (w *Wget) FromArgs() (string, int, error) {
	flagL := flag.Int("l", 0, "specifies the maximum nesting depth")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		return "", 0, fmt.Errorf("usage: go run main.go <URL>")
	}
	inputURL := args[0]

	if inputURL[len(inputURL)-1] == '/' {
		inputURL = inputURL[:len(inputURL)-1]
	}

	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", 0, fmt.Errorf("invalid URL: %v", err)
	}

	folderName := parsedURL.Hostname()
	err = os.Mkdir(folderName, 0o700)
	if err != nil && !os.IsExist(err) {
		return "", 0, fmt.Errorf("failed to create folder: %v", err)
	}
	fmt.Println("Folder successfully created")

	if err := os.Chdir(folderName); err != nil {
		return "", 0, fmt.Errorf("error changing directory: %v", err)
	}

	w.processed = make(map[string]struct{}, 100)
	return inputURL, *flagL, nil
}

// getBody fetches the body of the GET request and returns it
func (w *Wget) getBody(pageURL string) (*http.Response, error) {
	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching file: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	return resp, nil
}

// CreateHTML создаем имя файла из url и сохраняем его
func (w *Wget) createHTML(pageURL string, response *http.Response) error {
	defer response.Body.Close()

	fn := fileName(pageURL)
	path := fmt.Sprintf("%s.html", fn)

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	fmt.Printf("File successfully downloaded and saved as %s\n", path)
	return nil
}

func (w *Wget) Run() {
	startURL, depth, err := w.FromArgs()
	if err != nil {
		log.Fatal(err)
	}

	err = w.download(startURL, depth)
	if err != nil {
		log.Fatal(err)
	}
}

// download рекурсивно проходим по всем ссылкам с учетом глубины l
func (w *Wget) download(pageURL string, depth int) error {
	response, err := w.getBody(pageURL)
	if err != nil {
		return err
	}

	err = w.createHTML(pageURL, response)
	if err != nil {
		return err
	}

	if depth > 0 {
		nestedURLs, err := extractLinks(pageURL)
		if err != nil {
			return err
		}

		for _, link := range nestedURLs {
			// проверям встречался ли этот url ранее
			if _, ok := w.processed[link]; ok {
				continue
			} else {
				w.processed[link] = struct{}{}
			}
			err = w.download(link, depth-1)
			if err != nil {
				log.Printf("Failed to download %s: %v", link, err)
			}
		}
	}

	return nil
}

// extractLinks возвращает найденные на сайте ссылки
func extractLinks(pageURL string) ([]string, error) {
	var links []string

	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %v", err)
	}

	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}
		token := tokenizer.Token()
		if tokenType == html.StartTagToken && token.Data == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					link := attr.Val
					parsedLink, err := baseURL.Parse(link)
					if err != nil {
						continue
					}
					if strings.HasPrefix(parsedLink.String(), "http") {
						fmt.Println(parsedLink.String())
						links = append(links, parsedLink.String())
					}
				}
			}
		}
	}
	uniqueLinks := dedupeStrings(links)

	return uniqueLinks, nil
}

// fileName удаляем то что излишне или не может быть в имени файла
func fileName(pageURL string) string {
	parsedURL, _ := url.Parse(pageURL)
	fn := strings.ReplaceAll(parsedURL.Host+parsedURL.Path, "/", "_")
	extra := []string{".", "?", "=", ","}
	for _, v := range extra {
		fn = strings.ReplaceAll(fn, v, "")
	}
	return fn
}

// dedupeStrings возвращает уникальные элементы слайса
func dedupeStrings(arr []string) []string {
	m, uniq := make(map[string]struct{}), make([]string, 0, len(arr))
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v], uniq = struct{}{}, append(uniq, v)
		}
	}
	return uniq
}

func main() {
	var wget Wget
	wget.Run()
}
