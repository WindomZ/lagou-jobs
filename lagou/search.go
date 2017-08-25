package lagou

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/WindomZ/grequests"
	"github.com/WindomZ/lagou-jobs/lagou/entity/mobile"
)

const pageSize int = 15

func (s Spider) searchWithPage(city, keyword string, pageNo int) (*mobile.SearchResponse, error) {
	if !s.running {
		return nil, errors.New("spider stopped")
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
			RequestTimeout: s.Request.RequestTimeout,
		})
	if err != nil {
		return nil, err
	}

	if !resp.Ok {
		return nil, fmt.Errorf("status code: %v", resp.StatusCode)
	}

	response := new(mobile.SearchResponse)
	if err := json.Unmarshal(resp.Bytes(), &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Spider) search(city, keyword string) (<-chan *message, error) {
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

	s.progress.start(totalCount)

	msg := make(chan *message, pages)
	msg <- newMsgData(r)

	go func() {
		wait := sync.WaitGroup{}
		for i := 1; s.running && i < pages; i++ {
			time.Sleep(s.Request.RequestInterval)
			wait.Add(1)
			go func(i int) {
				if r, err := s.searchWithPage(city, keyword, i); err != nil {
					msg <- newMsgError(err)
				} else {
					msg <- newMsgData(r)
				}
				wait.Done()
			}(i)
		}
		wait.Wait()
		close(msg)
	}()

	return msg, nil
}

func (s *Spider) searchPositions(city, keyword string) (<-chan *message, error) {
	msgSR, err := s.search(city, keyword)
	if err != nil {
		return nil, err
	}

	msg := make(chan *message, cap(msgSR)*pageSize)
	go func() {
		wait := sync.WaitGroup{}
		for keep := s.running; keep; {
			select {
			case m := <-msgSR:
				if m == nil {
					keep = false
				} else if m.hasData() {
					wait.Add(1)
					go func(ps []mobile.Position) {
						for _, p := range ps {
							if s.filterCompanyID(p.CompanyID) &&
								s.filterPositionID(p.PositionID) &&
								s.filterSalary(p.Salary) {
								if s.filterString(p.PositionName) {
									msg <- newMsgData(p)
									s.progress.increment()
								} else {
									wait.Add(1)
									time.Sleep(s.Request.RequestInterval)
									go func() {
										if s.filterJobDetail(p.PositionID) {
											msg <- newMsgData(p)
										}
										s.progress.increment()
										wait.Done()
									}()
								}
							}
						}
						wait.Done()
					}(m.data.(*mobile.SearchResponse).Content.Data.Page.Result)
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
		s.progress.finish()
		close(msg)
	}()

	return msg, nil
}

// SearchPositionMap search all positions via city and keyword, returns a PositionMap.
func (s Spider) SearchPositionMap(city, keyword string) (mobile.PositionMap, error) {
	msg, err := s.searchPositions(city, keyword)
	if err != nil {
		return nil, err
	}

	ps := make([]mobile.Position, 0, cap(msg))

	for keep := s.running; keep; {
		select {
		case m := <-msg:
			if m != nil && m.hasData() {
				ps = append(ps, m.data.(mobile.Position))
			} else {
				keep = false
			}
		case <-s.interrupt:
			keep = false
		}
	}

	positions := mobile.PositionMap{}
	positions.Add(ps...)
	return positions, nil
}

// SearchPositionMaps search all positions via city and some keywords, returns a PositionMap.
func (s Spider) SearchPositionMaps(city string, keywords ...string) (mobile.PositionMap, error) {
	positions := mobile.PositionMap{}

	for _, keyword := range keywords {
		pm, err := s.SearchPositionMap(city, keyword)
		if err != nil {
			return nil, err
		}
		positions.Concat(pm)
	}

	return positions, nil
}
