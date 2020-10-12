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

type Heram struct {
	CommentaryList []struct {
		CommText       string  `json:"commText"`
		Timestamp      int64   `json:"timestamp"`
		BallNbr        int     `json:"ballNbr"`
		OverNumber     float64 `json:"overNumber,omitempty"`
		InningsID      int     `json:"inningsId"`
		Event          string  `json:"event"`
		BatTeamName    string  `json:"batTeamName"`
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
	} `json:"commentaryList"`
	MatchHeader struct {
		MatchID                int    `json:"matchId"`
		MatchDescription       string `json:"matchDescription"`
		MatchFormat            string `json:"matchFormat"`
		MatchType              string `json:"matchType"`
		Complete               bool   `json:"complete"`
		Domestic               bool   `json:"domestic"`
		MatchStartTimestamp    int64  `json:"matchStartTimestamp"`
		MatchCompleteTimestamp int64  `json:"matchCompleteTimestamp"`
		DayNight               bool   `json:"dayNight"`
		Year                   int    `json:"year"`
		State                  string `json:"state"`
		Status                 string `json:"status"`
		TossResults            struct {
			TossWinnerID   int    `json:"tossWinnerId"`
			TossWinnerName string `json:"tossWinnerName"`
			Decision       string `json:"decision"`
		} `json:"tossResults"`
		Result struct {
			WinningTeam  string `json:"winningTeam"`
			WinByRuns    bool   `json:"winByRuns"`
			WinByInnings bool   `json:"winByInnings"`
		} `json:"result"`
		RevisedTarget struct {
			Reason string `json:"reason"`
		} `json:"revisedTarget"`
		PlayersOfTheMatch  []interface{} `json:"playersOfTheMatch"`
		PlayersOfTheSeries []interface{} `json:"playersOfTheSeries"`
		MatchTeamInfo      []struct {
			BattingTeamID        int    `json:"battingTeamId"`
			BattingTeamShortName string `json:"battingTeamShortName"`
			BowlingTeamID        int    `json:"bowlingTeamId"`
			BowlingTeamShortName string `json:"bowlingTeamShortName"`
		} `json:"matchTeamInfo"`
		IsMatchNotCovered bool `json:"isMatchNotCovered"`
		Team1             struct {
			ID            int           `json:"id"`
			Name          string        `json:"name"`
			PlayerDetails []interface{} `json:"playerDetails"`
			ShortName     string        `json:"shortName"`
		} `json:"team1"`
		Team2 struct {
			ID            int           `json:"id"`
			Name          string        `json:"name"`
			PlayerDetails []interface{} `json:"playerDetails"`
			ShortName     string        `json:"shortName"`
		} `json:"team2"`
		SeriesDesc string `json:"seriesDesc"`
	} `json:"matchHeader"`
	Miniscore struct {
		InningsID      int `json:"inningsId"`
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
		} `json:"batsmanStriker"`
		BatsmanNonStriker struct {
			BatBalls      int     `json:"batBalls"`
			BatDots       int     `json:"batDots"`
			BatFours      int     `json:"batFours"`
			BatID         int     `json:"batId"`
			BatName       string  `json:"batName"`
			BatMins       int     `json:"batMins"`
			BatRuns       int     `json:"batRuns"`
			BatSixes      int     `json:"batSixes"`
			BatStrikeRate float64 `json:"batStrikeRate"`
		} `json:"batsmanNonStriker"`
		BatTeam struct {
			TeamID    int `json:"teamId"`
			TeamScore int `json:"teamScore"`
			TeamWkts  int `json:"teamWkts"`
		} `json:"batTeam"`
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
		} `json:"bowlerStriker"`
		BowlerNonStriker struct {
			BowlID      int     `json:"bowlId"`
			BowlName    string  `json:"bowlName"`
			BowlMaidens int     `json:"bowlMaidens"`
			BowlNoballs int     `json:"bowlNoballs"`
			BowlOvs     float64 `json:"bowlOvs"`
			BowlRuns    int     `json:"bowlRuns"`
			BowlWides   int     `json:"bowlWides"`
			BowlWkts    int     `json:"bowlWkts"`
			BowlEcon    float64 `json:"bowlEcon"`
		} `json:"bowlerNonStriker"`
		Overs          float64 `json:"overs"`
		RecentOvsStats string  `json:"recentOvsStats"`
		Target         int     `json:"target"`
		PartnerShip    struct {
			Balls int `json:"balls"`
			Runs  int `json:"runs"`
		} `json:"partnerShip"`
		CurrentRunRate    float64 `json:"currentRunRate"`
		RequiredRunRate   float64 `json:"requiredRunRate"`
		MatchScoreDetails struct {
			MatchID          int `json:"matchId"`
			InningsScoreList []struct {
				InningsID   int     `json:"inningsId"`
				BatTeamID   int     `json:"batTeamId"`
				BatTeamName string  `json:"batTeamName"`
				Score       int     `json:"score"`
				Wickets     int     `json:"wickets"`
				Overs       float64 `json:"overs"`
				IsDeclared  bool    `json:"isDeclared"`
				IsFollowOn  bool    `json:"isFollowOn"`
			} `json:"inningsScoreList"`
			TossResults struct {
				TossWinnerID   int    `json:"tossWinnerId"`
				TossWinnerName string `json:"tossWinnerName"`
				Decision       string `json:"decision"`
			} `json:"tossResults"`
			MatchTeamInfo []struct {
				BattingTeamID        int    `json:"battingTeamId"`
				BattingTeamShortName string `json:"battingTeamShortName"`
				BowlingTeamID        int    `json:"bowlingTeamId"`
				BowlingTeamShortName string `json:"bowlingTeamShortName"`
			} `json:"matchTeamInfo"`
			IsMatchNotCovered bool   `json:"isMatchNotCovered"`
			MatchFormat       string `json:"matchFormat"`
			State             string `json:"state"`
			CustomStatus      string `json:"customStatus"`
			HighlightedTeamID int    `json:"highlightedTeamId"`
		} `json:"matchScoreDetails"`
		LatestPerformance []interface{} `json:"latestPerformance"`
		PpData            struct {
			Pp1 struct {
				PpID        int     `json:"ppId"`
				PpOversFrom float64 `json:"ppOversFrom"`
				PpOversTo   float64 `json:"ppOversTo"`
				PpType      string  `json:"ppType"`
				RunsScored  int     `json:"runsScored"`
			} `json:"pp_1"`
		} `json:"ppData"`
		OverSummaryList []interface{} `json:"overSummaryList"`
	} `json:"miniscore"`
	CommentarySnippetList []struct {
		CommID        int           `json:"commId"`
		MatchID       int           `json:"matchId"`
		InningsID     int           `json:"inningsId"`
		InfraType     string        `json:"infraType"`
		Headline      string        `json:"headline"`
		Content       string        `json:"content"`
		Timestamp     int64         `json:"timestamp"`
		ParsedContent []interface{} `json:"parsedContent"`
		IsLive        bool          `json:"isLive"`
	} `json:"commentarySnippetList"`
	Page            string `json:"page"`
	EnableNoContent bool   `json:"enableNoContent"`
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
		MatchHeader struct {
			MatchID                int    `json:"matchId"`
			MatchDescription       string `json:"matchDescription"`
			MatchFormat            string `json:"matchFormat"`
			MatchType              string `json:"matchType"`
			Complete               bool   `json:"complete"`
			Domestic               bool   `json:"domestic"`
			MatchStartTimestamp    int64  `json:"matchStartTimestamp"`
			MatchCompleteTimestamp int64  `json:"matchCompleteTimestamp"`
			DayNight               bool   `json:"dayNight"`
			Year                   int    `json:"year"`
			State                  string `json:"state"`
			Status                 string `json:"status"`
			TossResults            struct {
				TossWinnerID   int    `json:"tossWinnerId"`
				TossWinnerName string `json:"tossWinnerName"`
				Decision       string `json:"decision"`
			} `json:"tossResults"`
			Result struct {
				WinningTeam  string `json:"winningTeam"`
				WinByRuns    bool   `json:"winByRuns"`
				WinByInnings bool   `json:"winByInnings"`
			} `json:"result"`
			RevisedTarget struct {
				Reason string `json:"reason"`
			} `json:"revisedTarget"`
			PlayersOfTheMatch  []interface{} `json:"playersOfTheMatch"`
			PlayersOfTheSeries []interface{} `json:"playersOfTheSeries"`
			MatchTeamInfo      []struct {
				BattingTeamID        int    `json:"battingTeamId"`
				BattingTeamShortName string `json:"battingTeamShortName"`
				BowlingTeamID        int    `json:"bowlingTeamId"`
				BowlingTeamShortName string `json:"bowlingTeamShortName"`
			} `json:"matchTeamInfo"`
			IsMatchNotCovered bool `json:"isMatchNotCovered"`
			Team1             struct {
				ID            int           `json:"id"`
				Name          string        `json:"name"`
				PlayerDetails []interface{} `json:"playerDetails"`
				ShortName     string        `json:"shortName"`
			} `json:"team1"`
			Team2 struct {
				ID            int           `json:"id"`
				Name          string        `json:"name"`
				PlayerDetails []interface{} `json:"playerDetails"`
				ShortName     string        `json:"shortName"`
			} `json:"team2"`
			SeriesDesc string `json:"seriesDesc"`
		} `json:"matchHeader"`
		Miniscore struct {
			InningsID      int `json:"inningsId"`
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
			} `json:"batsmanStriker"`
			BatsmanNonStriker struct {
				BatBalls      int     `json:"batBalls"`
				BatDots       int     `json:"batDots"`
				BatFours      int     `json:"batFours"`
				BatID         int     `json:"batId"`
				BatName       string  `json:"batName"`
				BatMins       int     `json:"batMins"`
				BatRuns       int     `json:"batRuns"`
				BatSixes      int     `json:"batSixes"`
				BatStrikeRate float64 `json:"batStrikeRate"`
			} `json:"batsmanNonStriker"`
			BatTeam struct {
				TeamID    int `json:"teamId"`
				TeamScore int `json:"teamScore"`
				TeamWkts  int `json:"teamWkts"`
			} `json:"batTeam"`
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
			} `json:"bowlerStriker"`
			BowlerNonStriker struct {
				BowlID      int     `json:"bowlId"`
				BowlName    string  `json:"bowlName"`
				BowlMaidens int     `json:"bowlMaidens"`
				BowlNoballs int     `json:"bowlNoballs"`
				BowlOvs     float64 `json:"bowlOvs"`
				BowlRuns    int     `json:"bowlRuns"`
				BowlWides   int     `json:"bowlWides"`
				BowlWkts    int     `json:"bowlWkts"`
				BowlEcon    float64 `json:"bowlEcon"`
			} `json:"bowlerNonStriker"`
			Overs          float64 `json:"overs"`
			RecentOvsStats string  `json:"recentOvsStats"`
			Target         int     `json:"target"`
			PartnerShip    struct {
				Balls int `json:"balls"`
				Runs  int `json:"runs"`
			} `json:"partnerShip"`
			CurrentRunRate    float64 `json:"currentRunRate"`
			RequiredRunRate   float64 `json:"requiredRunRate"`
			MatchScoreDetails struct {
				MatchID          int `json:"matchId"`
				InningsScoreList []struct {
					InningsID   int     `json:"inningsId"`
					BatTeamID   int     `json:"batTeamId"`
					BatTeamName string  `json:"batTeamName"`
					Score       int     `json:"score"`
					Wickets     int     `json:"wickets"`
					Overs       float64 `json:"overs"`
					IsDeclared  bool    `json:"isDeclared"`
					IsFollowOn  bool    `json:"isFollowOn"`
				} `json:"inningsScoreList"`
				TossResults struct {
					TossWinnerID   int    `json:"tossWinnerId"`
					TossWinnerName string `json:"tossWinnerName"`
					Decision       string `json:"decision"`
				} `json:"tossResults"`
				MatchTeamInfo []struct {
					BattingTeamID        int    `json:"battingTeamId"`
					BattingTeamShortName string `json:"battingTeamShortName"`
					BowlingTeamID        int    `json:"bowlingTeamId"`
					BowlingTeamShortName string `json:"bowlingTeamShortName"`
				} `json:"matchTeamInfo"`
				IsMatchNotCovered bool   `json:"isMatchNotCovered"`
				MatchFormat       string `json:"matchFormat"`
				State             string `json:"state"`
				CustomStatus      string `json:"customStatus"`
				HighlightedTeamID int    `json:"highlightedTeamId"`
			} `json:"matchScoreDetails"`
			LatestPerformance []interface{} `json:"latestPerformance"`
			PpData            struct {
				Pp1 struct {
					PpID        int     `json:"ppId"`
					PpOversFrom float64 `json:"ppOversFrom"`
					PpOversTo   float64 `json:"ppOversTo"`
					PpType      string  `json:"ppType"`
					RunsScored  int     `json:"runsScored"`
				} `json:"pp_1"`
			} `json:"ppData"`
			OverSummaryList []interface{} `json:"overSummaryList"`
		} `json:"miniscore"`
		CommentarySnippetList []struct {
			CommID        int           `json:"commId"`
			MatchID       int           `json:"matchId"`
			InningsID     int           `json:"inningsId"`
			InfraType     string        `json:"infraType"`
			Headline      string        `json:"headline"`
			Content       string        `json:"content"`
			Timestamp     int64         `json:"timestamp"`
			ParsedContent []interface{} `json:"parsedContent"`
			IsLive        bool          `json:"isLive"`
		} `json:"commentarySnippetList"`
		Page            string `json:"page"`
		EnableNoContent bool   `json:"enableNoContent"`
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

func GetCommentaryUpdates(url, id string) (commentary Heram) {
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
