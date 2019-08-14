package check

import (
	"rotates/conf"
	"time"
)

type Checker struct {
	checkables []Checkable
	Tic        <-chan time.Time
}

func NewChecker(config conf.Service, checks conf.Check) *Checker {
	cs := make([]Checkable, 0)
	for _, uc := range checks.Urls {
		cs = append(cs, NewUrlChecker(config, uc))
	}
	for _, pc := range checks.Ports {
		cs = append(cs, NewPortChecker(config, pc))
	}
	return &Checker{
		checkables: cs,
		Tic:        time.Tick(time.Second * time.Duration(config.Interval)),
	}
}

func (c *Checker) CronCheck(q chan<- string) {
	for {
		for _, s := range c.checkables {
			go s.Check(q)
		}
		<-c.Tic
	}
}
