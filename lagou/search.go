package lagou

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/WindomZ/grequests"
	. "github.com/WindomZ/lagou-jobs/lagou/entity/mobile"
)

const (
	pageSize int           = 15
	timeout  time.Duration = time.Second * 15
)

func (s Spider) searchWithPage(city, keyword string, pageNo int) (*SearchResponse, error) {
	if !s.running {
		return nil, errors.New("spider stopped!")
	}
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
	r, err := s.searchWithPage(city, keyword, 0)
	if err != nil {
		return nil, err
	}

	totalCount, err := strconv.Atoi(r.Content.Data.Page.TotalCount)
	if err != nil {
		return nil, err
	}
	pages := totalCount / pageSize
	if totalCount%pageSize != 0 {
		pages++
	}

	msg := make(chan *Msg, pages)
	msg <- MsgData(r)

	go func() {
		for i := 1; s.running && i < pages; i++ {
			if s.RequestInterval > 0 {
				time.Sleep(time.Millisecond * time.Duration(s.RequestInterval))
			}
			r, err := s.searchWithPage(city, keyword, i)
			if err != nil {
				msg <- MsgError(err)
				break
			}
			msg <- MsgData(r)
		}
		msg <- nil
		time.Sleep(time.Second)
		close(msg)
	}()

	return msg, nil
}

func (s Spider) searchPositions(city, keyword string) (<-chan *Msg, error) {
	msgSR, err := s.Search(city, keyword)
	if err != nil {
		return nil, err
	}

	msg := make(chan *Msg, cap(msgSR)*pageSize)
	go func() {
		wait := &sync.WaitGroup{}
		for keep := s.running; keep; {
			select {
			case m := <-msgSR:
				if m == nil {
					keep = false
				} else if m.HasData() {
					wait.Add(1)
					go s.filterPositions(msg, wait,
						m.data.(*SearchResponse).Content.Data.Page.Result)
				} else {
					keep = false
					msg <- m
				}
			case <-s.interrupt:
				keep = false
				//case <-time.After(timeout):
				//	msg <- MsgInterrupt()
				//	keep = false
			}
		}
		wait.Wait()
		close(msg)
	}()

	return msg, nil
}

func (s Spider) SearchPositionMap(city, keyword string) (PositionMap, error) {
	msg, err := s.searchPositions(city, keyword)
	if err != nil {
		return nil, err
	}

	ps := make([]Position, 0, cap(msg))

	for keep := s.running; keep; {
		select {
		case m := <-msg:
			if m != nil && m.HasData() {
				ps = append(ps, m.data.(Position))
			} else {
				keep = false
			}
		case <-s.interrupt:
			keep = false
		}
	}

	positions := PositionMap{}
	positions.Add(ps...)
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
