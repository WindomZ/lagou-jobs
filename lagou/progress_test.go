package lagou

import (
	"testing"
	"time"
)

func TestProgress_1(t *testing.T) {
	p := Progress{}
	p.Start(100)
	for i := 0; i < 100; i++ {
		p.Increment()
		time.Sleep(time.Millisecond * 25)
	}
	p.Finish()
}

func TestProgress_2(t *testing.T) {
	p := Progress{}
	p.Start(100)
	for i := 0; i < 10; i++ {
		p.Add(10)
		time.Sleep(time.Millisecond * 250)
	}
	p.Finish()
}