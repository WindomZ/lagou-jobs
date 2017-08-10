package lagou

import (
	"time"

	"gopkg.in/cheggaaa/pb.v1"
)

type Progress struct {
	bar *pb.ProgressBar
}

func (p *Progress) Start(total int) {
	p.bar = pb.StartNew(total)
	p.bar.SetRefreshRate(time.Millisecond * 25)
}

func (p *Progress) Add(add int) {
	if p.bar != nil {
		p.bar.Add(add)
	}
}

func (p *Progress) Increment() {
	if p.bar != nil {
		p.bar.Increment()
	}
}

func (p *Progress) Finish() {
	if p.bar != nil {
		p.bar.Finish()
	}
}
