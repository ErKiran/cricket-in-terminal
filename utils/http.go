package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CurrentTopic struct {
	Title string
	ID    string
}

func GetCurrentTopics(url string) (c []CurrentTopic) {
	resp, err := http.Get(fmt.Sprintf("%v/html/homepage-scag", url))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := goquery.NewDocumentFromReader(resp.Body)

	body.Find(".cb-font-12").Each(func(i int, s *goquery.Selection) {
		title, exists := s.Attr("title")

		href, exists := s.Attr("href")
		var id string
		if exists {
			id = strings.Split(href, "/")[2]
		}

		c = append(c, CurrentTopic{Title: title, ID: id})
	})
	if err != nil {
		log.Fatalln(err)
	}
	return c
}

func GetCommentaryUpdates(url, id string) {
	resp, err := http.Get(fmt.Sprintf("%v/cricket-match/commentary/%v", url, id))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
