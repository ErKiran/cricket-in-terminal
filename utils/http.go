package utils

import (
	"encoding/json"
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

type CommentryResponse struct {
	CommentaryList []struct {
		CommText      string  `json:"commText"`
		Timestamp     int64   `json:"timestamp"`
		BallNbr       int     `json:"ballNbr"`
		OverNumber    float64 `json:"overNumber,omitempty"`
		InningsID     int     `json:"inningsId"`
		Event         string  `json:"event"`
		BatTeamName   string  `json:"batTeamName"`
		OverSeparator struct {
			Score              int      `json:"score"`
			Wickets            int      `json:"wickets"`
			InningsID          int      `json:"inningsId"`
			OSummary           string   `json:"o_summary"`
			Runs               int      `json:"runs"`
			BatStrikerIds      []int    `json:"batStrikerIds"`
			BatStrikerNames    []string `json:"batStrikerNames"`
			BatStrikerRuns     int      `json:"batStrikerRuns"`
			BatStrikerBalls    int      `json:"batStrikerBalls"`
			BatNonStrikerIds   []int    `json:"batNonStrikerIds"`
			BatNonStrikerNames []string `json:"batNonStrikerNames"`
			BatNonStrikerRuns  int      `json:"batNonStrikerRuns"`
			BatNonStrikerBalls int      `json:"batNonStrikerBalls"`
			BowlIds            []int    `json:"bowlIds"`
			BowlNames          []string `json:"bowlNames"`
			BowlOvers          float64  `json:"bowlOvers"`
			BowlMaidens        int      `json:"bowlMaidens"`
			BowlRuns           int      `json:"bowlRuns"`
			BowlWickets        int      `json:"bowlWickets"`
			Timestamp          int64    `json:"timestamp"`
			OverNum            float64  `json:"overNum"`
			BatTeamName        string   `json:"batTeamName"`
			Event              string   `json:"event"`
		} `json:"overSeparator,omitempty"`
		BatsmanStriker struct {
			BatBalls      int     `json:"batBalls"`
			BatDots       int     `json:"batDots"`
			BatFours      int     `json:"batFours"`
			BatID         int     `json:"batId"`
			BatName       string  `json:"batName"`
			BatMins       int     `json:"batMins"`
			BatRuns       int     `json:"batRuns"`
			BatSixes      int     `json:"batSixes"`
			BatStrikeRate float64 `json:"batStrikeRate"`
		} `json:"batsmanStriker,omitempty"`
		BowlerStriker struct {
			BowlID      int     `json:"bowlId"`
			BowlName    string  `json:"bowlName"`
			BowlMaidens int     `json:"bowlMaidens"`
			BowlNoballs int     `json:"bowlNoballs"`
			BowlOvs     float64 `json:"bowlOvs"`
			BowlRuns    int     `json:"bowlRuns"`
			BowlWides   int     `json:"bowlWides"`
			BowlWkts    int     `json:"bowlWkts"`
			BowlEcon    float64 `json:"bowlEcon"`
		} `json:"bowlerStriker,omitempty"`
		CommentaryFormats struct {
			Bold struct {
				FormatID    []string `json:"formatId"`
				FormatValue []string `json:"formatValue"`
			} `json:"bold"`
		} `json:"commentaryFormats,omitempty"`
	}
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

func GetCommentaryUpdates(url, id string) (commentary CommentryResponse) {
	resp, err := http.Get(fmt.Sprintf("%v/cricket-match/commentary/%v", url, id))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &commentary)
	if err != nil {
		log.Fatalln(err)
	}
	return commentary
}
