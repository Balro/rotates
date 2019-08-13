package alert

import (
	"log"
	"rotate/conf"
)

type Alerter struct {
	sendables []Sendable
}

func NewAlerter(config conf.Alert) *Alerter {
	ss := make([]Sendable, 0)
	for _, s := range config.Alis {
		ss = append(ss, NewAliAlert(s))
	}
	for _, s := range config.Dings {
		ss = append(ss, NewDingAlert(s))
	}
	return &Alerter{sendables: ss}
}

func (a *Alerter) CronAlert(q <-chan string) {
	for {
		info := <-q
		log.Println("Received alert request.")
		for _, s := range a.sendables {
			go s.Send(info)
		}
	}
}
