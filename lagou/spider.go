package lagou

import (
	"errors"
	"sync"
)

type Spider struct {
	Config
	Filter
	Request

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
	if err := s.initFilter(); err != nil {
		return err
	}
	if err := s.initRequestOption(); err != nil {
		return err
	}
	if err := s.initFiles(); err != nil {
		return err
	}
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

	pm, err := s.SearchPositionMaps(s.Config.Search.City,
		s.Config.Search.Keywords...)
	if err != nil {
		return err
	}

	if err := s.writeToFiles(pm); err != nil {
		return err
	}

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
