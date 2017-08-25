package lagou

import (
	"time"

	"gopkg.in/cheggaaa/pb.v1"
)

type progress struct {
	bar *pb.ProgressBar
}

func (p *progress) start(total int) {
	p.bar = pb.StartNew(total)
	p.bar.SetRefreshRate(time.Millisecond * 25)
}

func (p *progress) add(add int) {
	if p.bar != nil {
		p.bar.Add(add)
	}
}

func (p *progress) increment() {
	if p.bar != nil {
		p.bar.Increment()
	}
}

func (p *progress) finish() {
	if p.bar != nil {
		p.bar.Finish()
	}
}
