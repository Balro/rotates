package check

import (
	"bytes"
	"encoding/json"
	"log"
	"net"
	"rotates/conf"
	"strconv"
	"time"
)

type portCheck struct {
	info         portInfo
	suppressed   bool
	suppressTime int
	maxAlert     int
	curAlert     int
}

type portInfo struct {
	Name string
	Host string
	Port int
}

func NewPortChecker(s conf.Service, p conf.Port) *portCheck {
	pc := portCheck{
		info: portInfo{
			Name: p.Name,
			Host: p.Host,
			Port: p.Port,
		},
		suppressed:   false,
		suppressTime: s.SuppressionTime,
		maxAlert:     s.AlertTimes,
		curAlert:     0,
	}
	return &pc
}

func (pc *portCheck) Check(q chan<- string) {
	if pc.checkAlive() {
		log.Printf("%+v is ok.", pc.info)
		pc.suppressed = false
		pc.curAlert = 0
	} else {
		if pc.suppressed {
			log.Printf("Check suppressed. %+v", pc.info)
			return
		}
		q <- pc.GetInfo()
		log.Printf("%+v is lost!", pc.info)
		pc.curAlert += 1
		if pc.curAlert >= pc.maxAlert {
			pc.suppressed = true
			log.Printf("alert reached max alert times %+v", pc.info)
		} else {
			defer pc.unSuppress()
			pc.suppressed = true
			log.Printf("suppressed alert %+v", pc.info)
			time.Sleep(time.Second * time.Duration(pc.suppressTime))
		}
	}
}

func (pc *portCheck) GetInfo() string {
	j, err := json.Marshal(pc.info)
	if err != nil {
		log.Println(err)
	}
	var info bytes.Buffer
	info.WriteString("port lost: \n")
	info.Write(j)
	return info.String()
}

func (pc *portCheck) unSuppress() {
	pc.suppressed = false
	log.Printf("Slept enough time, unsuppressed %+v", pc.info)
}

func (pc *portCheck) checkAlive() bool {
	conn, err := net.DialTimeout("tcp", pc.info.Host+":"+strconv.FormatInt(int64(pc.info.Port), 10), time.Second*5)
	if err != nil {
		log.Println(err)
		return false
	}
	defer conn.Close()
	return true
}
