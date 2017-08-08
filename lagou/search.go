package lagou

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/WindomZ/grequests"
	. "github.com/WindomZ/lagou-jobs/lagou/entity/mobile"
)

type SearchResponse struct {
	State   int    `json:"state"`
	Message string `json:"message"`
	Content struct {
		Data struct {
			Custom struct {
				City         string `json:"city"`
				PositionName string `json:"positionName"`
			} `json:"custom"`
			Page struct {
				PageNo     int        `json:"pageNo"`
				PageSize   int        `json:"pageSize"`
				Start      string     `json:"start"`
				TotalCount string     `json:"totalCount"`
				Result     []Position `json:"result"`
			} `json:"page"`
		} `json:"data"`
	} `json:"content"`
}

func (s Spider) search(city, keyword string, pageNo int) (*SearchResponse, error) {
	resp, err := grequests.Get(
		"https://m.lagou.com/search.json",
		&grequests.RequestOptions{
			Params: map[string]string{
				"city":         city,
				"positionName": keyword,
				"pageNo":       strconv.Itoa(pageNo + 1),
				"pageSize":     "15",
			},
			Headers: map[string]string{
				"Accept":           "application/json",
				"Accept-Encoding":  "gzip, deflate, sdch",
				"Host":             "m.lagou.com",
				"Referer":          "https://m.lagou.com/search.html",
				"User-Agent":       s.UserAgent,
				"X-Requested-With": "XMLHttpRequest",
				"Connection":       "keep-alive",
			},
			RequestTimeout: time.Second * 15,
		})
	if err != nil {
		return nil, err
	}

	if !resp.Ok {
		return nil, fmt.Errorf("status code: %v", resp.StatusCode)
	}

	response := new(SearchResponse)
	if err := json.Unmarshal(resp.Bytes(), &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (s Spider) Search(city, keyword string) (*SearchResponse, error) {
	response, err := s.search(city, keyword, 0)
	if err != nil {
		return nil, err
	}

	totalCount, err := strconv.Atoi(response.Content.Data.Page.TotalCount)
	if err != nil {
		return nil, err
	}
	totalCount /= 15

	for i := 1; i < totalCount; i++ {
		r, err := s.search(city, keyword, i)
		if err != nil {
			return nil, err
		}
		response.Content.Data.Page.Result = append(response.Content.Data.Page.Result,
			r.Content.Data.Page.Result...)
	}

	return response, nil
}

func (s Spider) SearchPositions(city, keyword string) (*Positions, error) {
	resp, err := s.Search(city, keyword)
	if err != nil {
		return nil, err
	}

	positions := Positions{}
	positions.Add(resp.Content.Data.Page.Result...)

	return &positions, nil
}
