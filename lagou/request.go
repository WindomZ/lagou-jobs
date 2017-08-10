package lagou

import "time"

type Request struct {
	RequestTimeout  time.Duration
	RequestInterval time.Duration
}

func (s *Spider) initRequestOption() error {
	s.Request.RequestTimeout = time.Second * time.Duration(s.Config.RequestTimeout)
	if s.Request.RequestTimeout <= 0 {
		s.Request.RequestTimeout = time.Second * 15
	}

	s.Request.RequestInterval = time.Millisecond * time.Duration(s.Config.RequestInterval)

	return nil
}
