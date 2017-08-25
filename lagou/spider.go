package lagou

import (
	"errors"
	"sync"
)

// Spider defines lagou spider main body
type Spider struct {
	Config
	Filter
	Request

	Cookies string

	lock      *sync.Mutex
	running   bool
	interrupt chan bool

	progress progress
}

// New returns a Spider instance
func New(configPath string) (*Spider, error) {
	s := &Spider{
		lock:      new(sync.Mutex),
		interrupt: make(chan bool),
	}
	return s, s.Config.ReadConfig(configPath)
}

func (s *Spider) ready() error {
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

	s.lock.Lock()
	if s.running {
		s.lock.Unlock()
		return errors.New("spider is running")
	}
	s.running = true
	s.lock.Unlock()

	return nil
}

func (s *Spider) run() error {
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

// Start start the searching jobs
func (s *Spider) Start() error {
	if err := s.ready(); err != nil {
		return nil
	}
	return s.run()
}

// Stop stop the searching jobs
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
