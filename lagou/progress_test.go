package lagou

import (
	"testing"
	"time"
)

func TestProgress_1(t *testing.T) {
	p := progress{}
	p.start(100)
	for i := 0; i < 100; i++ {
		p.increment()
		time.Sleep(time.Millisecond * 25)
	}
	p.finish()
}

func TestProgress_2(t *testing.T) {
	p := progress{}
	p.start(100)
	for i := 0; i < 10; i++ {
		p.add(10)
		time.Sleep(time.Millisecond * 250)
	}
	p.finish()
}
