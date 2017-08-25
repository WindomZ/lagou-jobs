package lagou

import "github.com/WindomZ/grequests"

// GetCookies fetch and set cookies from a http request.
func (s *Spider) GetCookies() error {
	resp, err := grequests.Get(
		"https://m.lagou.com",
		&grequests.RequestOptions{
			Headers: map[string]string{
				"Host": "m.lagou.com",
				"Upgrade-Insecure-Requests": "1",
				"User-Agent":                s.UserAgent,
			},
			RequestTimeout: s.Request.RequestTimeout,
		})
	if err != nil {
		return err
	}

	s.Cookies = resp.Header.Get("Set-Cookie")
	return nil
}
