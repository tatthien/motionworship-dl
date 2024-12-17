package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

const UrlRegex = `https:\/\/(www.)?motionworship.com\/download(-cloudfront)?\/?\?id=\w+\&type\=4`

func main() {
	targetUrl := os.Args[1]

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		var downloadUrls []string

		urls := e.ChildAttrs("a", "href")

		for _, url := range urls {
			re, err := regexp.Compile(UrlRegex)
			if err != nil {
				continue
			}

			if re.MatchString(url) {
				downloadUrls = append(downloadUrls, url)
			}
		}

		for _, url := range downloadUrls {
			c.Visit(url)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", os.Getenv("COOKIE"))
		fmt.Println("Visiting:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
		}

		fileName := strings.Split(r.Request.URL.Path, "/")[len(strings.Split(r.Request.URL.Path, "/"))-1]
		filePath := homeDir + "/Downloads/MotionWorship/" + fileName
		fmt.Println("Downloading:", fileName, "->", filePath)
		err = downloadFile(filePath, r.Request.URL.String())

		if err != nil {
			fmt.Println("cannot download", r.Request.URL, err)
		}
	})

	c.Visit(targetUrl)
	c.Wait()
}

func downloadFile(filePath, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
