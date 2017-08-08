package lagou

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/WindomZ/grequests"
	. "github.com/WindomZ/lagou-jobs/lagou/entity/mobile"
)

const (
	pageSize int           = 15
	timeout  time.Duration = time.Second * 15
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
				"pageSize":     strconv.Itoa(pageSize),
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
			RequestTimeout: timeout,
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

func (s Spider) Search(city, keyword string) (<-chan *Msg, error) {
	r, err := s.search(city, keyword, 0)
	if err != nil {
		return nil, err
	}

	totalCount, err := strconv.Atoi(r.Content.Data.Page.TotalCount)
	if err != nil {
		return nil, err
	}
	totalCount /= pageSize

	msg := make(chan *Msg, totalCount+1)
	msg <- MsgData(r)

	go func() {
		for i := 1; i < totalCount; i++ {
			if s.RequestInterval > 0 {
				time.Sleep(time.Millisecond * time.Duration(s.RequestInterval))
			}
			r, err := s.search(city, keyword, i)
			if err != nil {
				msg <- MsgError(err)
				break
			}
			msg <- MsgData(r)
		}
		close(msg)
	}()

	return msg, nil
}

func (s Spider) searchPositions(city, keyword string) (<-chan *Msg, error) {
	msg_sr, err := s.Search(city, keyword)
	if err != nil {
		return nil, err
	}

	msg := make(chan *Msg, cap(msg_sr)*pageSize)

	go func() {
		for keep := true; keep; {
			select {
			case m := <-msg_sr:
				if m == nil {
					keep = false
					break
				} else if m.HasData() {
					msg <- MsgData(m.data.(*SearchResponse).Content.Data.Page.Result)
				} else {
					msg <- m
				}
			case <-time.After(timeout):
				msg <- MsgInterrupt()
				keep = false
				break
			}
		}
		close(msg)
	}()

	return msg, nil
}

func (s Spider) SearchPositionMap(city, keyword string) (PositionMap, error) {
	msg, err := s.searchPositions(city, keyword)
	if err != nil {
		return nil, err
	}

	positions := PositionMap{}

	for keep := true; keep; {
		select {
		case m := <-msg:
			if m != nil && m.HasData() {
				positions.Add((m.data).([]Position)...)
			} else {
				keep = false
				break
			}
		}
	}

	return positions, nil
}

func (s Spider) SearchPositionMaps(city string, keywords ...string) (PositionMap, error) {
	positions := PositionMap{}

	for _, keyword := range keywords {
		pm, err := s.SearchPositionMap(city, keyword)
		if err != nil {
			return nil, err
		}
		positions.Concat(pm)
	}

	return positions, nil
}
