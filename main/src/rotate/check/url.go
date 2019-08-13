package check

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"rotate/conf"
	"time"
)

type urlCheck struct {
	info         urlInfo
	suppressed   bool
	suppressTime int
	maxAlert     int
	curAlert     int
}

type urlInfo struct {
	Name   string
	Url    string
	Expect *regexp.Regexp `json:"-"`
}

func NewUrlChecker(s conf.Service, u conf.Url) *urlCheck {
	uc := urlCheck{
		info: urlInfo{
			Name:   u.Name,
			Url:    u.Url,
			Expect: regexp.MustCompile(u.Expect),
		},
		suppressed:   false,
		suppressTime: s.SuppressionTime,
		maxAlert:     s.AlertTimes,
		curAlert:     0,
	}
	return &uc
}

func (uc *urlCheck) Check(q chan<- string) {
	if uc.checkAlive() {
		log.Printf("%+v is ok.", uc.info)
		uc.suppressed = false
		uc.curAlert = 0
	} else {
		if uc.suppressed {
			log.Printf("Check suppressed. %+v", uc.info)
			return
		}
		q <- uc.GetInfo()
		log.Printf("%+v is lost!", uc.info)
		uc.curAlert += 1
		if uc.curAlert >= uc.maxAlert {
			uc.suppressed = true
			log.Printf("alert reached max alert times %+v", uc.info)
		} else {
			defer uc.unSuppress()
			uc.suppressed = true
			log.Printf("suppressed alert %+v", uc.info)
			time.Sleep(time.Second * time.Duration(uc.suppressTime))
		}
	}
}

func (uc *urlCheck) GetInfo() string {
	j, err := json.Marshal(uc.info)
	if err != nil {
		log.Println(err)
	}
	var info bytes.Buffer
	info.WriteString("url lost: \n")
	info.Write(j)
	return info.String()
}

func (uc *urlCheck) unSuppress() {
	uc.suppressed = false
	log.Printf("Slept enough time, unsuppressed %+v", uc.info)
}

func (uc *urlCheck) checkAlive() bool {
	res, err := http.Get(uc.info.Url)
	if err != nil {
		log.Println(err)
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	if !uc.info.Expect.Match(body) {
		return false
	}
	return true
}
