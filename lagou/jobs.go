package lagou

import (
	"errors"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/WindomZ/grequests"
	"github.com/WindomZ/lagou-jobs/lagou/entity/mobile"
)

func (s Spider) crawlJobDetail(positionID int) (*mobile.JobDetail, error) {
	if !s.running {
		return nil, errors.New("spider stopped")
	}
	resp, err := grequests.Get(
		fmt.Sprintf("https://m.lagou.com/jobs/%v.html", positionID),
		&grequests.RequestOptions{
			Headers: map[string]string{
				"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
				"Host":                      "m.lagou.com",
				"Referer":                   "https://m.lagou.com/search.html",
				"Cookie":                    s.Cookies,
				"User-Agent":                s.UserAgent,
				"Upgrade-Insecure-Requests": "1",
			},
			RequestTimeout: s.Request.RequestTimeout,
		})
	if err != nil {
		return nil, err
	}

	if !resp.Ok {
		return nil, fmt.Errorf("status code: %v", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromResponse(resp.RawResponse)
	if err != nil {
		return nil, err
	}

	job := new(mobile.JobDetail)

	job.Meta.Keywords = doc.Find("meta[name=keywords]").
		AttrOr("content", "")
	job.Meta.Description = doc.Find("meta[name=description]").
		AttrOr("content", "")

	job.Title = strings.TrimSpace(doc.Find("div.postitle h2.title").Text())
	job.Salary = strings.TrimSpace(doc.Find("span.salary span.text").Text())
	job.WorkAddress = strings.TrimSpace(doc.Find("span.workaddress span.text").Text())
	job.JobNature = strings.TrimSpace(doc.Find("span.jobnature span.text").Text())
	job.WorkYear = strings.TrimSpace(doc.Find("span.workyear span.text").Text())
	job.Education = strings.TrimSpace(doc.Find("span.education span.text").Text())
	job.Temptation = strings.TrimSpace(doc.Find("div.temptation").Text())

	job.Company.Title = strings.TrimSpace(doc.Find("div.company div.desc h2.title").Text())
	job.Company.Title = strings.Replace(strings.Replace(job.Company.Title,
		" ", "", -1),
		"\n", "", -1)
	job.Company.Info = strings.TrimSpace(doc.Find("div.company div.desc p.info").Text())
	job.Company.Info = strings.Replace(strings.Replace(job.Company.Info,
		" ", "", -1),
		"\n", "", -1)

	return job, nil
}

func (s Spider) filterJobDetail(positionID int) bool {
	j, err := s.crawlJobDetail(positionID)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	if !s.filterSalary(j.Salary) {
		return false
	}

	switch {
	case s.filterString(j.Meta.Keywords):
	case s.filterString(j.Meta.Description):
	case s.filterString(j.Title):
	case s.filterString(j.Temptation):
	case s.filterString(j.PositionDesc):
	default:
		return false
	}

	return true
}
