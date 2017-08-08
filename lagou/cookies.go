package lagou

import (
	"time"

	"github.com/WindomZ/grequests"
)

func (s *Spider) GetCookies() error {
	resp, err := grequests.Get(
		"https://m.lagou.com/",
		&grequests.RequestOptions{
			Headers: map[string]string{
				"Host": "m.lagou.com",
				"Upgrade-Insecure-Requests": "1",
				"User-Agent":                s.UserAgent,
			},
			RequestTimeout: time.Second * 10,
		})
	if err != nil {
		return err
	}

	s.Cookies = resp.Header.Get("Set-Cookie")
	return nil
}
