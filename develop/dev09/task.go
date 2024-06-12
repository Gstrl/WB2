package main

import (
	"context"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Wget struct {
	depth      int
	nestedUrls []string
}

func (w *Wget) FromArgs() (string, error) {
	flagL := flag.Int("l", 0, "выбрать поля (колонки)")
	flag.Parse()

	args := flag.Args()
	fmt.Println(args)
	if len(args) != 1 {
		return "", fmt.Errorf("использование: go run main.go <URL>")
	}
	url := args[0]

	splitUrl := strings.Split(url, "/")
	folderName := splitUrl[len(splitUrl)-1]

	err := os.Mkdir(folderName, 0700)
	if err != nil {
		log.Fatalf("Не удалось создать папку:%s, %v", folderName, err)
	}
	fmt.Println("Папка успешно создана")

	if err := os.Chdir(folderName); err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		os.Exit(1)
	}

	w.depth = *flagL
	return url, nil
}

// GetBody стоит оставить
func (w *Wget) GetBody(url string) (*http.Response, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("ошибка при скачивании файла: %v", err)
	}

	return response, nil
}

// CreateHTML стоит оставить
func (w *Wget) CreateHTML(url string, response *http.Response) error {
	defer response.Body.Close()

	fn := fileName(url)
	path := fmt.Sprintf("%s.html", fn)

	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return fmt.Errorf("ошибка при создании файла: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("ошибка при записи файла: %v", err)
	}

	fmt.Printf("Файл успешно скачан и сохранен как %s\n", path)
	return nil
}

func (w *Wget) Run() {
	url, err := w.FromArgs()
	if err != nil {
		log.Fatal(err)
	}

	response, err := w.GetBody(url)
	if err != nil {
		log.Fatal(err)
	}
	err = w.CreateHTML(url, response)

	if err != nil {
		log.Fatal(err)
	}

	if w.depth != 0 {
		RunRecursiv(w.depth, url)
	}
}

func RunRecursiv(depth int, url string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*80))
	defer cancel()

	nestedUrls, err := extractLinks(ctx, url)
	if err != nil {
		return
	}

	wget := Wget{
		depth:      depth,
		nestedUrls: nestedUrls,
	}

	defer cancel()

	for _, v := range wget.nestedUrls {
		if v != url {
			response, err := wget.GetBody(v)
			if err != nil {
				continue
			}
			err = wget.CreateHTML(v, response)
			if err != nil {
				continue
			}

			if wget.depth != 0 {
				RunRecursiv(wget.depth-1, v)
			}
		}
	}
}
func extractLinks(ctx context.Context, url string) ([]string, error) {

	links := make([]string, 0)
	select {
	case <-ctx.Done():
		return links, nil
	default:
		resp, err := http.Get(url)
		tokenizer := html.NewTokenizer(resp.Body)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
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
						if strings.HasPrefix(link, "http") {
							fmt.Println(link)
							links = append(links, link)
						}
					}
				}
			}
		}
	}

	return links, nil
}

func fileName(url string) string {
	fn := strings.ReplaceAll(url, "/", "")
	fn = strings.ReplaceAll(fn, ".", "")
	fn = strings.ReplaceAll(fn, "https:", "")
	fn = strings.ReplaceAll(fn, "http:", "")
	fn = strings.ReplaceAll(fn, "?", "")
	fn = strings.ReplaceAll(fn, "=", "")
	fn = strings.ReplaceAll(fn, ",", "")
	return fn
}

func main() {
	var wget Wget
	wget.Run()
}
