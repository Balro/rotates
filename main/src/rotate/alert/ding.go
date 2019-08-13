package alert

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"rotate/conf"
)

type DingAlert struct {
	Name  string
	Token string
	AtAll bool
	At    []int
}

func NewDingAlert(c conf.Ding) *DingAlert {
	da := DingAlert{
		Name:  c.Name,
		Token: c.Token,
		AtAll: c.AtAll,
		At:    c.At,
	}
	return &da
}

func (da *DingAlert) Send(info string) {
	dJson, err := json.Marshal(getDingMsg(info, da.AtAll, da.At))
	if err != nil {
		log.Println(err)
	}
	log.Printf("DingAlert to %+v\n", string(dJson))
	res, err := http.Post(da.Token, "application/json", bytes.NewReader(dJson))
	if err != nil {
		log.Println(err)
	} else {
		defer res.Body.Close()
	}
}

type dingMsg struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	At struct {
		AtMobiles []int `json:"atMobiles,omitempty"`
		IsAtAll   bool  `json:"isAtAll"`
	} `json:"at"`
}

func getDingMsg(content string, atAll bool, at []int) *dingMsg {
	return &dingMsg{
		Msgtype: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
		At: struct {
			AtMobiles []int `json:"atMobiles,omitempty"`
			IsAtAll   bool  `json:"isAtAll"`
		}{
			IsAtAll:   atAll,
			AtMobiles: at,
		},
	}
}
