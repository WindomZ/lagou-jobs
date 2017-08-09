package lagou

import (
	"errors"
	"sync"
)

type Spider struct {
	Config
	Cookies string

	lock      sync.Mutex
	running   bool
	interrupt chan bool
}

func New() *Spider {
	return &Spider{
		interrupt: make(chan bool),
	}
}

func (s *Spider) Start() error {
	if len(s.Cookies) == 0 {
		if err := s.GetCookies(); err != nil {
			return err
		}
	}
	return s.Run()
}

func (s *Spider) Run() error {
	s.lock.Lock()
	if s.running {
		s.lock.Unlock()
		return errors.New("spider is running...")
	}
	s.running = true
	s.lock.Unlock()
	// TODO: ...
	return nil
}

func (s *Spider) Stop() {
	s.lock.Lock()
	if s.running {
		s.running = false
		s.lock.Unlock()
		s.interrupt <- true
	} else {
		s.lock.Unlock()
	}
}
